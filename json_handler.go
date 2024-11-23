package logx

import "log/slog"

// NewJSONHandler creates a [JSONHandler] that writes to w,
// using the given options. If opts is nil, the default options are used.
var NewJSONHandler = slog.NewJSONHandler
