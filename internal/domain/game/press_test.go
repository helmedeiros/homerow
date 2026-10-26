package game_test

import (
	"testing"
	"time"
)

func TestPressingALetterClearsTheLowestMatch(t *testing.T) {
	g := newGame()
	g.Advance(6 * time.Second)
	drops := g.Drops()
	lowest := drops[0]

	for _, d := range drops[1:] {
		if d.Char == lowest.Char && d.Row > lowest.Row {
			t.Fatal("expected the first drop to be the lowest")
		}
	}

	if !g.Press(lowest.Char) {
		t.Fatalf("Press(%q) = false", lowest.Char)
	}
	for _, d := range g.Drops() {
		if d == lowest {
			t.Fatalf("drop %+v still falling", lowest)
		}
	}
	if g.Hits() != 1 || g.Score() != 10 {
		t.Fatalf("hits = %d, score = %d", g.Hits(), g.Score())
	}
}

func TestPressingALetterNotFallingIsAMiss(t *testing.T) {
	g := newGame()
	g.Advance(2 * time.Second)
	before := len(g.Drops())

	if g.Press('z') {
		t.Fatal("Press('z') = true")
	}
	if g.Misses() != 1 || len(g.Drops()) != before {
		t.Fatalf("misses = %d, drops = %v", g.Misses(), g.Drops())
	}
}
