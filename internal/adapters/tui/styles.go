package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorMuted   = lipgloss.Color("244")
	colorAccent  = lipgloss.Color("#3B82F6")
	colorCorrect = lipgloss.Color("#22C55E")
	colorStar    = lipgloss.Color("#FACC15")
	colorWrong   = lipgloss.Color("#EF4444")

	titleStyle     = lipgloss.NewStyle().Bold(true)
	mutedStyle     = lipgloss.NewStyle().Foreground(colorMuted)
	accentStyle    = lipgloss.NewStyle().Foreground(colorAccent).Bold(true)
	correctStyle   = lipgloss.NewStyle().Foreground(colorCorrect)
	wrongTextStyle = lipgloss.NewStyle().Foreground(colorWrong)
	wrongStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(colorWrong)
	cursorStyle    = lipgloss.NewStyle().Underline(true).Bold(true).Foreground(colorAccent)
	pendingStyle   = lipgloss.NewStyle()
	starStyle      = lipgloss.NewStyle().Foreground(colorStar)
	keyStyle       = lipgloss.NewStyle().Foreground(colorMuted)
	pressStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(colorAccent).Bold(true)
)
