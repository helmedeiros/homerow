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
	tick  ticker
}

func newGameScreen(cfg game.Config, rng *rand.Rand, now func() time.Time) gameScreen {
	return gameScreen{game: game.New(cfg, rng), cfg: cfg, now: now, last: now(), tick: realTicker}
}

func (g gameScreen) Init() tea.Cmd {
	return g.tick(frameEvery, frameMsg{})
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
	return g, g.tick(frameEvery, frameMsg{})
}

func (g gameScreen) pressed(key tea.KeyMsg) (screen, tea.Cmd) {
	switch key.Type {
	case tea.KeyEsc:
		return g, send(leaveLessonMsg{})
	case tea.KeyCtrlC:
		return g, tea.Quit
	case tea.KeyRunes:
		for _, r := range key.Runes {
			g.game.Press(r)
		}
	}
	return g, nil
}

func (g gameScreen) View() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s   score %d   time %s\n", hearts(g.game.Lives(), g.cfg.Lives), g.game.Score(), clockFace(g.game.Remaining()))
	b.WriteString(g.renderSky())
	b.WriteString("\n" + mutedStyle.Render("type the falling letters before they land · esc lessons"))
	return b.String()
}

func (g gameScreen) renderSky() string {
	rows := emptySky(g.cfg.Width, g.cfg.Height)
	for _, d := range g.game.Drops() {
		rows[d.Row][d.Column] = accentStyle.Render(string(d.Char))
	}
	return framed(rows, g.cfg.Width)
}

func hearts(lives, total int) string {
	return wrongTextStyle.Render(strings.Repeat("♥", lives)) + mutedStyle.Render(strings.Repeat("♡", total-lives))
}

func emptySky(width, height int) [][]string {
	rows := make([][]string, height)
	for r := range rows {
		rows[r] = strings.Split(strings.Repeat(" ", width), "")
	}
	return rows
}

func framed(rows [][]string, width int) string {
	var b strings.Builder
	b.WriteString("┌" + strings.Repeat("─", width) + "┐\n")
	for _, row := range rows {
		b.WriteString("│" + strings.Join(row, "") + "│\n")
	}
	b.WriteString("└" + strings.Repeat("─", width) + "┘")
	return b.String()
}
