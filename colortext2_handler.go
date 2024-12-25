package logx

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

func NewColorText2Handler(w io.Writer, opts *tint.Options) slog.Handler {
	return tint.NewHandler(w, opts)
}
