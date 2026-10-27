package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorMuted  = lipgloss.Color("244")
	colorAccent = lipgloss.Color("#3B82F6")

	mutedStyle  = lipgloss.NewStyle().Foreground(colorMuted)
	accentStyle = lipgloss.NewStyle().Foreground(colorAccent).Bold(true)
	keyStyle    = lipgloss.NewStyle().Foreground(colorMuted)
	pressStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(colorAccent).Bold(true)
)
