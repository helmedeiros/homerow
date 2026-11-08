package game_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func TestEveryLetterTypedScores(t *testing.T) {
	b := newBalloons()
	b.Advance(2 * time.Second)
	word := b.Balloons()[0].Word

	typeWord(b, word)

	if b.Score() != 10*len(word) {
		t.Fatalf("score = %d after popping %q", b.Score(), word)
	}
}

func TestBalloonResultCountsWrongKeysAndEscapes(t *testing.T) {
	b := newBalloons()
	b.Advance(2 * time.Second)
	word := b.Balloons()[0].Word
	typeWord(b, word)
	b.Press('z')

	if r := b.Result(); r.Accuracy != float64(len(word))/float64(len(word)+1)*100 {
		t.Fatalf("accuracy = %v", r.Accuracy)
	}
	b.Advance(58 * time.Second)
	r := b.Result()
	if r.Accuracy >= 50 {
		t.Fatalf("accuracy = %v, want escapes to count against it", r.Accuracy)
	}
	want := float64(len(word)) / 5 / b.Elapsed().Minutes()
	if r.WPM < want-0.001 || r.WPM > want+0.001 {
		t.Fatalf("wpm = %v, want %v", r.WPM, want)
	}
}

func TestAnUntouchedBalloonRoundIsFullyAccurate(t *testing.T) {
	if r := newBalloons().Result(); r.Accuracy != 100 || r.WPM != 0 {
		t.Fatalf("result = %+v", r)
	}
}

func TestBalloonsRiseFasterAsTheRoundGoesOn(t *testing.T) {
	cfg := balloonConfig()
	cfg.SpawnEvery = 4 * time.Second
	cfg.Height = 1000
	cfg.RampEvery = 10 * time.Second
	cfg.Ramp = 0.5
	b := game.NewBalloons(cfg, rand.New(rand.NewPCG(1, 1)))

	if b.RiseInterval() != 500*time.Millisecond || b.SpawnInterval() != 4*time.Second {
		t.Fatalf("intervals = %v, %v at start", b.RiseInterval(), b.SpawnInterval())
	}
	b.Advance(10 * time.Second)
	if b.RiseInterval() != 250*time.Millisecond || b.SpawnInterval() != 2*time.Second {
		t.Fatalf("intervals = %v, %v after one ramp", b.RiseInterval(), b.SpawnInterval())
	}
	b.Advance(20 * time.Second)
	if b.RiseInterval() != game.FastestRise {
		t.Fatalf("RiseInterval() = %v, want the floor %v", b.RiseInterval(), game.FastestRise)
	}
}
