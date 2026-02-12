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

	colorSource *color.Color
}

func newPalette(noColor bool) *palette {
	return &palette{
		colorTime: newColor(noColor, color.Faint),

		colorDebug:   newColor(noColor, color.FgHiCyan),
		colorVerbose: newColor(noColor, color.FgHiCyan),
		colorInfo:    newColor(noColor, color.FgHiGreen),
		colorWarn:    newColor(noColor, color.FgHiYellow),
		colorError:   newColor(noColor, color.FgHiRed),

		colorFgCyan: newColor(noColor, color.FgCyan),
		colorFgRed:  newColor(noColor, color.FgRed),

		colorSource: newColor(noColor, color.Faint),
	}
}

func newColor(noColor bool, value ...color.Attribute) *color.Color {
	if noColor {
		return color.New()
	}
	return color.New(value...)
}

var colorPrefix = color.New(color.BgHiWhite, color.FgBlack)
