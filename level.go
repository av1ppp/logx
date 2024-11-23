package logx

import "log/slog"

type Level = slog.Level

const (
	LevelDebug   Level = -4
	LevelVerbose Level = -2
	LevelInfo    Level = 0
	LevelWarn    Level = 4
	LevelError   Level = 8
	LevelPanic   Level = 10
)
