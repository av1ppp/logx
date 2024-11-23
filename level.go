package logx

import (
	"errors"
	"log/slog"
	"strings"
)

type Level = slog.Level

const (
	LevelDebug   Level = -4
	LevelVerbose Level = -2
	LevelInfo    Level = 0
	LevelWarn    Level = 4
	LevelError   Level = 8
	LevelPanic   Level = 10
)

func ParseLevel(s string) (Level, error) {
	switch strings.ToLower(s) {
	case "debug":
		return LevelDebug, nil
	case "verbose":
		return LevelVerbose, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "panic":
		return LevelPanic, nil
	default:
		return 0, errors.New("unknown log level")
	}
}

func MustParseLevel(s string) Level {
	level, err := ParseLevel(s)
	if err != nil {
		panic(err)
	}
	return level
}
