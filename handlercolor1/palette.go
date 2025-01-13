package handlercolor1

import "github.com/fatih/color"

type palette struct {
	colorTime *color.Color

	colorDebug   *color.Color
	colorVerbose *color.Color
	colorInfo    *color.Color
	colorWarn    *color.Color
	colorError   *color.Color

	colorFgCyan *color.Color
	colorFgRed  *color.Color
}

func newPalette(noColor bool) *palette {
	return &palette{
		colorTime: newColor(noColor, color.Faint),

		colorDebug:   newColor(noColor, color.BgCyan, color.FgHiWhite),
		colorVerbose: newColor(noColor, color.BgCyan, color.FgHiWhite),
		colorInfo:    newColor(noColor, color.BgGreen, color.FgHiWhite),
		colorWarn:    newColor(noColor, color.BgYellow, color.FgHiWhite),
		colorError:   newColor(noColor, color.BgRed, color.FgHiWhite),

		colorFgCyan: newColor(noColor, color.FgCyan),
		colorFgRed:  newColor(noColor, color.FgRed),
	}
}

func newColor(noColor bool, value ...color.Attribute) *color.Color {
	if noColor {
		return color.New()
	}
	return color.New(value...)
}

var colorPrefix = color.New(color.BgHiWhite, color.FgBlack)
