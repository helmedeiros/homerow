package tui

import tea "github.com/charmbracelet/bubbletea"

type screen interface {
	Update(tea.Msg) (screen, tea.Cmd)
	View() string
}

type startLessonMsg struct{ number int }

func send(msg tea.Msg) tea.Cmd {
	return func() tea.Msg { return msg }
}
