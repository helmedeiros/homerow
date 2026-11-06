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

func (m mapScreen) focus(number int) mapScreen {
	for i, card := range m.overview.Lessons {
		if card.Number == number {
			m.cursor = i
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
	case "up":
		m.cursor = max(m.cursor-1, 0)
	case "down":
		m.cursor = min(m.cursor+1, len(m.overview.Lessons)-1)
	case "right":
		m = m.jump(1)
	case "left":
		m = m.jump(-1)
	case "enter":
		return m, send(startLessonMsg{number: m.selected()})
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	}
	if key.Type == tea.KeyRunes {
		m = m.moved(key.Runes)
	}
	return m, nil
}

func (m mapScreen) moved(burst []rune) mapScreen {
	for _, r := range burst {
		switch r {
		case 'k':
			m.cursor = max(m.cursor-1, 0)
		case 'j':
			m.cursor = min(m.cursor+1, len(m.overview.Lessons)-1)
		case 'l':
			m = m.jump(1)
		case 'h':
			m = m.jump(-1)
		}
	}
	return m
}

const mapRows = 18

func (m mapScreen) jump(delta int) mapScreen {
	modules := m.overview.Modules
	for i, section := range modules {
		if m.cursor < section.First || m.cursor > section.Last {
			continue
		}
		target := i + delta
		if target < 0 || target >= len(modules) {
			return m
		}
		m.cursor = m.firstUnfinished(modules[target])
		return m
	}
	return m
}

func (m mapScreen) firstUnfinished(section app.Section) int {
	for i := section.First; i <= section.Last; i++ {
		if !m.overview.Lessons[i].Entry.Completed {
			return i
		}
	}
	return section.First
}

func (m mapScreen) View() string {
	var b strings.Builder
	o := m.overview
	section := o.ModuleOf(m.cursor)
	fmt.Fprintf(&b, "%s  %s\n\n", titleStyle.Render("homerow"),
		mutedStyle.Render(fmt.Sprintf("%d%% progress · %d stars", o.Percent, o.TotalStars)))
	fmt.Fprintf(&b, "%s  %s\n\n", titleStyle.Render(fmt.Sprintf("Module %d · %s", section.Number, section.Title)),
		mutedStyle.Render(fmt.Sprintf("%d%% · %d stars", section.Percent, section.Stars)))
	top, bottom := m.window(section)
	if top > section.First {
		b.WriteString(mutedStyle.Render("     ↑ more") + "\n")
	}
	for i := top; i <= bottom; i++ {
		b.WriteString(m.row(i) + "\n")
	}
	if bottom < section.Last {
		b.WriteString(mutedStyle.Render("     ↓ more") + "\n")
	}
	b.WriteString("\n" + mutedStyle.Render("↑/↓ choose · ←/→ module · enter start · q quit"))
	return b.String()
}

func (m mapScreen) window(section app.Section) (int, int) {
	size := section.Last - section.First + 1
	if size <= mapRows {
		return section.First, section.Last
	}
	start := min(max(m.cursor-section.First-mapRows/2, 0), size-mapRows)
	return section.First + start, section.First + start + mapRows - 1
}

func (m mapScreen) row(i int) string {
	card := m.overview.Lessons[i]
	pointer := "  "
	title := fmt.Sprintf("%-26s", card.Title)
	if i == m.cursor {
		pointer = accentStyle.Render("› ")
		title = accentStyle.Render(title)
	}
	return fmt.Sprintf("%s%3d  %s %s", pointer, card.Number, title, status(card))
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
