package logx

import (
	"log/slog"

	multi "github.com/samber/slog-multi"
)

type Handler = slog.Handler

// JoinHandlers creates a Handler that writes to all handlers in the given list.
func JoinHandlers(handlers ...Handler) Handler {
	return multi.Fanout(handlers...)
}
