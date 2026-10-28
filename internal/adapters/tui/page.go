package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const pageWidth = 60

type pageScreen struct {
	text string
}

func newPageScreen(text string) pageScreen {
	return pageScreen{text: text}
}

func (p pageScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return p, nil
	}
	switch key.String() {
	case "enter", " ", "right":
		return p, send(stepDoneMsg{})
	case "esc":
		return p, send(leaveLessonMsg{})
	case "ctrl+c":
		return p, tea.Quit
	}
	return p, nil
}

func (p pageScreen) View() string {
	body := lipgloss.NewStyle().Width(pageWidth).Render(p.text)
	return body + "\n\n" + mutedStyle.Render("enter continue · esc lessons")
}
