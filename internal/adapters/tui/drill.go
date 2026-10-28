package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/drill"
	"github.com/helmedeiros/homerow/internal/domain/keyboard"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

const drillWidth = 48

type drillScreen struct {
	session *drill.Session
	limit   time.Duration
	now     func() time.Time
}

func newDrillScreen(text string, limit time.Duration, now func() time.Time) drillScreen {
	return drillScreen{session: drill.NewSession(text), limit: limit, now: now}
}

func (d drillScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return d, nil
	}
	switch key.Type {
	case tea.KeyEsc:
		return d, send(leaveLessonMsg{})
	case tea.KeyCtrlC:
		return d, tea.Quit
	case tea.KeySpace:
		return d.typed(' ')
	case tea.KeyRunes:
		if len(key.Runes) == 1 {
			return d.typed(key.Runes[0])
		}
	}
	return d, nil
}

func (d drillScreen) typed(r rune) (screen, tea.Cmd) {
	d.session.Type(r, d.now())
	if d.session.Done() {
		return d, send(d.finished())
	}
	return d, nil
}

func (d drillScreen) finished() stepDoneMsg {
	return stepDoneMsg{
		result: scoring.Result{Accuracy: d.session.Accuracy(), WPM: d.session.WPM()},
		scored: true,
	}
}

func (d drillScreen) View() string {
	var b strings.Builder
	b.WriteString(d.renderText() + "\n\n")
	next, more := d.session.Expected()
	info, _ := keyboard.Lookup(next)
	b.WriteString(renderHands(info.Finger, more) + "\n\n")
	fmt.Fprintf(&b, "%s   %s\n\n",
		mutedStyle.Render(fmt.Sprintf("speed %.0f wpm", d.session.WPM())),
		mutedStyle.Render(fmt.Sprintf("accuracy %.0f%%", d.session.Accuracy())))
	b.WriteString(mutedStyle.Render("esc lessons"))
	return b.String()
}

func (d drillScreen) renderText() string {
	text := []rune(d.session.Text())
	marks := d.session.Marks()
	cursor := d.session.Position()
	var lines []string
	for _, span := range wrap(text, drillWidth) {
		var b strings.Builder
		for i := span[0]; i < span[1]; i++ {
			b.WriteString(renderChar(text[i], marks[i], i == cursor))
		}
		lines = append(lines, b.String())
	}
	return strings.Join(lines, "\n")
}

func renderChar(r rune, mark drill.Mark, cursor bool) string {
	shown := string(r)
	switch {
	case cursor:
		if r == ' ' {
			shown = "_"
		}
		return cursorStyle.Render(shown)
	case mark == drill.Correct:
		return correctStyle.Render(shown)
	case mark == drill.Wrong:
		return wrongStyle.Render(shown)
	default:
		return pendingStyle.Render(shown)
	}
}

func wrap(text []rune, width int) [][2]int {
	var spans [][2]int
	start, lastSpace := 0, -1
	for i, r := range text {
		if r == ' ' {
			lastSpace = i
		}
		if i-start >= width && lastSpace > start {
			spans = append(spans, [2]int{start, lastSpace + 1})
			start = lastSpace + 1
		}
	}
	return append(spans, [2]int{start, len(text)})
}
