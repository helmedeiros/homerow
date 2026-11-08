package tui

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func newTestBalloons(clock *fakeClock) balloonScreen {
	cfg := game.DefaultBalloons([]string{"sad", "lad"})
	cfg.Duration = 10 * time.Second
	return newBalloonScreen(cfg, rand.New(rand.NewPCG(3, 3)), clock.Now)
}

func rise(s screen, clock *fakeClock, d time.Duration) (screen, tea.Cmd) {
	var cmd tea.Cmd
	for elapsed := time.Duration(0); elapsed < d; elapsed += frameEvery {
		clock.tick(frameEvery)
		s, cmd = s.Update(frameMsg{})
		if s.(balloonScreen).ended {
			return s, cmd
		}
	}
	return s, cmd
}

func TestBalloonsShowLivesScoreAndTime(t *testing.T) {
	s := newTestBalloons(newClock())
	if s.Init() == nil {
		t.Fatal("the round should start ticking")
	}

	for _, want := range []string{"♥♥♥", "score 0", "0:10", "rising words"} {
		if !strings.Contains(s.View(), want) {
			t.Errorf("missing %q:\n%s", want, s.View())
		}
	}
}

func TestWordsRiseIntoView(t *testing.T) {
	clock := newClock()
	s, _ := rise(newTestBalloons(clock), clock, 3*time.Second)

	if !strings.Contains(s.View(), "ad") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestTypingAWordPopsItAndScores(t *testing.T) {
	clock := newClock()
	s, _ := rise(newTestBalloons(clock), clock, 3*time.Second)
	word := s.(balloonScreen).game.Balloons()[0].Word

	s, _ = s.Update(runes(word))

	if len(s.(balloonScreen).game.Balloons()) != 0 || !strings.Contains(s.View(), "score 30") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestBalloonRoundEndsWithAScoredResult(t *testing.T) {
	clock := newClock()
	s := screen(newTestBalloons(clock))
	var cmd tea.Cmd
	for range 200 {
		s, cmd = rise(s, clock, frameEvery)
		if s.(balloonScreen).ended {
			if done := message(cmd).(stepDoneMsg); !done.scored || done.result.Accuracy != 100 {
				t.Fatalf("result = %+v", done)
			}
			if _, cmd = s.Update(frameMsg{}); cmd != nil {
				t.Fatal("an ended round should stop ticking")
			}
			return
		}
		for _, b := range s.(balloonScreen).game.Balloons() {
			s, _ = s.Update(runes(b.Word[b.Typed:]))
		}
	}
	t.Fatal("round never ended")
}

func TestBalloonRoundCanBeLeftOrQuit(t *testing.T) {
	s := newTestBalloons(newClock())

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
