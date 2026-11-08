package game_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func balloonConfig() game.BalloonConfig {
	return game.BalloonConfig{
		Words:      []string{"sad", "lad", "ask"},
		Width:      20,
		Height:     10,
		Lives:      3,
		Duration:   time.Minute,
		SpawnEvery: 2 * time.Second,
		RiseEvery:  500 * time.Millisecond,
	}
}

func newBalloons() *game.Balloons {
	return game.NewBalloons(balloonConfig(), rand.New(rand.NewPCG(5, 5)))
}

func typeWord(b *game.Balloons, word string) {
	for _, r := range word {
		b.Press(r)
	}
}

func TestBalloonsStartWithAnEmptySky(t *testing.T) {
	b := newBalloons()

	if len(b.Balloons()) != 0 || b.Lives() != 3 || b.Over() {
		t.Fatalf("balloons = %v, lives = %d, over = %v", b.Balloons(), b.Lives(), b.Over())
	}
}

func TestBalloonsAppearAtTheBottomWithAWordThatFits(t *testing.T) {
	b := newBalloons()

	b.Advance(1999 * time.Millisecond)
	if len(b.Balloons()) != 0 {
		t.Fatal("a balloon appeared too early")
	}
	b.Advance(time.Millisecond)

	got := b.Balloons()
	if len(got) != 1 {
		t.Fatalf("balloons = %v", got)
	}
	balloon := got[0]
	if balloon.Row != 9 || balloon.Column+len(balloon.Word) > 20 || balloon.Typed != 0 {
		t.Fatalf("balloon = %+v", balloon)
	}
}

func TestBalloonsRiseOneRowAtATime(t *testing.T) {
	b := newBalloons()
	b.Advance(2 * time.Second)

	b.Advance(500 * time.Millisecond)

	if row := b.Balloons()[0].Row; row != 8 {
		t.Fatalf("row = %d, want 8", row)
	}
}

func TestABalloonThatEscapesCostsALife(t *testing.T) {
	b := newBalloons()

	b.Advance(2*time.Second + 10*500*time.Millisecond)

	if b.Lives() != 2 || b.Escaped() != 1 {
		t.Fatalf("lives = %d, escaped = %d", b.Lives(), b.Escaped())
	}
}

func TestTypingTheWholeWordPopsTheBalloon(t *testing.T) {
	b := newBalloons()
	b.Advance(2 * time.Second)
	word := b.Balloons()[0].Word

	typeWord(b, word[:2])
	if typed := b.Balloons()[0].Typed; typed != 2 {
		t.Fatalf("typed = %d, want 2", typed)
	}
	typeWord(b, word[2:])

	if len(b.Balloons()) != 0 || b.Popped() != 1 {
		t.Fatalf("balloons = %v, popped = %d", b.Balloons(), b.Popped())
	}
}

func TestTheBalloonClosestToTheTopIsTypedFirst(t *testing.T) {
	cfg := balloonConfig()
	cfg.Words = []string{"sad"}
	b := game.NewBalloons(cfg, rand.New(rand.NewPCG(5, 5)))
	b.Advance(4 * time.Second)

	b.Press('s')

	got := b.Balloons()
	if got[0].Typed != 1 || got[1].Typed != 0 {
		t.Fatalf("balloons = %+v", got)
	}
}

func TestAStartedWordKeepsTheFocus(t *testing.T) {
	cfg := balloonConfig()
	cfg.Words = []string{"sad"}
	b := game.NewBalloons(cfg, rand.New(rand.NewPCG(5, 5)))
	b.Advance(4 * time.Second)
	b.Press('s')

	if b.Press('s') {
		t.Fatal("a second s should not start the other balloon")
	}
	b.Press('a')

	got := b.Balloons()
	if got[0].Typed != 2 || got[1].Typed != 0 {
		t.Fatalf("balloons = %+v", got)
	}
}

func TestAWrongKeyIsIgnored(t *testing.T) {
	b := newBalloons()
	b.Advance(2 * time.Second)

	if b.Press('z') {
		t.Fatal("z should not match any balloon")
	}
	if b.Balloons()[0].Typed != 0 {
		t.Fatalf("balloon = %+v", b.Balloons()[0])
	}
}

func TestBalloonRoundEndsWhenLivesRunOut(t *testing.T) {
	b := newBalloons()

	b.Advance(30 * time.Second)

	if !b.Over() || b.Lives() != 0 {
		t.Fatalf("over = %v, lives = %d", b.Over(), b.Lives())
	}
	elapsed := b.Elapsed()
	b.Advance(time.Minute)
	if b.Elapsed() != elapsed || b.Press('s') {
		t.Fatal("the round kept going after it ended")
	}
}

func TestBalloonRoundEndsWhenTimeIsUp(t *testing.T) {
	cfg := balloonConfig()
	cfg.Duration = 10 * time.Second
	b := game.NewBalloons(cfg, rand.New(rand.NewPCG(5, 5)))

	for !b.Over() {
		b.Advance(100 * time.Millisecond)
		for _, balloon := range b.Balloons() {
			typeWord(b, balloon.Word[balloon.Typed:])
		}
	}

	if b.Remaining() != 0 || b.Lives() != 3 || b.Popped() != 4 {
		t.Fatalf("remaining = %v, lives = %d, popped = %d", b.Remaining(), b.Lives(), b.Popped())
	}
}

func TestDefaultBalloonsUseTheGivenWords(t *testing.T) {
	cfg := game.DefaultBalloons([]string{"sad"})

	if len(cfg.Words) != 1 || cfg.Lives != 3 || cfg.Duration != time.Minute {
		t.Fatalf("cfg = %+v", cfg)
	}
}
