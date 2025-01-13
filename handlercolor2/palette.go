package handlercolor2

import "github.com/fatih/color"

type palette struct {
	colorFaint      *color.Color
	colorHiGreen    *color.Color
	colorHiYellow   *color.Color
	colorHiRed      *color.Color
	colorHiRedFaint *color.Color
}

func newPalette(noColor bool) *palette {
	return &palette{
		colorFaint:      newColor(noColor, color.Faint),
		colorHiGreen:    newColor(noColor, color.FgHiGreen),
		colorHiYellow:   newColor(noColor, color.FgHiYellow),
		colorHiRed:      newColor(noColor, color.FgHiRed),
		colorHiRedFaint: newColor(noColor, color.FgHiRed, color.Faint),
	}
}

func newColor(noColor bool, value ...color.Attribute) *color.Color {
	if noColor {
		return color.New()
	}
	return color.New(value...)
}

var colorPrefix = color.New(color.BgHiWhite, color.FgBlack)
