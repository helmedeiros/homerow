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
	started time.Time
	tick    ticker
}

func newDrillScreen(text string, limit time.Duration, now func() time.Time) drillScreen {
	return drillScreen{session: drill.NewSession(text), limit: limit, now: now, tick: realTicker}
}

func (d drillScreen) Init() tea.Cmd {
	if d.limit <= 0 {
		return nil
	}
	return d.tick(time.Second, tickMsg{})
}

func (d drillScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	if _, ok := msg.(tickMsg); ok {
		return d.ticked()
	}
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return d, nil
	}
	switch key.Type {
	case tea.KeyEsc:
		return d, send(leaveLessonMsg{})
	case tea.KeyCtrlC:
		return d, tea.Quit
	case tea.KeyBackspace:
		d.session.Back()
		return d, nil
	case tea.KeySpace:
		return d.typed(' ')
	case tea.KeyTab:
		return d.typed('\t')
	case tea.KeyRunes:
		return d.typed(key.Runes...)
	}
	return d, nil
}

func (d drillScreen) ticked() (screen, tea.Cmd) {
	if d.limit <= 0 {
		return d, nil
	}
	if d.timeLeft() <= 0 {
		return d, send(d.finished())
	}
	return d, d.tick(time.Second, tickMsg{})
}

func (d drillScreen) timeLeft() time.Duration {
	if d.started.IsZero() {
		return d.limit
	}
	return d.limit - d.now().Sub(d.started)
}

func (d drillScreen) typed(burst ...rune) (screen, tea.Cmd) {
	if d.started.IsZero() {
		d.started = d.now()
	}
	for _, r := range burst {
		d.session.Type(r, d.now())
		if d.session.Done() {
			return d, send(d.finished())
		}
	}
	return d, nil
}

func (d drillScreen) finished() stepDoneMsg {
	return stepDoneMsg{
		result: scoring.Result{Accuracy: d.session.Accuracy(), WPM: d.session.WPM(), Mistakes: d.session.Mistakes()},
		scored: true,
	}
}

func (d drillScreen) View() string {
	var b strings.Builder
	b.WriteString(d.renderText() + "\n\n")
	next, more := d.session.Expected()
	info, _ := keyboard.Lookup(next)
	b.WriteString(renderHands(info, more) + "\n\n")
	stats := fmt.Sprintf("speed %.0f wpm   accuracy %.0f%%", d.session.WPM(), d.session.Accuracy())
	if d.limit > 0 {
		stats += "   time left " + clockFace(d.timeLeft())
	}
	b.WriteString(mutedStyle.Render(stats) + "\n\n")
	b.WriteString(mutedStyle.Render("backspace fix   esc lessons"))
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
	if r == '\t' {
		shown = "→"
	}
	switch {
	case cursor:
		if r == ' ' {
			shown = "_"
		}
		return cursorStyle.Render(shown)
	case mark == drill.Correct:
		return correctStyle.Render(shown)
	case mark == drill.Corrected:
		return correctedStyle.Render(shown)
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

type tickMsg struct{}

type ticker func(time.Duration, tea.Msg) tea.Cmd

func realTicker(d time.Duration, msg tea.Msg) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return msg })
}

func clockFace(d time.Duration) string {
	seconds := max(int(d.Seconds()), 0)
	return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
}
