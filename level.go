package logx

import (
	"log/slog"
	"strconv"
	"strings"
)

const (
	TimeKey    = slog.TimeKey
	LevelKey   = slog.LevelKey
	MessageKey = slog.MessageKey
	SourceKey  = slog.SourceKey
)

type Leveler = slog.Leveler

type Level = slog.Level

const (
	LevelDebug   Level = -4
	LevelVerbose Level = -2
	LevelInfo    Level = 0
	LevelWarn    Level = 4
	LevelError   Level = 8
	LevelPanic   Level = 10
)

func LevelStringWithDelta(level Level) string {
	switch {
	case level < LevelVerbose:
		return levelStringWithDelta("DEBUG", level-LevelDebug)

	case level < LevelInfo:
		return levelStringWithDelta("VERBOSE", level-LevelVerbose)

	case level < LevelWarn:
		return levelStringWithDelta("INFO", level-LevelInfo)

	case level < LevelError:
		return levelStringWithDelta("WARN", level-LevelWarn)

	case level < LevelPanic:
		return levelStringWithDelta("ERROR", level-LevelError)

	default:
		return levelStringWithDelta("PANIC", level-LevelPanic)
	}
}

func levelStringWithDelta(levelText string, delta Level) string {
	if delta == 0 {
		return levelText
	}
	return levelText + "+" + strconv.Itoa(int(delta))
}

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
		return 0, commonErrors.New("unknown log level")
	}
}

func MustParseLevel(s string) Level {
	level, err := ParseLevel(s)
	if err != nil {
		panic(err)
	}
	return level
}
