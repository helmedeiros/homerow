package game

import (
	"math/rand/v2"
	"time"
)

type Config struct {
	Keys       []rune
	Width      int
	Height     int
	Lives      int
	Duration   time.Duration
	SpawnEvery time.Duration
	FallEvery  time.Duration
}

type Drop struct {
	Char   rune
	Column int
	Row    int
}

type Game struct {
	cfg       Config
	rng       *rand.Rand
	drops     []Drop
	lives     int
	elapsed   time.Duration
	nextSpawn time.Duration
	nextFall  time.Duration
}

func New(cfg Config, rng *rand.Rand) *Game {
	return &Game{
		cfg:       cfg,
		rng:       rng,
		lives:     cfg.Lives,
		nextSpawn: cfg.SpawnEvery,
		nextFall:  cfg.FallEvery,
	}
}

func (g *Game) Advance(dt time.Duration) {
	end := g.elapsed + dt
	for {
		next := min(g.nextSpawn, g.nextFall)
		if next > end {
			break
		}
		g.elapsed = next
		if g.nextFall == next {
			g.fall()
			g.nextFall += g.cfg.FallEvery
		}
		if g.nextSpawn == next {
			g.spawn()
			g.nextSpawn += g.cfg.SpawnEvery
		}
	}
	g.elapsed = end
}

func (g *Game) spawn() {
	g.drops = append(g.drops, Drop{
		Char:   g.cfg.Keys[g.rng.IntN(len(g.cfg.Keys))],
		Column: g.rng.IntN(g.cfg.Width),
	})
}

func (g *Game) fall() {
	for i := range g.drops {
		g.drops[i].Row++
	}
}

func (g *Game) Drops() []Drop { return append([]Drop(nil), g.drops...) }

func (g *Game) Lives() int { return g.lives }

func (g *Game) Over() bool { return false }
