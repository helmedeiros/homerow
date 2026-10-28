package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type screen interface {
	Update(tea.Msg) (screen, tea.Cmd)
	View() string
}

type startLessonMsg struct{ number int }

type stepDoneMsg struct {
	result scoring.Result
	scored bool
}

type leaveLessonMsg struct{}

func send(msg tea.Msg) tea.Cmd {
	return func() tea.Msg { return msg }
}
