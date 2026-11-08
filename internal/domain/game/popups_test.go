package game_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func popUpConfig() game.PopUpConfig {
	return game.PopUpConfig{
		Keys:       []rune("fjdk"),
		Width:      20,
		Height:     10,
		Duration:   30 * time.Second,
		SpawnEvery: time.Second,
		ShowFor:    2500 * time.Millisecond,
	}
}

func newPopUps() *game.PopUps {
	return game.NewPopUps(popUpConfig(), rand.New(rand.NewPCG(5, 5)))
}

func TestKeysPopUpOnScheduleInsideTheBoard(t *testing.T) {
	p := newPopUps()

	p.Advance(999 * time.Millisecond)
	if len(p.Shown()) != 0 {
		t.Fatal("a key popped up too early")
	}
	p.Advance(time.Second + time.Millisecond)

	shown := p.Shown()
	if len(shown) != 2 {
		t.Fatalf("shown = %v", shown)
	}
	for _, pop := range shown {
		if pop.Column >= 20 || pop.Row >= 10 {
			t.Fatalf("pop-up outside the board: %+v", pop)
		}
	}
}

func TestAKeyHidesAgainUnlessTypedInTime(t *testing.T) {
	p := newPopUps()
	p.Advance(time.Second)

	p.Advance(2499 * time.Millisecond)
	if p.Hidden() != 0 {
		t.Fatal("the key hid too early")
	}
	p.Advance(time.Millisecond)

	if p.Hidden() != 1 || len(p.Shown()) != 2 {
		t.Fatalf("hidden = %d, shown = %v", p.Hidden(), p.Shown())
	}
}

func TestTypingAShownKeyCatchesTheOneAboutToHide(t *testing.T) {
	cfg := popUpConfig()
	cfg.Keys = []rune("f")
	p := game.NewPopUps(cfg, rand.New(rand.NewPCG(5, 5)))
	p.Advance(2 * time.Second)

	if !p.Press('f') {
		t.Fatal("f should catch a shown key")
	}

	shown := p.Shown()
	if len(shown) != 1 || shown[0].Hides != 4500*time.Millisecond || p.Hits() != 1 || p.Score() != 10 {
		t.Fatalf("shown = %v, hits = %d, score = %d", shown, p.Hits(), p.Score())
	}
}

func TestAKeyThatIsNotShownIsAMiss(t *testing.T) {
	p := newPopUps()
	p.Advance(time.Second)

	if p.Press('z') || len(p.Shown()) != 1 {
		t.Fatalf("z caught something: %v", p.Shown())
	}
	if r := p.Result(); r.Accuracy != 0 {
		t.Fatalf("accuracy = %v after one miss", r.Accuracy)
	}
}

func TestPopUpRoundEndsWhenTimeIsUp(t *testing.T) {
	p := newPopUps()

	for !p.Over() {
		p.Advance(100 * time.Millisecond)
		for _, pop := range p.Shown() {
			p.Press(pop.Char)
		}
	}

	if p.Remaining() != 0 || p.Hidden() != 0 || p.Hits() != 29 {
		t.Fatalf("remaining = %v, hidden = %d, hits = %d", p.Remaining(), p.Hidden(), p.Hits())
	}
	if r := p.Result(); r.Accuracy != 100 || r.WPM <= 0 {
		t.Fatalf("result = %+v", r)
	}
	elapsed := p.Elapsed()
	p.Advance(time.Second)
	if p.Elapsed() != elapsed || p.Press('f') {
		t.Fatal("the round kept going after it ended")
	}
}

func TestDefaultPopUpsUseTheGivenKeys(t *testing.T) {
	cfg := game.DefaultPopUps([]rune("fj"))

	if string(cfg.Keys) != "fj" || cfg.ShowFor <= 0 || cfg.Duration <= 0 {
		t.Fatalf("cfg = %+v", cfg)
	}
}
