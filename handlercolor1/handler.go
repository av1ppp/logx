// ! Original repository: github.com/MatusOllah/slogcolor

package handlercolor1

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

type Options = slogcolor.Options

var DefaultOptions = slogcolor.DefaultOptions

type Handler = slogcolor.Handler

func New(w io.Writer, opts *slogcolor.Options) *Handler {
	return slogcolor.NewHandler(w, opts)
}
