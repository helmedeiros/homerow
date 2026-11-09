package tui

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func newTestPopUps(clock *fakeClock) popUpScreen {
	cfg := game.DefaultPopUps([]rune("fj"))
	cfg.Duration = 10 * time.Second
	return newPopUpScreen(cfg, rand.New(rand.NewPCG(3, 3)), clock.Now)
}

func popFor(s screen, clock *fakeClock, d time.Duration) (screen, tea.Cmd) {
	var cmd tea.Cmd
	for elapsed := time.Duration(0); elapsed < d; elapsed += frameEvery {
		clock.tick(frameEvery)
		s, cmd = s.Update(frameMsg{})
		if s.(popUpScreen).ended {
			return s, cmd
		}
	}
	return s, cmd
}

func TestPopUpsShowCaughtMissedAndTime(t *testing.T) {
	s := newTestPopUps(newClock())
	if s.Init() == nil {
		t.Fatal("the warm-up should start ticking")
	}

	for _, want := range []string{"warm-up", "caught 0", "missed 0", "0:10", "hides again"} {
		if !strings.Contains(s.View(), want) {
			t.Errorf("missing %q:\n%s", want, s.View())
		}
	}
}

func TestKeysPopUpAndAreCaught(t *testing.T) {
	clock := newClock()
	s, _ := popFor(newTestPopUps(clock), clock, 1500*time.Millisecond)
	shown := s.(popUpScreen).game.Shown()
	if len(shown) != 1 || !strings.ContainsRune(s.View(), shown[0].Char) {
		t.Fatalf("view:\n%s", s.View())
	}

	s, _ = s.Update(runes(string(shown[0].Char)))

	if !strings.Contains(s.View(), "caught 1") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestTheWarmUpEndsUnscored(t *testing.T) {
	clock := newClock()
	s, cmd := popFor(newTestPopUps(clock), clock, 11*time.Second)

	if done, ok := message(cmd).(stepDoneMsg); !ok || done.scored {
		t.Fatalf("message = %#v", message(cmd))
	}
	if _, cmd = s.Update(frameMsg{}); cmd != nil {
		t.Fatal("an ended warm-up should stop ticking")
	}
}

func TestTheWarmUpCanBeLeftOrQuit(t *testing.T) {
	s := newTestPopUps(newClock())

	if _, cmd := s.Update(special(tea.KeyEsc)); message(cmd) != (leaveLessonMsg{}) {
		t.Fatalf("esc message = %#v", message(cmd))
	}
	if _, cmd := s.Update(special(tea.KeyCtrlC)); message(cmd) != tea.Quit() {
		t.Fatalf("ctrl+c message = %#v", message(cmd))
	}
	if _, cmd := s.Update(tea.WindowSizeMsg{}); cmd != nil {
		t.Fatal("other messages should be ignored")
	}
}
