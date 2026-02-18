package handlercolor

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	styleTime = lipgloss.NewStyle()

	styleLevelDebugRendered = lipgloss.NewStyle().
				SetString("DBG").
				Bold(true).
				Foreground(lipgloss.Color("63")).
				Render()

	styleLevelVerboseRendered = lipgloss.NewStyle().
					SetString("VRB").
					Bold(true).
					Foreground(lipgloss.Color("63")).
					Render()

	styleLevelInfoRendered = lipgloss.NewStyle().
				SetString("INF").
				Bold(true).
				Foreground(lipgloss.Color("86")).
				Render()

	styleLevelWarnRendered = lipgloss.NewStyle().
				SetString("WRN").
				Bold(true).
				Foreground(lipgloss.Color("192")).
				Render()

	styleLevelErrorRendered = lipgloss.NewStyle().
				SetString("ERR").
				Bold(true).
				Foreground(lipgloss.Color("204")).
				Render()

	styleLevelPanicRendered = lipgloss.NewStyle().
				SetString("PNC").
				Bold(true).
				Foreground(lipgloss.Color("134")).
				Render()

	styleLevelUnknownRendered = lipgloss.NewStyle().
					SetString(" ? ").
					Bold(true).
					Foreground(lipgloss.Color("134")).
					Render()

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
