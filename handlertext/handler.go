package handlertext

import "log/slog"

type Options = slog.HandlerOptions

// New creates a [TextHandler] that writes to w,
// using the given options.
// If opts is nil, the default options are used.
var New = slog.NewTextHandler
