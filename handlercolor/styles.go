package handlercolor

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	styleTime = lipgloss.NewStyle()

	styleLevelDebug = lipgloss.NewStyle().
			SetString("DBG").
			Bold(true).
			Foreground(lipgloss.Color("63"))

	styleLevelVerbose = lipgloss.NewStyle().
				SetString("VRB").
				Bold(true).
				Foreground(lipgloss.Color("63"))

	styleLevelInfo = lipgloss.NewStyle().
			SetString("INF").
			Bold(true).
			Foreground(lipgloss.Color("86"))

	styleLevelWarn = lipgloss.NewStyle().
			SetString("WRN").
			Bold(true).
			Foreground(lipgloss.Color("192"))

	styleLevelError = lipgloss.NewStyle().
			SetString("ERR").
			Bold(true).
			Foreground(lipgloss.Color("204"))

	styleLevelPanic = lipgloss.NewStyle().
			SetString("PNC").
			Bold(true).
			Foreground(lipgloss.Color("134"))

	styleUnknown = lipgloss.NewStyle().
			SetString(" ? ").
			Bold(true).
			Foreground(lipgloss.Color("134"))

	styleSource = lipgloss.NewStyle().
			Faint(true)

	stylePrefix = lipgloss.NewStyle().
			Bold(true).
			Faint(true)

	styleAttrKey = lipgloss.NewStyle().
			Faint(true)

	styleAttrValue = lipgloss.NewStyle().
			Foreground(lipgloss.Color("195")).
			Bold(true)

	styleCause = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))
)
