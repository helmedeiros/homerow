package game_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func TestALetterThatLandsCostsALife(t *testing.T) {
	g := newGame()

	g.Advance(2*time.Second + 10*500*time.Millisecond)

	if g.Lives() != 2 || g.Landed() != 1 {
		t.Fatalf("lives = %d, landed = %d", g.Lives(), g.Landed())
	}
	for _, d := range g.Drops() {
		if d.Row >= 10 {
			t.Fatalf("landed drop still on screen: %+v", d)
		}
	}
}

func TestRoundEndsWhenLivesRunOut(t *testing.T) {
	g := newGame()

	g.Advance(30 * time.Second)

	if !g.Over() || g.Lives() != 0 {
		t.Fatalf("over = %v, lives = %d", g.Over(), g.Lives())
	}
	elapsed := g.Elapsed()
	g.Advance(time.Minute)
	if g.Elapsed() != elapsed {
		t.Fatal("time kept passing after the round ended")
	}
	if g.Press('f') {
		t.Fatal("presses should be ignored after the round ended")
	}
}

func TestRoundEndsWhenTimeIsUp(t *testing.T) {
	cfg := config()
	cfg.Duration = 10 * time.Second
	g := game.New(cfg, rand.New(rand.NewPCG(5, 5)))

	for !g.Over() {
		g.Advance(100 * time.Millisecond)
		for _, d := range g.Drops() {
			g.Press(d.Char)
		}
	}

	if g.Elapsed() != 10*time.Second || g.Lives() != 3 || g.Remaining() != 0 {
		t.Fatalf("elapsed = %v, lives = %d, remaining = %v", g.Elapsed(), g.Lives(), g.Remaining())
	}
}

func TestResultCountsHitsAgainstEverythingThatWentWrong(t *testing.T) {
	g := newGame()
	g.Advance(2 * time.Second)
	g.Press(g.Drops()[0].Char)
	g.Press('z')
	g.Advance(58 * time.Second)

	r := g.Result()

	if r.Accuracy <= 0 || r.Accuracy >= 50 {
		t.Fatalf("accuracy = %v, want one hit among several misses and landings", r.Accuracy)
	}
	want := 1.0 / 5 / g.Elapsed().Minutes()
	if r.WPM < want-0.001 || r.WPM > want+0.001 {
		t.Fatalf("wpm = %v, want %v", r.WPM, want)
	}
}

func TestPerfectRoundIsFullyAccurate(t *testing.T) {
	g := newGame()
	if r := g.Result(); r.Accuracy != 100 || r.WPM != 0 {
		t.Fatalf("empty round result = %+v", r)
	}
}
