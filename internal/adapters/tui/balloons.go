package tui

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

type balloonScreen struct {
	game  *game.Balloons
	cfg   game.BalloonConfig
	now   func() time.Time
	last  time.Time
	ended bool
	tick  ticker
}

func newBalloonScreen(cfg game.BalloonConfig, rng *rand.Rand, now func() time.Time) balloonScreen {
	return balloonScreen{game: game.NewBalloons(cfg, rng), cfg: cfg, now: now, last: now(), tick: realTicker}
}

func (b balloonScreen) Init() tea.Cmd {
	return b.tick(frameEvery, frameMsg{})
}

func (b balloonScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	switch msg := msg.(type) {
	case frameMsg:
		return b.advanced()
	case tea.KeyMsg:
		return b.pressed(msg)
	}
	return b, nil
}

func (b balloonScreen) advanced() (screen, tea.Cmd) {
	if b.ended {
		return b, nil
	}
	now := b.now()
	b.game.Advance(now.Sub(b.last))
	b.last = now
	if b.game.Over() {
		b.ended = true
		return b, send(stepDoneMsg{result: b.game.Result(), scored: true})
	}
	return b, b.tick(frameEvery, frameMsg{})
}

func (b balloonScreen) pressed(key tea.KeyMsg) (screen, tea.Cmd) {
	switch key.Type {
	case tea.KeyEsc:
		return b, send(leaveLessonMsg{})
	case tea.KeyCtrlC:
		return b, tea.Quit
	case tea.KeyRunes:
		for _, r := range key.Runes {
			b.game.Press(r)
		}
	}
	return b, nil
}

func (b balloonScreen) View() string {
	var v strings.Builder
	fmt.Fprintf(&v, "%s   score %d   time %s\n", hearts(b.game.Lives(), b.cfg.Lives), b.game.Score(), clockFace(b.game.Remaining()))
	v.WriteString(b.renderSky())
	v.WriteString("\n" + mutedStyle.Render("type the rising words before they float away · esc lessons"))
	return v.String()
}

func (b balloonScreen) renderSky() string {
	rows := emptySky(b.cfg.Width, b.cfg.Height)
	for _, balloon := range b.game.Balloons() {
		for i, r := range balloon.Word {
			style := accentStyle
			if i < balloon.Typed {
				style = correctStyle
			}
			if col := balloon.Column + i; col < b.cfg.Width {
				rows[balloon.Row][col] = style.Render(string(r))
			}
		}
	}
	return framed(rows, b.cfg.Width)
}
