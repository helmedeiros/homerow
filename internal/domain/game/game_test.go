package game_test

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func config() game.Config {
	return game.Config{
		Keys:       []rune("fjdk"),
		Width:      20,
		Height:     10,
		Lives:      3,
		Duration:   time.Minute,
		SpawnEvery: 2 * time.Second,
		FallEvery:  500 * time.Millisecond,
	}
}

func newGame() *game.Game {
	return game.New(config(), rand.New(rand.NewPCG(5, 5)))
}

func TestNewGameHasAnEmptySky(t *testing.T) {
	g := newGame()

	if len(g.Drops()) != 0 || g.Lives() != 3 || g.Over() {
		t.Fatalf("drops = %v, lives = %d, over = %v", g.Drops(), g.Lives(), g.Over())
	}
}

func TestLettersSpawnAtTheTopOnSchedule(t *testing.T) {
	g := newGame()

	g.Advance(1999 * time.Millisecond)
	if len(g.Drops()) != 0 {
		t.Fatalf("spawned too early: %v", g.Drops())
	}
	g.Advance(time.Millisecond)

	drops := g.Drops()
	if len(drops) != 1 {
		t.Fatalf("drops = %v, want one", drops)
	}
	d := drops[0]
	if d.Row != 0 || d.Column < 0 || d.Column >= 20 || !strings.ContainsRune("fjdk", d.Char) {
		t.Fatalf("drop = %+v", d)
	}
}

func TestLettersFallOneRowPerFallInterval(t *testing.T) {
	g := newGame()
	g.Advance(2 * time.Second)

	g.Advance(1500 * time.Millisecond)

	if row := g.Drops()[0].Row; row != 3 {
		t.Fatalf("row = %d, want 3", row)
	}
}

func TestLargeStepsBehaveLikeManySmallOnes(t *testing.T) {
	big, small := newGame(), newGame()

	big.Advance(7 * time.Second)
	for range 70 {
		small.Advance(100 * time.Millisecond)
	}

	if len(big.Drops()) != len(small.Drops()) {
		t.Fatalf("big = %v, small = %v", big.Drops(), small.Drops())
	}
	for i := range big.Drops() {
		if big.Drops()[i] != small.Drops()[i] {
			t.Fatalf("big = %v, small = %v", big.Drops(), small.Drops())
		}
	}
}
