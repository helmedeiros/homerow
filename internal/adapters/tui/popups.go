package tui

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

type popUpScreen struct {
	game  *game.PopUps
	cfg   game.PopUpConfig
	now   func() time.Time
	last  time.Time
	ended bool
	tick  ticker
}

func newPopUpScreen(cfg game.PopUpConfig, rng *rand.Rand, now func() time.Time) popUpScreen {
	return popUpScreen{game: game.NewPopUps(cfg, rng), cfg: cfg, now: now, last: now(), tick: realTicker}
}

func (p popUpScreen) Init() tea.Cmd {
	return p.tick(frameEvery, frameMsg{})
}

func (p popUpScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	switch msg := msg.(type) {
	case frameMsg:
		return p.advanced()
	case tea.KeyMsg:
		return p.pressed(msg)
	}
	return p, nil
}

func (p popUpScreen) advanced() (screen, tea.Cmd) {
	if p.ended {
		return p, nil
	}
	now := p.now()
	p.game.Advance(now.Sub(p.last))
	p.last = now
	if p.game.Over() {
		p.ended = true
		return p, send(stepDoneMsg{})
	}
	return p, p.tick(frameEvery, frameMsg{})
}

func (p popUpScreen) pressed(key tea.KeyMsg) (screen, tea.Cmd) {
	switch key.Type {
	case tea.KeyEsc:
		return p, send(leaveLessonMsg{})
	case tea.KeyCtrlC:
		return p, tea.Quit
	case tea.KeyRunes:
		for _, r := range key.Runes {
			p.game.Press(r)
		}
	}
	return p, nil
}

func (p popUpScreen) View() string {
	var v strings.Builder
	fmt.Fprintf(&v, "warm-up   caught %d   missed %d   time %s\n", p.game.Hits(), p.game.Hidden(), clockFace(p.game.Remaining()))
	rows := emptySky(p.cfg.Width, p.cfg.Height)
	for _, pop := range p.game.Shown() {
		rows[pop.Row][pop.Column] = accentStyle.Render(string(pop.Char))
	}
	v.WriteString(framed(rows, p.cfg.Width))
	v.WriteString("\n" + mutedStyle.Render("catch each key before it hides again · esc lessons"))
	return v.String()
}
