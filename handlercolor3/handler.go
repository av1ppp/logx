package handlercolor3

import (
	"io"
	"log/slog"

	"github.com/golang-cz/devslog"
)

type Options = devslog.Options
type Handler = slog.Handler

func New(out io.Writer, opts *Options) Handler {
	return devslog.NewHandler(out, opts)
}
