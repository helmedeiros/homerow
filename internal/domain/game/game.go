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
	hits      int
	misses    int
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

const pointsPerHit = 10

func (g *Game) Press(r rune) bool {
	lowest := -1
	for i, d := range g.drops {
		if d.Char == r && (lowest < 0 || d.Row > g.drops[lowest].Row) {
			lowest = i
		}
	}
	if lowest < 0 {
		g.misses++
		return false
	}
	g.drops = append(g.drops[:lowest], g.drops[lowest+1:]...)
	g.hits++
	return true
}

func (g *Game) Hits() int { return g.hits }

func (g *Game) Misses() int { return g.misses }

func (g *Game) Score() int { return g.hits * pointsPerHit }

func (g *Game) Drops() []Drop { return append([]Drop(nil), g.drops...) }

func (g *Game) Lives() int { return g.lives }

func (g *Game) Over() bool { return false }
