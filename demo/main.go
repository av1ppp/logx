package main

import (
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/logx/handlercolor1"
)

func main() {
	if err := innerMain(); err != nil {
		panic(err)
	}
}

func innerMain() error {
	logger := logx.New(handlercolor1.New(os.Stderr, &handlercolor1.Options{
		Level:         logx.LevelDebug,
		TimeFormat:    time.DateTime,
		SrcFileMode:   handlercolor1.ShortFile,
		SrcFileLength: 0,
		MsgPrefix:     color.HiWhiteString("| "),
		MsgLength:     0,
		MsgColor:      color.New(),
		NoColor:       false,
	}))

	logger.Debug("debug message")
	logger.Verbose("verbose message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	return nil
}
