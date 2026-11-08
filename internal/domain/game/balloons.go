package game

import (
	"math/rand/v2"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type BalloonConfig struct {
	Words      []string
	Width      int
	Height     int
	Lives      int
	Duration   time.Duration
	SpawnEvery time.Duration
	RiseEvery  time.Duration
	RampEvery  time.Duration
	Ramp       float64
}

const FastestRise = 250 * time.Millisecond

func DefaultBalloons(words []string) BalloonConfig {
	return BalloonConfig{
		Words:      words,
		Width:      40,
		Height:     12,
		Lives:      3,
		Duration:   time.Minute,
		SpawnEvery: 2500 * time.Millisecond,
		RiseEvery:  900 * time.Millisecond,
		RampEvery:  15 * time.Second,
		Ramp:       0.85,
	}
}

type Balloon struct {
	Word   string
	Typed  int
	Column int
	Row    int
}

type Balloons struct {
	cfg       BalloonConfig
	rng       *rand.Rand
	balloons  []Balloon
	lives     int
	popped    int
	escaped   int
	hits      int
	misses    int
	elapsed   time.Duration
	nextSpawn time.Duration
	nextRise  time.Duration
}

func NewBalloons(cfg BalloonConfig, rng *rand.Rand) *Balloons {
	return &Balloons{
		cfg:       cfg,
		rng:       rng,
		lives:     cfg.Lives,
		nextSpawn: cfg.SpawnEvery,
		nextRise:  cfg.RiseEvery,
	}
}

func (b *Balloons) Advance(dt time.Duration) {
	if b.Over() {
		return
	}
	end := min(b.elapsed+dt, b.cfg.Duration)
	for !b.Over() {
		next := min(b.nextSpawn, b.nextRise)
		if next > end {
			break
		}
		b.elapsed = next
		if b.nextRise == next {
			b.rise()
			b.nextRise += b.RiseInterval()
		}
		if b.nextSpawn == next {
			b.spawn()
			b.nextSpawn += b.SpawnInterval()
		}
	}
	if !b.Over() {
		b.elapsed = end
	}
}

func (b *Balloons) RiseInterval() time.Duration {
	return max(ramped(b.cfg.RiseEvery, b.elapsed, b.cfg.RampEvery, b.cfg.Ramp), FastestRise)
}

func (b *Balloons) SpawnInterval() time.Duration {
	return ramped(b.cfg.SpawnEvery, b.elapsed, b.cfg.RampEvery, b.cfg.Ramp)
}

func (b *Balloons) spawn() {
	word := b.cfg.Words[b.rng.IntN(len(b.cfg.Words))]
	b.balloons = append(b.balloons, Balloon{
		Word:   word,
		Column: b.rng.IntN(max(b.cfg.Width-len(word)+1, 1)),
		Row:    b.cfg.Height - 1,
	})
}

func (b *Balloons) rise() {
	floating := b.balloons[:0]
	for _, balloon := range b.balloons {
		balloon.Row--
		if balloon.Row < 0 {
			b.escaped++
			b.lives = max(b.lives-1, 0)
			continue
		}
		floating = append(floating, balloon)
	}
	b.balloons = floating
}

func (b *Balloons) Press(r rune) bool {
	if b.Over() {
		return false
	}
	i := b.target(r)
	if i < 0 || rune(b.balloons[i].Word[b.balloons[i].Typed]) != r {
		b.misses++
		return false
	}
	b.hits++
	b.balloons[i].Typed++
	if b.balloons[i].Typed == len(b.balloons[i].Word) {
		b.balloons = append(b.balloons[:i], b.balloons[i+1:]...)
		b.popped++
	}
	return true
}

func (b *Balloons) target(r rune) int {
	highest := -1
	for i, balloon := range b.balloons {
		if balloon.Typed > 0 {
			return i
		}
		if rune(balloon.Word[0]) == r && (highest < 0 || balloon.Row < b.balloons[highest].Row) {
			highest = i
		}
	}
	return highest
}

func (b *Balloons) Balloons() []Balloon { return append([]Balloon(nil), b.balloons...) }

func (b *Balloons) Lives() int { return b.lives }

func (b *Balloons) Popped() int { return b.popped }

func (b *Balloons) Score() int { return b.hits * pointsPerHit }

func (b *Balloons) Result() scoring.Result {
	return result(b.hits, b.misses+b.escaped, b.elapsed)
}

func (b *Balloons) Escaped() int { return b.escaped }

func (b *Balloons) Elapsed() time.Duration { return b.elapsed }

func (b *Balloons) Remaining() time.Duration { return b.cfg.Duration - b.elapsed }

func (b *Balloons) Over() bool {
	return b.lives == 0 || b.elapsed >= b.cfg.Duration
}
