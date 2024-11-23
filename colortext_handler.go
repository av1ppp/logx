package logx

import (
	"io"

	"github.com/MatusOllah/slogcolor"
)

type SourceFileMode = slogcolor.SourceFileMode

const (
	// Nop does nothing.
	Nop = slogcolor.Nop

	// ShortFile produces only the filename.
	ShortFile = slogcolor.ShortFile

	// LongFile produces the full file path.
	LongFile = slogcolor.LongFile
)

type ColorTextHandlerOptions = slogcolor.Options

type ColorTextHandler = slogcolor.Handler

func NewColorTextHandler(w io.Writer, opts *slogcolor.Options) *ColorTextHandler {
	return slogcolor.NewHandler(w, opts)
}
