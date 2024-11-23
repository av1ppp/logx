package logx

import "log/slog"

// NewTextHandler creates a [TextHandler] that writes to w,
// using the given options.
// If opts is nil, the default options are used.
var NewTextHandler = slog.NewTextHandler
