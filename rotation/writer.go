package rotation

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/joomcode/errorx"
)

type WriterOptions struct {
	Prefix     string
	MaxSize    int64
	MaxBackups int
	MaxAge     time.Duration
}

type Writer struct {
	prefix     string
	maxSize    int64
	maxBackups int
	maxAge     time.Duration

	fileName    string
	file        *os.File
	fileWritten int64

	rotateMu *sync.Mutex

	writeMu *sync.Mutex

	closed bool
}

func NewWriter(opts *WriterOptions) (*Writer, error) {
	if opts == nil || opts.Prefix == "" {
		return nil, commonErrors.New("prefix must be provided")
	}

	w := &Writer{
		prefix:     opts.Prefix,
		maxSize:    1024 * 1024, // 1MB
		maxBackups: 3,
		maxAge:     time.Hour * 24 * 7, // 7 days

		fileName:    makeLogName(opts.Prefix),
		file:        nil,
		fileWritten: 0,

		rotateMu: &sync.Mutex{},

		writeMu: &sync.Mutex{},

		closed: false,
	}
	if opts.MaxSize > 0 {
		w.maxSize = opts.MaxSize
	}
	if opts.MaxBackups > 0 {
		w.maxBackups = opts.MaxBackups
	}
	if opts.MaxAge > 0 {
		w.maxAge = opts.MaxAge
	}

	logName := makeLogName(w.prefix)

	// first rotate if necessary
	if stat, err := os.Stat(logName); err == nil && stat.Size() > 0 {
		if err := firstRotate(w.prefix, logName); err != nil {
			return nil, commonErrors.Wrap(err, "failed to make first rotate")
		}
	}

	file, err := os.OpenFile(logName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to open log file")
	}
	w.file = file

	return w, nil
}

func (self *Writer) Write(data []byte) (int, error) {
	self.writeMu.Lock()
	defer self.writeMu.Unlock()

	if self.closed {
		return 0, commonErrors.New("writer is closed")
	}

	dataLen := len(data)
	if self.fileWritten+int64(dataLen) > self.maxSize {
		if err := self.rotate(true); err != nil {
			panic(errorx.Decorate(err, "failed to rotate log file"))
		}
	}

	if _, err := self.file.Write(data); err != nil {
		panic(errorx.Decorate(err, "failed to write to log file"))
	}
	self.fileWritten += int64(dataLen)

	return 0, nil
}

func (self *Writer) Close() {
	self.writeMu.Lock()
	self.closed = true
	written := self.fileWritten
	self.writeMu.Unlock()

	if written > 0 {
		if err := self.rotate(false); err != nil {
			panic(errorx.Decorate(err, "failed to rotate log file"))
		}
	}
}

func (self *Writer) rotate(inGoro bool) error {
	buf := self.flushToBuffer()

	if inGoro {
		go func() {
			self.rotateMu.Lock()
			defer self.rotateMu.Unlock()
			self.removeFilesAndCompress(buf)
		}()
	} else {
		self.rotateMu.Lock()
		defer self.rotateMu.Unlock()
		self.removeFilesAndCompress(buf)
	}

	return nil
}

func (self *Writer) removeFilesAndCompress(buf *bytes.Buffer) {
	removeExpiredFiles(self.prefix, self.fileName, self.maxAge)
	removeOldFiles(self.prefix, self.fileName, self.maxBackups-1) // cause a new file will be created

	if err := compressBuffer(buf, self.prefix); err != nil {
		panic(errorx.Decorate(err, "failed to compress log file"))
	}
}

func (self *Writer) flushToBuffer() *bytes.Buffer {
	// panic is better, cause we can't recover

	if _, err := self.file.Seek(0, 0); err != nil {
		panic(errorx.Decorate(err, "failed to seek log file"))
	}

	buf := &bytes.Buffer{}

	if _, err := io.Copy(buf, self.file); err != nil {
		panic(errorx.Decorate(err, "failed to write rotated log file"))
	}

	if _, err := self.file.Seek(0, 0); err != nil {
		panic(errorx.Decorate(err, "failed to seek log file"))
	}

	if err := self.file.Truncate(0); err != nil {
		panic(errorx.Decorate(err, "failed to truncate log file"))
	}

	self.fileWritten = 0

	return buf
}

func compressBuffer(buf *bytes.Buffer, prefix string) error {
	z, err := createLogZip(prefix)
	if err != nil {
		return errorx.Decorate(err, "failed to create zip")
	}

	if _, err := io.Copy(z, buf); err != nil {
		return errorx.DecorateMany("failed to write zipped log file", err,
			z.closeAndRemove())
	}

	if err := z.close(); err != nil {
		return errorx.Decorate(err, "failed to close zip writer")
	}

	return nil
}

func firstRotate(prefix, logName string) error {
	src, err := os.Open(logName)
	if err != nil {
		return errorx.Decorate(err, "failed to open log file")
	}

	z, err := createLogZip(prefix)
	if err != nil {
		return errorx.Decorate(err, "failed to create zip")
	}

	_, err = io.Copy(z, src)
	if err != nil {
		return errorx.DecorateMany("failed to write rotated log file", err, z.close())
	}
	if err := z.close(); err != nil {
		return errorx.Decorate(err, "failed to close zip")
	}

	return nil
}

func removeOldFiles(prefix, logName string, maxBackups int) {
	files := []*zipFile{}

	if err := iterZipFiles(prefix, logName, func(f *zipFile) {
		files = append(files, f)
	}); err != nil {
		panic(errorx.Decorate(err, "failed to iterate zip files"))
	}

	slices.SortFunc(files, func(a, b *zipFile) int {
		return a.createTime.Compare(b.createTime)
	})

	if len(files) > maxBackups {
		for i := 0; i < len(files)-maxBackups; i++ {
			if err := os.Remove(files[i].name); err != nil {
				panic(errorx.Decorate(err, "failed to remove old log file"))
			}
		}
	}
}

func removeExpiredFiles(prefix, logName string, maxAge time.Duration) {
	now := time.Now().UTC()

	if err := iterZipFiles(prefix, logName, func(f *zipFile) {
		if now.Sub(f.createTime) > maxAge {
			if err := os.Remove(f.name); err != nil {
				panic(errorx.Decorate(err, "failed to remove expired log file"))
			}
		}
	}); err != nil {
		panic(errorx.Decorate(err, "failed to iterate zip files"))
	}
}

type zipFile struct {
	name       string
	createTime time.Time
}

func iterZipFiles(prefix, logName string, f func(*zipFile)) error {
	dir := filepath.Dir(logName)
	prefix = filepath.Base(prefix)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return errorx.Decorate(err, "failed to read log directory")
	}

	for _, entry := range entries {
		if entry.IsDir() ||
			!strings.HasPrefix(entry.Name(), prefix) ||
			!strings.HasSuffix(entry.Name(), zipExt) {
			continue
		}

		t, ok := extractCreateTimeFromZipName(entry.Name())
		if !ok {
			continue
		}

		f(&zipFile{filepath.Join(dir, entry.Name()), t})
	}

	return nil
}
