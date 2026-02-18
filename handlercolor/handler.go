package handlercolor

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/av1ppp/logx"
	"github.com/charmbracelet/lipgloss"
)

type Handler struct {
	groups []string
	attrs  []slog.Attr

	opts Options

	mu  sync.Mutex
	out io.Writer
}

// New creates a new Handler with the specified options. If opts is nil, uses [DefaultOptions].
func New(out io.Writer, opts *Options) *Handler {
	h := &Handler{
		mu:  sync.Mutex{},
		out: out,
	}

	if opts == nil {
		_opts := *DefaultOptions
		opts = &_opts
	}

	h.opts = *opts
	return h
}

func (h *Handler) clone() *Handler {
	return &Handler{
		groups: h.groups,
		attrs:  h.attrs,
		opts:   h.opts,
		// mu:     h.mu,
		out: h.out,
	}
}

// Enabled implements slog.Handler.Enabled .
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// Handle implements slog.Handler.Handle .
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	bf := getBuffer()
	defer freeBuffer(bf)

	bf.Reset()

	if !r.Time.IsZero() {
		bf.WriteString(styleTime.Render(r.Time.Format(h.opts.TimeFormat)))
		bf.WriteByte(' ')
	}

	switch r.Level {
	case logx.LevelDebug:
		bf.WriteString(styleLevelDebugRendered)
	case logx.LevelVerbose:
		bf.WriteString(styleLevelVerboseRendered)
	case logx.LevelInfo:
		bf.WriteString(styleLevelInfoRendered)
	case logx.LevelWarn:
		bf.WriteString(styleLevelWarnRendered)
	case logx.LevelError:
		bf.WriteString(styleLevelErrorRendered)
	case logx.LevelPanic:
		bf.WriteString(styleLevelPanicRendered)
	default:
		bf.WriteString(styleLevelUnknownRendered)
	}
	bf.WriteByte(' ')

	if h.opts.SrcFileMode != Nop {
		if r.PC != 0 {
			f, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()

			var filename string
			switch h.opts.SrcFileMode {
			case Nop:
				break
			case ShortFile:
				filename = filepath.Base(f.File)
			case LongFile:
				filename = f.File
			}
			lineStr := ":" + strconv.Itoa(f.Line)
			formatted := filename + lineStr + " "
			if h.opts.SrcFileLength > 0 {
				maxFilenameLen := h.opts.SrcFileLength - len(lineStr) - 1
				if len(filename) > maxFilenameLen {
					filename = filename[:maxFilenameLen] // Truncate if too long
				}
				lenStr := strconv.Itoa(h.opts.SrcFileLength)
				formatted = fmt.Sprintf("%-"+lenStr+"s", filename+lineStr)
			}
			bf.WriteString(styleSource.Render(formatted))
		}
	}

	//we need the attributes here, as we can print a longer string if there are no attributes
	var attrs []slog.Attr
	attrs = append(attrs, h.attrs...)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	bf.WriteString(h.opts.MsgPrefix)
	formattedMessage := r.Message
	if h.opts.MsgLength > 0 && len(attrs) > 0 {
		if len(formattedMessage) > h.opts.MsgLength {
			formattedMessage = formattedMessage[:h.opts.MsgLength-1] + "…" // Truncate and add ellipsis if too long
		} else {
			// Pad with spaces if too short
			lenStr := strconv.Itoa(h.opts.MsgLength)
			formattedMessage = fmt.Sprintf("%-"+lenStr+"s", formattedMessage)
		}
	}
	fmt.Fprintf(bf, "%s", formattedMessage)

	// * Gap between message and attributes
	// if len(attrs) > 0 {
	// 	bf.WriteString("   ")
	// }
	for _, a := range attrs {
		var styleKey, styleValue lipgloss.Style
		if a.Key == "cause" {
			styleKey = styleCause
			styleValue = styleCause
		} else {
			styleKey = styleAttrKey
			styleValue = styleAttrValue
		}

		bf.WriteByte(' ')
		if len(h.groups) > 0 {
			for _, g := range h.groups {
				bf.WriteString(styleKey.Render(g + "."))
			}
		}

		bf.WriteString(styleKey.Render(a.Key + "="))
		bf.WriteString(styleValue.Render(a.Value.String()))
	}

	bf.WriteByte('\n')

	h.mu.Lock()
	_, err := io.Copy(h.out, bf)
	h.mu.Unlock()

	return err
}

// WithGroup implements slog.Handler.WithGroup .
func (h *Handler) WithGroup(name string) slog.Handler {
	h2 := h.clone()
	h2.groups = append(h2.groups, name)
	return h2
}

// WithAttrs implements slog.Handler.WithAttrs .
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := h.clone()
	h2.attrs = append(h2.attrs, attrs...)
	return h2
}
