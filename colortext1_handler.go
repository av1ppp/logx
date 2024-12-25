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

type ColorText1HandlerOptions = slogcolor.Options

var ColorText1HandlerDefaultOptions = slogcolor.DefaultOptions

type ColorText1Handler = slogcolor.Handler

func NewColorText1Handler(w io.Writer, opts *slogcolor.Options) *ColorText1Handler {
	return slogcolor.NewHandler(w, opts)
}
