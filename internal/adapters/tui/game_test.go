package tui

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/game"
)

func newTestGame(clock *fakeClock) gameScreen {
	cfg := game.DefaultConfig([]rune("fj"))
	cfg.Duration = 10 * time.Second
	return newGameScreen(cfg, rand.New(rand.NewPCG(3, 3)), clock.Now)
}

func advance(s screen, clock *fakeClock, d time.Duration) (screen, tea.Cmd) {
	var cmd tea.Cmd
	for elapsed := time.Duration(0); elapsed < d; elapsed += frameEvery {
		clock.tick(frameEvery)
		s, cmd = s.Update(frameMsg{})
		if s.(gameScreen).ended {
			return s, cmd
		}
	}
	return s, cmd
}

func TestGameShowsLivesScoreAndTime(t *testing.T) {
	view := newTestGame(newClock()).View()

	for _, want := range []string{"♥♥♥", "score 0", "0:10"} {
		if !strings.Contains(view, want) {
			t.Errorf("missing %q:\n%s", want, view)
		}
	}
}

func TestGameStartsTheFrameClock(t *testing.T) {
	if newTestGame(newClock()).Init() == nil {
		t.Fatal("the game should start ticking")
	}
}

func TestLettersFallIntoView(t *testing.T) {
	clock := newClock()
	s, _ := advance(newTestGame(clock), clock, 2*time.Second)

	g := s.(gameScreen)
	if len(g.game.Drops()) == 0 {
		t.Fatal("no letters fell")
	}
	if !strings.ContainsAny(s.View(), "fj") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestTypingALetterScores(t *testing.T) {
	clock := newClock()
	s, _ := advance(newTestGame(clock), clock, 2*time.Second)
	char := s.(gameScreen).game.Drops()[0].Char

	s, _ = s.Update(runes(string(char)))

	if !strings.Contains(s.View(), "score 10") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestGameEndsWithAScoredResult(t *testing.T) {
	clock := newClock()
	s := screen(newTestGame(clock))
	var cmd tea.Cmd
	for range 200 {
		s, cmd = advance(s, clock, frameEvery)
		if s.(gameScreen).ended {
			if done := message(cmd).(stepDoneMsg); !done.scored {
				t.Fatal("game result should be scored")
			}
			return
		}
		for _, d := range s.(gameScreen).game.Drops() {
			s, _ = s.Update(runes(string(d.Char)))
		}
	}
	t.Fatal("game never ended")
}

func TestGameCanBeLeft(t *testing.T) {
	_, cmd := newTestGame(newClock()).Update(special(tea.KeyEsc))

	if _, ok := message(cmd).(leaveLessonMsg); !ok {
		t.Fatalf("message = %#v, want leave", message(cmd))
	}
}

func TestGameTakesEveryLetterOfABurst(t *testing.T) {
	clock := newClock()
	s, _ := advance(newTestGame(clock), clock, 6*time.Second)
	var burst strings.Builder
	for _, d := range s.(gameScreen).game.Drops() {
		burst.WriteRune(d.Char)
	}

	s, _ = s.Update(runes(burst.String()))

	if left := len(s.(gameScreen).game.Drops()); left != 0 || burst.Len() < 2 {
		t.Fatalf("burst %q left %d drops", burst.String(), left)
	}
}
