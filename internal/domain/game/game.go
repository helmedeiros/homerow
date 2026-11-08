package game

import (
	"math/rand/v2"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type Config struct {
	Keys       []rune
	Width      int
	Height     int
	Lives      int
	Duration   time.Duration
	SpawnEvery time.Duration
	FallEvery  time.Duration
	RampEvery  time.Duration
	Ramp       float64
}

const FastestFall = 150 * time.Millisecond

func DefaultConfig(keys []rune) Config {
	return Config{
		Keys:       keys,
		Width:      40,
		Height:     12,
		Lives:      3,
		Duration:   time.Minute,
		SpawnEvery: 1600 * time.Millisecond,
		FallEvery:  600 * time.Millisecond,
		RampEvery:  15 * time.Second,
		Ramp:       0.85,
	}
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
	landed    int
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
	if g.Over() {
		return
	}
	end := min(g.elapsed+dt, g.cfg.Duration)
	for !g.Over() {
		next := min(g.nextSpawn, g.nextFall)
		if next > end {
			break
		}
		g.elapsed = next
		if g.nextFall == next {
			g.fall()
			g.nextFall += g.FallInterval()
		}
		if g.nextSpawn == next {
			g.spawn()
			g.nextSpawn += g.SpawnInterval()
		}
	}
	if !g.Over() {
		g.elapsed = end
	}
}

func (g *Game) FallInterval() time.Duration {
	return max(ramped(g.cfg.FallEvery, g.elapsed, g.cfg.RampEvery, g.cfg.Ramp), FastestFall)
}

func (g *Game) SpawnInterval() time.Duration {
	return ramped(g.cfg.SpawnEvery, g.elapsed, g.cfg.RampEvery, g.cfg.Ramp)
}

func ramped(interval, elapsed, every time.Duration, factor float64) time.Duration {
	if every <= 0 {
		return interval
	}
	for range int(elapsed / every) {
		interval = time.Duration(float64(interval) * factor)
	}
	return interval
}

func (g *Game) spawn() {
	g.drops = append(g.drops, Drop{
		Char:   g.cfg.Keys[g.rng.IntN(len(g.cfg.Keys))],
		Column: g.rng.IntN(g.cfg.Width),
	})
}

func (g *Game) fall() {
	falling := g.drops[:0]
	for _, d := range g.drops {
		d.Row++
		if d.Row >= g.cfg.Height {
			g.landed++
			g.lives = max(g.lives-1, 0)
			continue
		}
		falling = append(falling, d)
	}
	g.drops = falling
}

const pointsPerHit = 10

func (g *Game) Press(r rune) bool {
	if g.Over() {
		return false
	}
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

func (g *Game) Landed() int { return g.landed }

func (g *Game) Elapsed() time.Duration { return g.elapsed }

func (g *Game) Remaining() time.Duration { return g.cfg.Duration - g.elapsed }

func (g *Game) Over() bool {
	return g.lives == 0 || g.elapsed >= g.cfg.Duration
}

const charactersPerWord = 5

func (g *Game) Result() scoring.Result {
	return result(g.hits, g.misses+g.landed, g.elapsed)
}

func result(hits, failures int, elapsed time.Duration) scoring.Result {
	accuracy := 100.0
	if attempts := hits + failures; attempts > 0 {
		accuracy = float64(hits) / float64(attempts) * 100
	}
	wpm := 0.0
	if minutes := elapsed.Minutes(); minutes > 0 {
		wpm = float64(hits) / charactersPerWord / minutes
	}
	return scoring.Result{Accuracy: accuracy, WPM: wpm}
}
