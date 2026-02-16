package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/logx/handlercolor"
	"github.com/av1ppp/logx/handlerdev"
	"github.com/av1ppp/logx/handlerjson"
	"github.com/av1ppp/logx/handlertext"
	"github.com/av1ppp/timex"
)

const timeFormat = time.Kitchen
const level = logx.LevelDebug
const addSource = true

func main() {
	demoHandlerText()

	fmt.Println()
	fmt.Println()

	demoHandlerJson()

	fmt.Println()
	fmt.Println()

	demoHandlerDev()

	fmt.Println()
	fmt.Println()

	demoHandlerColor()
}

func demoHandlerColor() {
	logger := logx.New(handlercolor.New(os.Stdout, &handlercolor.Options{
		Level:       level,
		SrcFileMode: handlercolor.ShortFile,
		TimeFormat:  timeFormat,
	}))
	demoLogger(logger)
}

func demoHandlerDev() {
	logger := logx.New(handlerdev.New(os.Stdout, &handlerdev.Options{
		HandlerOptions: &slog.HandlerOptions{
			Level:     level,
			AddSource: addSource,
		},
	}))
	demoLogger(logger)
}

func demoHandlerJson() {
	logger := logx.New(handlerjson.New(os.Stdout, &handlerjson.Options{
		Level:     level,
		AddSource: addSource,
	}))
	demoLogger(logger)
}

func demoHandlerText() {
	logger := logx.New(handlertext.New(os.Stdout, &handlertext.Options{
		Level:     level,
		AddSource: addSource,
	}))
	demoLogger(logger)
}

type Object struct {
	Name string
}

func demoLogger(logger *logx.Logger) {
	start := time.Now()

	obj := Object{Name: "name"}

	args := []any{
		logx.Int("attr.int", 42),
		logx.String("attr.str", "value"),
		logx.Any("attr.array", []string{"a", "b", "c"}),
		logx.Any("attr.pointer", &obj),
		logx.Any("attr.object", obj),
		logx.Time("attr.time", time.Now()),
		logx.Duration("attr.duration", time.Second),
		logx.Durationx("attr.durationx", timex.Second*5),
	}

	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", args...)
	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", args...)
	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose message", args...)
	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn message", args...)
	logger.Error("error message", logx.Cause(errors.New("something was wrong")))
	logger.Log(context.Background(), logx.LevelError+1, "error message", args...)

	logger = logger.WithGroup("group1")
	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", args...)
	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", args...)
	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose message", args...)
	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn message", args...)
	logger.Error("error message", logx.Cause(errors.New("something was wrong")))
	logger.Log(context.Background(), logx.LevelError+1, "error message", args...)

	logger = logger.WithGroup("group2")
	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", args...)
	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", args...)
	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose message", args...)
	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn message", args...)
	logger.Error("error message", logx.Cause(errors.New("something was wrong")))
	logger.Log(context.Background(), logx.LevelError+1, "error message", args...)

	fmt.Printf("duration: %s\n", time.Since(start))
}
