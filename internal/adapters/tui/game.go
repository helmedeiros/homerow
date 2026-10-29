package tui

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

const frameEvery = 100 * time.Millisecond

type frameMsg struct{}

type gameScreen struct {
	game  *game.Game
	cfg   game.Config
	now   func() time.Time
	last  time.Time
	ended bool
}

func newGameScreen(cfg game.Config, rng *rand.Rand, now func() time.Time) gameScreen {
	return gameScreen{game: game.New(cfg, rng), cfg: cfg, now: now, last: now()}
}

func (g gameScreen) Init() tea.Cmd {
	return frame()
}

func frame() tea.Cmd {
	return tea.Tick(frameEvery, func(time.Time) tea.Msg { return frameMsg{} })
}

func (g gameScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	switch msg := msg.(type) {
	case frameMsg:
		return g.advanced()
	case tea.KeyMsg:
		return g.pressed(msg)
	}
	return g, nil
}

func (g gameScreen) advanced() (screen, tea.Cmd) {
	if g.ended {
		return g, nil
	}
	now := g.now()
	g.game.Advance(now.Sub(g.last))
	g.last = now
	if g.game.Over() {
		g.ended = true
		return g, send(stepDoneMsg{result: g.game.Result(), scored: true})
	}
	return g, frame()
}

func (g gameScreen) pressed(key tea.KeyMsg) (screen, tea.Cmd) {
	switch key.Type {
	case tea.KeyEsc:
		return g, send(leaveLessonMsg{})
	case tea.KeyCtrlC:
		return g, tea.Quit
	case tea.KeyRunes:
		if len(key.Runes) == 1 {
			g.game.Press(key.Runes[0])
		}
	}
	return g, nil
}

func (g gameScreen) View() string {
	var b strings.Builder
	hearts := wrongTextStyle.Render(strings.Repeat("♥", g.game.Lives())) + mutedStyle.Render(strings.Repeat("♡", g.cfg.Lives-g.game.Lives()))
	fmt.Fprintf(&b, "%s   score %d   time %s\n", hearts, g.game.Score(), clockFace(g.game.Remaining()))
	b.WriteString(g.renderSky())
	b.WriteString("\n" + mutedStyle.Render("type the falling letters before they land · esc lessons"))
	return b.String()
}

func (g gameScreen) renderSky() string {
	rows := make([][]string, g.cfg.Height)
	for r := range rows {
		rows[r] = strings.Split(strings.Repeat(" ", g.cfg.Width), "")
	}
	for _, d := range g.game.Drops() {
		rows[d.Row][d.Column] = accentStyle.Render(string(d.Char))
	}
	var b strings.Builder
	b.WriteString("┌" + strings.Repeat("─", g.cfg.Width) + "┐\n")
	for _, row := range rows {
		b.WriteString("│" + strings.Join(row, "") + "│\n")
	}
	b.WriteString("└" + strings.Repeat("─", g.cfg.Width) + "┘")
	return b.String()
}
