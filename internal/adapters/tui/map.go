package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type mapScreen struct {
	overview app.Overview
	cursor   int
}

func newMapScreen(o app.Overview) mapScreen {
	m := mapScreen{overview: o}
	for i, card := range o.Lessons {
		if !card.Entry.Completed {
			m.cursor = i
			break
		}
	}
	return m
}

func (m mapScreen) selected() int {
	return m.overview.Lessons[m.cursor].Number
}

func (m mapScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "up", "k":
		m.cursor = max(m.cursor-1, 0)
	case "down", "j":
		m.cursor = min(m.cursor+1, len(m.overview.Lessons)-1)
	case "enter":
		return m, send(startLessonMsg{number: m.selected()})
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m mapScreen) View() string {
	var b strings.Builder
	o := m.overview
	fmt.Fprintf(&b, "%s  %s\n\n", titleStyle.Render("homerow · "+o.Module),
		mutedStyle.Render(fmt.Sprintf("%d%% progress · %d stars", o.Percent, o.TotalStars)))
	for i, card := range o.Lessons {
		pointer := "  "
		title := card.Title
		if i == m.cursor {
			pointer = accentStyle.Render("› ")
			title = accentStyle.Render(fmt.Sprintf("%-26s", title))
		} else {
			title = fmt.Sprintf("%-26s", title)
		}
		fmt.Fprintf(&b, "%s%3d  %s %s\n", pointer, card.Number, title, status(card))
	}
	b.WriteString("\n" + mutedStyle.Render("↑/↓ choose · enter start · q quit"))
	return b.String()
}

func status(card app.Card) string {
	switch {
	case card.Kind.Scored() && (card.Entry.Completed || card.Entry.Stars > 0):
		return renderStars(card.Entry.Stars)
	case card.Entry.Completed:
		return correctStyle.Render("✓")
	default:
		return mutedStyle.Render("·")
	}
}

func renderStars(n int) string {
	return starStyle.Render(strings.Repeat("★", n)) + mutedStyle.Render(strings.Repeat("☆", scoring.MaxStars-n))
}
