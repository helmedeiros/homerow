package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

type fakeClock struct{ now time.Time }

func (c *fakeClock) Now() time.Time { return c.now }

func (c *fakeClock) tick(d time.Duration) { c.now = c.now.Add(d) }

func newClock() *fakeClock {
	return &fakeClock{now: time.Date(2026, 10, 28, 20, 0, 0, 0, time.UTC)}
}

func typeText(s screen, clock *fakeClock, text string) (screen, tea.Cmd) {
	var cmd tea.Cmd
	for _, r := range text {
		clock.tick(time.Second)
		key := runes(string(r))
		if r == ' ' {
			key = special(tea.KeySpace)
		}
		s, cmd = s.Update(key)
	}
	return s, cmd
}

func TestDrillShowsTheTextAndTheFingerForTheNextKey(t *testing.T) {
	view := newDrillScreen("fj fj", 0, newClock().Now).View()

	if !strings.Contains(view, "fj fj") || !strings.Contains(view, "left index") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestDrillMovesTheFingerHintAlong(t *testing.T) {
	clock := newClock()
	s, _ := typeText(newDrillScreen("fj", 0, clock.Now), clock, "f")

	if !strings.Contains(s.View(), "right index") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestDrillShowsLiveSpeedAndAccuracy(t *testing.T) {
	clock := newClock()
	s, _ := typeText(newDrillScreen("fjfjf fjfjf", 0, clock.Now), clock, "fjfjd")

	view := s.View()
	if !strings.Contains(view, "80%") || !strings.Contains(view, "wpm") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestDrillFinishesWithTheResult(t *testing.T) {
	clock := newClock()
	_, cmd := typeText(newDrillScreen("fj fj", 0, clock.Now), clock, "fj fd")

	done, ok := message(cmd).(stepDoneMsg)
	if !ok {
		t.Fatalf("message = %#v, want step done", message(cmd))
	}
	if done.result.Accuracy != 80 || done.result.WPM <= 0 || !done.scored {
		t.Fatalf("result = %+v", done)
	}
}

func TestDrillCanBeLeft(t *testing.T) {
	_, cmd := newDrillScreen("fj", 0, newClock().Now).Update(special(tea.KeyEsc))

	if _, ok := message(cmd).(leaveLessonMsg); !ok {
		t.Fatalf("message = %#v, want leave", message(cmd))
	}
}

func TestBackspaceStepsBackToTheMistake(t *testing.T) {
	clock := newClock()
	s, _ := typeText(newDrillScreen("fjf", 0, clock.Now), clock, "fd")

	s, cmd := s.Update(special(tea.KeyBackspace))

	if cmd != nil || !strings.Contains(s.View(), "right index") {
		t.Fatalf("cmd = %#v, view:\n%s", message(cmd), s.View())
	}
}

func TestBackspaceAtTheStartDoesNothing(t *testing.T) {
	s, cmd := newDrillScreen("fj", 0, newClock().Now).Update(special(tea.KeyBackspace))

	if cmd != nil || !strings.Contains(s.View(), "left index") {
		t.Fatalf("cmd = %#v, view:\n%s", message(cmd), s.View())
	}
}

func TestFixedMistakesStillCostAccuracy(t *testing.T) {
	clock := newClock()
	s, _ := typeText(newDrillScreen("fjf", 0, clock.Now), clock, "fd")
	s, _ = s.Update(special(tea.KeyBackspace))
	_, cmd := typeText(s, clock, "jf")

	done, ok := message(cmd).(stepDoneMsg)
	if !ok || done.result.Accuracy != 75 {
		t.Fatalf("message = %#v, want 75%% accuracy", message(cmd))
	}
}

func TestFixedMistakesAreYellowAndOpenOnesRed(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	t.Cleanup(func() { lipgloss.SetColorProfile(termenv.Ascii) })

	if got := renderChar('f', drill.Corrected, false); !strings.Contains(got, "38;2;234;179;8") {
		t.Errorf("corrected = %q, want yellow", got)
	}
	if got := renderChar('f', drill.Wrong, false); !strings.Contains(got, "48;2;239;68;68") {
		t.Errorf("wrong = %q, want red", got)
	}
}

func TestDrillWrapsLongTextOnWords(t *testing.T) {
	text := strings.TrimSpace(strings.Repeat("asdf jkl; ", 12))
	view := newDrillScreen(text, 0, newClock().Now).View()

	for _, line := range strings.Split(view, "\n") {
		if len([]rune(line)) > drillWidth+4 && strings.Contains(line, "asdf") {
			t.Fatalf("line too long (%d): %q", len([]rune(line)), line)
		}
	}
}

func TestDrillTypesEveryCharacterOfABurst(t *testing.T) {
	_, cmd := newDrillScreen("fjf", 0, newClock().Now).Update(runes("fjf"))

	done, ok := message(cmd).(stepDoneMsg)
	if !ok || done.result.Accuracy != 100 {
		t.Fatalf("message = %#v, want a finished drill", message(cmd))
	}
}
