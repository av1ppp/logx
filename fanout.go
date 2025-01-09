package logx

import multi "github.com/samber/slog-multi"

// Fanout distributes records to multiple [Handler] in parallel
func Fanout(handlers ...Handler) Handler {
	return multi.Fanout(handlers...)
}
