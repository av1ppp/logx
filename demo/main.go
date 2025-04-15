package main

import (
	"context"
	"errors"
	"os"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/logx/handlercolor1"
	"github.com/av1ppp/logx/handlercolor2"
)

func main() {
	if err := innerMain(); err != nil {
		panic(err)
	}
}

func innerMain() error {
	logger := logx.New(handlercolor1.New(os.Stdout, &handlercolor1.Options{
		Level: logx.LevelDebug,
	}))

	// logger := logx.New(handlercolor1.New(os.Stderr, &handlercolor1.Options{
	// 	Level:         logx.LevelDebug,
	// 	TimeFormat:    time.DateTime,
	// 	SrcFileMode:   handlercolor1.ShortFile,
	// 	SrcFileLength: 0,
	// 	MsgPrefix:     color.HiWhiteString("| "),
	// 	MsgLength:     0,
	// 	MsgColor:      color.New(),
	// 	NoColor:       false,
	// }))

	logger.Debug("debug message 1", handlercolor2.Err(errors.New("something was wrong")), "key", "value")
	logger.Log(context.Background(), logx.LevelDebug+1, "debug message 2")

	logger.Verbose("verbose message 1")
	logger.Log(context.Background(), logx.LevelVerbose+1, "verbose message 2")

	logger.Info("info message 1")
	logger.Log(context.Background(), logx.LevelInfo+1, "verbose info 2")

	logger.Warn("warn message 1")
	logger.Log(context.Background(), logx.LevelWarn+1, "warn info 2")

	logger.Error("error message 1")
	logger.Log(context.Background(), logx.LevelError+1, "error info 2")

	return nil
}
