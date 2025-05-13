package grpc

import (
	"context"
	"fmt"

	"github.com/av1ppp/logx"
)

type Logger struct {
	logger *logx.Logger
}

func NewLogger(logger *logx.Logger) *Logger {
	return &Logger{logger}
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (self *Logger) Info(args ...any) {
	self.logger.Info(fmt.Sprint(args...))
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (self *Logger) Infoln(args ...any) {
	self.logger.Info(fmt.Sprintln(args...))
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (self *Logger) Infof(format string, args ...any) {
	self.logger.Info(fmt.Sprintf(format, args...))
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (self *Logger) Warning(args ...any) {
	self.logger.Warn(fmt.Sprint(args...))
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (self *Logger) Warningln(args ...any) {
	self.logger.Warn(fmt.Sprintln(args...))
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (self *Logger) Warningf(format string, args ...any) {
	self.logger.Warn(fmt.Sprintf(format, args...))
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (self *Logger) Error(args ...any) {
	self.logger.Error(fmt.Sprint(args...))
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (self *Logger) Errorln(args ...any) {
	self.logger.Error(fmt.Sprintln(args...))
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (self *Logger) Errorf(format string, args ...any) {
	self.logger.Error(fmt.Sprintf(format, args...))
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (self *Logger) Fatal(args ...any) {
	self.logger.Panic(fmt.Sprint(args...))
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (self *Logger) Fatalln(args ...any) {
	self.logger.Panic(fmt.Sprintln(args...))
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (self *Logger) Fatalf(format string, args ...any) {
	self.logger.Panic(fmt.Sprintf(format, args...))
}

const (
	grpcLvlInfo int = iota
	grpcLvlWarn
	grpcLvlError
	grpcLvlFatal
)

// _grpcToLogxLevel maps gRPC log levels to logx log levels.
var _grpcToLogxLevel = map[int]logx.Level{
	grpcLvlInfo:  logx.LevelInfo,
	grpcLvlWarn:  logx.LevelWarn,
	grpcLvlError: logx.LevelError,
	grpcLvlFatal: logx.LevelPanic,
}

// V reports whether verbosity level l is at least the requested verbose level.
func (self *Logger) V(l int) bool {
	return self.logger.Enabled(context.Background(), _grpcToLogxLevel[l])
}
