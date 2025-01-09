package rotation

import (
	"compress/gzip"
	"os"

	"github.com/joomcode/errorx"
)

const zipExt = ".log.gz"
const zipTimeLayout = "2006-01-02T15-04-05"

type zip struct {
	filename string
	file     *os.File
	writer   *gzip.Writer
}

// * zip must be closed
func createLogZip(prefix string) (*zip, error) {
	filename := makeZipName(prefix)

	f, err := os.Create(filename)
	if err != nil {
		return nil, errorx.Decorate(err, "failed to create log file")
	}

	z := gzip.NewWriter(f)

	return &zip{filename, f, z}, nil
}

func (self *zip) Write(data []byte) (int, error) {
	return self.writer.Write(data)
}

func (self *zip) closeAndRemove() error {
	if err := self.close(); err != nil {
		return err
	}
	return os.Remove(self.filename)
}

func (self *zip) close() error {
	if err := self.writer.Close(); err != nil {
		return errorx.Decorate(err, "failed to close zip writer")
	}
	if err := self.file.Close(); err != nil {
		return errorx.Decorate(err, "failed to close log file")
	}
	return nil
}
