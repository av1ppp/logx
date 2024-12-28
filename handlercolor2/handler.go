package handlercolor2

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

// New creates a new handler that writes logs to the specified io.Writer,
// using the provided tint.Options for configuration. If opts is nil,
// default options are used. This handler formats logs with colorized output.
func New(w io.Writer, opts *tint.Options) slog.Handler {
	return tint.NewHandler(w, opts)
}
