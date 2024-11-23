package logx

import (
	"context"
	"log/slog"
)

type Logger struct {
	*slog.Logger
}

func New(h Handler) *Logger {
	return &Logger{slog.New(h)}
}

// NewLogLogger returns a new [log.Logger] such that each call to its Output method
// dispatches a Record to the specified handler. The logger acts as a bridge from
// the older log API to newer structured logging handlers.
var NewLogLogger = slog.NewLogLogger

// NewRecord creates a [Record] from the given arguments.
// Use [Record.AddAttrs] to add attributes to the Record.
//
// NewRecord is intended for logging APIs that want to support a [Handler] as
// a backend.
var NewRecord = slog.NewRecord

func (self *Logger) With(args ...any) *Logger {
	if len(args) == 0 {
		return self
	}
	return &Logger{self.Logger.With(args...)}
}

func (self *Logger) WithGroup(name string) *Logger {
	if name == "" {
		return self
	}
	return &Logger{self.Logger.WithGroup(name)}
}

func (self *Logger) Verbose(msg string, args ...any) {
	self.Logger.Log(context.Background(), LevelVerbose, msg, args...)
}

func (self *Logger) VerboseContext(ctx context.Context, msg string, args ...any) {
	self.Logger.Log(ctx, LevelVerbose, msg, args...)
}

func (self *Logger) Panic(msg string, args ...any) {
	self.Logger.Log(context.Background(), LevelPanic, msg, args...)
	panic(msg + ", see logs for details")
}

func (self *Logger) PanicContext(ctx context.Context, msg string, args ...any) {
	self.Logger.Log(ctx, LevelPanic, msg, args...)
	panic(msg + ", see logs for details")
}
