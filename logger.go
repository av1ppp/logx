package logx

import (
	"context"
	"log/slog"
	"runtime"
	"time"
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
	self.log(context.Background(), slog.Level(LevelVerbose), msg, args...)
}

func (self *Logger) VerboseContext(ctx context.Context, msg string, args ...any) {
	self.log(ctx, slog.Level(LevelVerbose), msg, args...)
}

func (self *Logger) Panic(msg string, args ...any) {
	self.log(context.Background(), slog.Level(LevelPanic), msg, args...)
	panic(msg + ", see logs for details")
}

func (self *Logger) PanicContext(ctx context.Context, msg string, args ...any) {
	self.log(ctx, slog.Level(LevelPanic), msg, args...)
	panic(msg + ", see logs for details")
}

// log is the method for fix stacktrace.
func (self *Logger) log(ctx context.Context, level Level, msg string, args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}
	if !self.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	// if !internal.IgnorePC {
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	// }
	r := NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	_ = self.Handler().Handle(ctx, r)
}
