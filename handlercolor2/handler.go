package handlercolor2

import (
	"context"
	"encoding"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unicode"

	"github.com/av1ppp/logx"
	"github.com/fatih/color"
)

const errKey = "err"

var (
	defaultLevel      = logx.LevelInfo
	defaultTimeFormat = time.StampMilli
)

// Options for a slog.Handler that writes tinted logs. A zero Options consists
// entirely of default values.
//
// Options can be used as a drop-in replacement for [slog.HandlerOptions].
type Options struct {
	// Enable source code location (Default: false)
	AddSource bool

	// Minimum level to log (Default: logx.LevelInfo)
	Level logx.Leveler

	// ReplaceAttr is called to rewrite each non-group attribute before it is logged.
	// See https://pkg.go.dev/log/slog#HandlerOptions for details.
	ReplaceAttr func(groups []string, attr slog.Attr) slog.Attr

	// Time format (Default: time.StampMilli)
	TimeFormat string

	// Disable color (Default: false)
	NoColor bool
}

// NewHandler creates a [slog.Handler] that writes tinted logs to Writer w,
// using the default options. If opts is nil, the default options are used.
func New(w io.Writer, opts *Options) slog.Handler {
	h := &handler{
		w:          w,
		level:      defaultLevel,
		timeFormat: defaultTimeFormat,
	}
	if opts == nil {
		return h
	}

	h.addSource = opts.AddSource
	if opts.Level != nil {
		h.level = opts.Level
	}
	h.replaceAttr = opts.ReplaceAttr
	if opts.TimeFormat != "" {
		h.timeFormat = opts.TimeFormat
	}
	h.noColor = opts.NoColor
	h.palette = newPalette(h.noColor)
	return h
}

// handler implements a [slog.Handler].
type handler struct {
	attrsPrefix string
	groupPrefix string
	groups      []string
	palette     *palette

	mu sync.Mutex
	w  io.Writer

	addSource   bool
	level       logx.Leveler
	replaceAttr func([]string, slog.Attr) slog.Attr
	timeFormat  string
	noColor     bool
}

func (h *handler) clone() *handler {
	return &handler{
		attrsPrefix: h.attrsPrefix,
		groupPrefix: h.groupPrefix,
		groups:      h.groups,
		palette:     h.palette,
		w:           h.w,
		addSource:   h.addSource,
		level:       h.level,
		replaceAttr: h.replaceAttr,
		timeFormat:  h.timeFormat,
		noColor:     h.noColor,
	}
}

func (h *handler) Enabled(_ context.Context, level logx.Level) bool {
	return level >= h.level.Level()
}

func (h *handler) Handle(_ context.Context, r slog.Record) error {
	// get a buffer from the sync pool
	buf := newBuffer()
	defer buf.Free()

	rep := h.replaceAttr

	// write time
	if !r.Time.IsZero() {
		val := r.Time.Round(0) // strip monotonic to match Attr behavior
		if rep == nil {
			h.appendTime(buf, r.Time)
			buf.WriteByte(' ')
		} else if a := rep(nil /* groups */, slog.Time(slog.TimeKey, val)); a.Key != "" {
			if a.Value.Kind() == slog.KindTime {
				h.appendTime(buf, a.Value.Time())
			} else {
				h.appendValue(buf, a.Value, false)
			}
			buf.WriteByte(' ')
		}
	}

	// write level
	if rep == nil {
		h.appendLevel(buf, r.Level)
		buf.WriteByte(' ')
	} else if a := rep(nil /* groups */, slog.Any(slog.LevelKey, r.Level)); a.Key != "" {
		h.appendValue(buf, a.Value, false)
		buf.WriteByte(' ')
	}

	// write source
	if h.addSource {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		if f.File != "" {
			src := &slog.Source{
				Function: f.Function,
				File:     f.File,
				Line:     f.Line,
			}

			if rep == nil {
				h.appendSource(buf, src)
				buf.WriteByte(' ')
			} else if a := rep(nil /* groups */, slog.Any(slog.SourceKey, src)); a.Key != "" {
				h.appendValue(buf, a.Value, false)
				buf.WriteByte(' ')
			}
		}
	}

	// write message
	if rep == nil {
		buf.WriteString(r.Message)
		buf.WriteByte(' ')
	} else if a := rep(nil /* groups */, slog.String(slog.MessageKey, r.Message)); a.Key != "" {
		h.appendValue(buf, a.Value, false)
		buf.WriteByte(' ')
	}

	// write handler attributes
	if len(h.attrsPrefix) > 0 {
		buf.WriteString(h.attrsPrefix)
	}

	// write attributes
	r.Attrs(func(attr slog.Attr) bool {
		h.appendAttr(buf, attr, h.groupPrefix, h.groups)
		return true
	})

	if len(*buf) == 0 {
		return nil
	}
	(*buf)[len(*buf)-1] = '\n' // replace last space with newline

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err := h.w.Write(*buf)
	return err
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := h.clone()

	buf := newBuffer()
	defer buf.Free()

	// write attributes to buffer
	for _, attr := range attrs {
		h.appendAttr(buf, attr, h.groupPrefix, h.groups)
	}
	h2.attrsPrefix = h.attrsPrefix + string(*buf)
	return h2
}

func (h *handler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := h.clone()
	h2.groupPrefix += name + "."
	h2.groups = append(h2.groups, name)
	return h2
}

func (h *handler) appendTime(buf *buffer, t time.Time) {
	h.palette.colorFaint.Fprint(buf, t.Format(h.timeFormat))
}

func (h *handler) appendLevel(buf *buffer, level logx.Level) {
	switch {
	case level < logx.LevelVerbose:
		appendLevelWithDelta(buf, "DBG", h.palette.colorFaint, level-logx.LevelDebug)

	case level < logx.LevelInfo:
		appendLevelWithDelta(buf, "VRB", h.palette.colorFaint, level-logx.LevelVerbose)

	case level < logx.LevelWarn:
		appendLevelWithDelta(buf, "INF", h.palette.colorHiGreen, level-logx.LevelInfo)

	case level < logx.LevelError:
		appendLevelWithDelta(buf, "WRN", h.palette.colorHiYellow, level-logx.LevelWarn)

	default:
		appendLevelWithDelta(buf, "ERR", h.palette.colorHiRed, level-logx.LevelError)
	}
}

func appendLevelWithDelta(buf *buffer, levelText string, color *color.Color, delta logx.Level) {
	if delta == 0 {
		buf.WriteString(color.Sprint(levelText))
		return
	}
	buf.WriteString(color.Sprint(levelText + "+" + strconv.Itoa(int(delta))))
}

func (h *handler) appendSource(buf *buffer, src *slog.Source) {
	dir, file := filepath.Split(src.File)
	s := filepath.Join(filepath.Base(dir), file) + ":" + strconv.Itoa(src.Line)
	buf.WriteString(h.palette.colorFaint.Sprint(s))
}

func (h *handler) appendAttr(buf *buffer, attr slog.Attr, groupsPrefix string, groups []string) {
	attr.Value = attr.Value.Resolve()
	if rep := h.replaceAttr; rep != nil && attr.Value.Kind() != slog.KindGroup {
		attr = rep(groups, attr)
		attr.Value = attr.Value.Resolve()
	}

	if attr.Equal(slog.Attr{}) {
		return
	}

	switch attr.Value.Kind() {
	case slog.KindAny:
		if err, ok := attr.Value.Any().(tintError); ok {
			h.appendTintError(buf, err, attr.Key, groupsPrefix)
			buf.WriteByte(' ')
			return
		}
	case slog.KindGroup:
		if attr.Key != "" {
			groupsPrefix += attr.Key + "."
			groups = append(groups, attr.Key)
		}
		for _, groupAttr := range attr.Value.Group() {
			h.appendAttr(buf, groupAttr, groupsPrefix, groups)
		}
		return
	}

	h.appendKey(buf, attr.Key, groupsPrefix)
	h.appendValue(buf, attr.Value, true)
	buf.WriteByte(' ')
}

func (h *handler) appendKey(buf *buffer, key, groups string) {
	s := groups + key
	if needsQuoting(s) {
		s = strconv.Quote(s)
	}
	h.palette.colorFaint.Fprint(buf, s+"=")
}

func (h *handler) appendValue(buf *buffer, v slog.Value, quote bool) {
	switch v.Kind() {
	case slog.KindString:
		appendString(buf, v.String(), quote)
	case slog.KindInt64:
		*buf = strconv.AppendInt(*buf, v.Int64(), 10)
	case slog.KindUint64:
		*buf = strconv.AppendUint(*buf, v.Uint64(), 10)
	case slog.KindFloat64:
		*buf = strconv.AppendFloat(*buf, v.Float64(), 'g', -1, 64)
	case slog.KindBool:
		*buf = strconv.AppendBool(*buf, v.Bool())
	case slog.KindDuration:
		appendString(buf, v.Duration().String(), quote)
	case slog.KindTime:
		appendString(buf, v.Time().String(), quote)
	case slog.KindAny:
		switch cv := v.Any().(type) {
		case logx.Level:
			h.appendLevel(buf, cv)
		case encoding.TextMarshaler:
			data, err := cv.MarshalText()
			if err != nil {
				break
			}
			appendString(buf, string(data), quote)
		case *slog.Source:
			h.appendSource(buf, cv)
		default:
			appendString(buf, fmt.Sprintf("%+v", v.Any()), quote)
		}
	}
}

func (h *handler) appendTintError(buf *buffer, err tintError, attrKey, groupsPrefix string) {
	s := groupsPrefix + attrKey
	if needsQuoting(s) {
		s = strconv.Quote(s)
	}
	h.palette.colorHiRedFaint.Fprint(buf, s+"=")

	s = err.Error()
	if needsQuoting(s) {
		s = strconv.Quote(s)
	}
	h.palette.colorHiRed.Fprint(buf, s)
}

func appendString(buf *buffer, s string, quote bool) {
	if quote && needsQuoting(s) {
		*buf = strconv.AppendQuote(*buf, s)
	} else {
		buf.WriteString(s)
	}
}

func needsQuoting(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, r := range s {
		if unicode.IsSpace(r) || r == '"' || r == '=' || !unicode.IsPrint(r) {
			return true
		}
	}
	return false
}

type tintError struct{ error }

// Err returns a tinted (colorized) [slog.Attr] that will be written in red color
// by the [tint.Handler]. When used with any other [slog.Handler], it behaves as
//
//	slog.Any("err", err)
func Err(err error) slog.Attr {
	if err != nil {
		err = tintError{err}
	}
	return slog.Any(errKey, err)
}
