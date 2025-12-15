package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/logx/handlercolor1"
	"github.com/av1ppp/logx/handlercolor2"
	"github.com/av1ppp/logx/handlercolor3"
	"github.com/av1ppp/logx/handlerjson"
	"github.com/av1ppp/logx/handlertext"
)

func main() {
	demoHandlerColor1()
	fmt.Println()
	demoHandlerColor2()
	fmt.Println()
	demoHandlerColor3()
	fmt.Println()
	demoHandlerJson()
	fmt.Println()
	demoHandlerText()

}

func demoHandlerColor1() {
	logger := logx.New(handlercolor1.New(os.Stdout, &handlercolor1.Options{
		Level: logx.LevelDebug,
	}))

	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", logx.String("key", "value"))

	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", logx.String("key", "value"))

	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose info", logx.String("key", "value"))

	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn info", logx.String("key", "value"))

	logger.Error("error message")
	logger.Log(context.Background(), logx.LevelError+1, "error info", logx.String("key", "value"))
}

func demoHandlerColor2() {
	logger := logx.New(handlercolor2.New(os.Stdout, &handlercolor2.Options{
		Level: logx.LevelDebug,
	}))

	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", logx.String("key", "value"))

	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", logx.String("key", "value"))

	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose info", logx.String("key", "value"))

	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn info", logx.String("key", "value"))

	logger.Error("error message")
	logger.Log(context.Background(), logx.LevelError+1, "error info", logx.String("key", "value"))
}

func demoHandlerColor3() {
	logger := logx.New(handlercolor3.New(os.Stdout, &handlercolor3.Options{
		HandlerOptions: &slog.HandlerOptions{
			Level: logx.LevelDebug,
		},
	}))

	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", logx.String("key", "value"))

	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", logx.String("key", "value"))

	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose info", logx.String("key", "value"))

	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn info", logx.String("key", "value"))

	logger.Error("error message")
	logger.Log(context.Background(), logx.LevelError+1, "error info", logx.String("key", "value"))
}

func demoHandlerJson() {
	logger := logx.New(handlerjson.New(os.Stdout, &handlerjson.Options{
		Level: logx.LevelDebug,
	}))

	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", logx.String("key", "value"))

	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", logx.String("key", "value"))

	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose info", logx.String("key", "value"))

	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn info", logx.String("key", "value"))

	logger.Error("error message")
	logger.Log(context.Background(), logx.LevelError+1, "error info", logx.String("key", "value"))
}

func demoHandlerText() {
	logger := logx.New(handlertext.New(os.Stdout, &handlertext.Options{
		Level: logx.LevelDebug,
	}))

	logger.Debug("debug message")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message", logx.String("key", "value"))

	logger.Verbose("verbose message")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message", logx.String("key", "value"))

	logger.Info("info message")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose info", logx.String("key", "value"))

	logger.Warn("warn message")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn info", logx.String("key", "value"))

	logger.Error("error message")
	logger.Log(context.Background(), logx.LevelError+1, "error info", logx.String("key", "value"))
}
