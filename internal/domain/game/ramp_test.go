package game_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func TestLettersFallFasterAsTheRoundGoesOn(t *testing.T) {
	cfg := config()
	cfg.SpawnEvery = time.Hour
	cfg.Height = 1000
	cfg.RampEvery = 10 * time.Second
	cfg.Ramp = 0.5
	g := game.New(cfg, rand.New(rand.NewPCG(1, 1)))

	if got := g.FallInterval(); got != 500*time.Millisecond {
		t.Fatalf("FallInterval() = %v at start", got)
	}
	g.Advance(10 * time.Second)
	if got := g.FallInterval(); got != 250*time.Millisecond {
		t.Fatalf("FallInterval() = %v after one ramp", got)
	}
	g.Advance(40 * time.Second)
	if got := g.FallInterval(); got != game.FastestFall {
		t.Fatalf("FallInterval() = %v, want the floor %v", got, game.FastestFall)
	}
}

func TestDefaultConfigIsPlayable(t *testing.T) {
	cfg := game.DefaultConfig([]rune("asdf"))

	if string(cfg.Keys) != "asdf" || cfg.Lives < 1 || cfg.Width < 10 || cfg.Height < 5 || cfg.Duration <= 0 {
		t.Fatalf("config = %+v", cfg)
	}
	if cfg.SpawnEvery <= 0 || cfg.FallEvery <= 0 || cfg.Ramp <= 0 || cfg.Ramp > 1 {
		t.Fatalf("config = %+v", cfg)
	}
}

func TestLettersSpawnFasterAsTheRoundGoesOn(t *testing.T) {
	cfg := config()
	cfg.RampEvery = 10 * time.Second
	cfg.Ramp = 0.5
	g := game.New(cfg, rand.New(rand.NewPCG(1, 1)))

	if got := g.SpawnInterval(); got != 2*time.Second {
		t.Fatalf("SpawnInterval() = %v at start", got)
	}
	for range 100 {
		g.Advance(100 * time.Millisecond)
		for _, d := range g.Drops() {
			g.Press(d.Char)
		}
	}
	if got := g.SpawnInterval(); got != time.Second {
		t.Fatalf("SpawnInterval() = %v after one ramp", got)
	}
}
