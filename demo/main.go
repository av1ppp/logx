package main

import (
	"os"
	"time"

	"github.com/MatusOllah/slogcolor"
	"github.com/av1ppp/logx"
	"github.com/fatih/color"
)

func main() {
	if err := innerMain(); err != nil {
		panic(err)
	}
}

func innerMain() error {
	logger := logx.New(logx.NewColorText1Handler(os.Stderr, &logx.ColorText1HandlerOptions{
		Level:         logx.LevelDebug,
		TimeFormat:    time.DateTime,
		SrcFileMode:   logx.ShortFile,
		SrcFileLength: 0,
		MsgPrefix:     color.HiWhiteString("| "),
		MsgLength:     0,
		MsgColor:      color.New(),
		NoColor:       false,
	}))
	logger.Verbose("hello world")
	logger.Info(slogcolor.Prefix("SceneController", "switching scene"), "scene", "MainMenuScene")
	return nil
}
