package game

import (
	"math/rand/v2"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type PopUpConfig struct {
	Keys       []rune
	Width      int
	Height     int
	Duration   time.Duration
	SpawnEvery time.Duration
	ShowFor    time.Duration
}

func DefaultPopUps(keys []rune) PopUpConfig {
	return PopUpConfig{
		Keys:       keys,
		Width:      40,
		Height:     12,
		Duration:   45 * time.Second,
		SpawnEvery: 1200 * time.Millisecond,
		ShowFor:    2500 * time.Millisecond,
	}
}

type PopUp struct {
	Char   rune
	Column int
	Row    int
	Hides  time.Duration
}

type PopUps struct {
	cfg       PopUpConfig
	rng       *rand.Rand
	shown     []PopUp
	hits      int
	misses    int
	hidden    int
	elapsed   time.Duration
	nextSpawn time.Duration
}

func NewPopUps(cfg PopUpConfig, rng *rand.Rand) *PopUps {
	return &PopUps{cfg: cfg, rng: rng, nextSpawn: cfg.SpawnEvery}
}

func (p *PopUps) Advance(dt time.Duration) {
	if p.Over() {
		return
	}
	end := min(p.elapsed+dt, p.cfg.Duration)
	for p.nextSpawn <= end {
		p.hide(p.nextSpawn)
		p.spawn(p.nextSpawn)
		p.nextSpawn += p.cfg.SpawnEvery
	}
	p.hide(end)
	p.elapsed = end
}

func (p *PopUps) spawn(at time.Duration) {
	p.shown = append(p.shown, PopUp{
		Char:   p.cfg.Keys[p.rng.IntN(len(p.cfg.Keys))],
		Column: p.rng.IntN(p.cfg.Width),
		Row:    p.rng.IntN(p.cfg.Height),
		Hides:  at + p.cfg.ShowFor,
	})
}

func (p *PopUps) hide(at time.Duration) {
	visible := p.shown[:0]
	for _, pop := range p.shown {
		if pop.Hides <= at {
			p.hidden++
			continue
		}
		visible = append(visible, pop)
	}
	p.shown = visible
}

func (p *PopUps) Press(r rune) bool {
	if p.Over() {
		return false
	}
	soonest := -1
	for i, pop := range p.shown {
		if pop.Char == r && (soonest < 0 || pop.Hides < p.shown[soonest].Hides) {
			soonest = i
		}
	}
	if soonest < 0 {
		p.misses++
		return false
	}
	p.shown = append(p.shown[:soonest], p.shown[soonest+1:]...)
	p.hits++
	return true
}

func (p *PopUps) Shown() []PopUp { return append([]PopUp(nil), p.shown...) }

func (p *PopUps) Hits() int { return p.hits }

func (p *PopUps) Hidden() int { return p.hidden }

func (p *PopUps) Score() int { return p.hits * pointsPerHit }

func (p *PopUps) Elapsed() time.Duration { return p.elapsed }

func (p *PopUps) Remaining() time.Duration { return p.cfg.Duration - p.elapsed }

func (p *PopUps) Over() bool { return p.elapsed >= p.cfg.Duration }

func (p *PopUps) Result() scoring.Result {
	return result(p.hits, p.misses+p.hidden, p.elapsed)
}
