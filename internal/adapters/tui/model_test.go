package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/adapters/memstore"
	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

type harness struct {
	t     *testing.T
	model tea.Model
	clock *fakeClock
	store app.ProgressStore
}

func newHarness(t *testing.T) *harness {
	return newHarnessWith(t, memstore.New())
}

func newHarnessWith(t *testing.T, store app.ProgressStore) *harness {
	clock := newClock()
	trainer := app.NewTrainer(content.Catalog{}, store, func() uint64 { return 7 })
	m, err := New(trainer, Options{Now: clock.Now, Seed: func() uint64 { return 7 }, tick: manualTicks})
	if err != nil {
		t.Fatal(err)
	}
	return &harness{t: t, model: m, clock: clock, store: store}
}

func (h *harness) send(msg tea.Msg) {
	h.t.Helper()
	var cmd tea.Cmd
	h.model, cmd = h.model.Update(msg)
	h.run(cmd)
}

func (h *harness) run(cmd tea.Cmd) {
	if cmd == nil {
		return
	}
	msg := cmd()
	if batch, ok := msg.(tea.BatchMsg); ok {
		for _, c := range batch {
			h.run(c)
		}
		return
	}
	if msg != nil {
		h.send(msg)
	}
}

func manualTicks(time.Duration, tea.Msg) tea.Cmd { return nil }

func (h *harness) view() string { return h.model.View() }

func (h *harness) current() screen { return h.model.(Model).current }

func (h *harness) typeDrill() {
	h.t.Helper()
	d, ok := h.current().(drillScreen)
	if !ok {
		h.t.Fatalf("not on a drill:\n%s", h.view())
	}
	for _, r := range d.session.Text() {
		h.clock.tick(time.Second)
		if r == ' ' {
			h.send(special(tea.KeySpace))
		} else {
			h.send(runes(string(r)))
		}
	}
}

func TestOpensOnTheLessonMap(t *testing.T) {
	h := newHarness(t)

	if !strings.Contains(h.view(), "Welcome to touch typing") || !strings.Contains(h.view(), "0% progress") {
		t.Fatalf("view:\n%s", h.view())
	}
}

func TestIntroLessonRunsToResultsAndSavesProgress(t *testing.T) {
	h := newHarness(t)

	h.send(special(tea.KeyEnter))
	if !strings.Contains(h.view(), "Lesson 1") {
		t.Fatalf("view:\n%s", h.view())
	}
	for range 4 {
		h.send(special(tea.KeyEnter))
	}

	if !strings.Contains(h.view(), "GOOD JOB!") {
		t.Fatalf("view:\n%s", h.view())
	}
	board, _ := h.store.Load()
	if !board.Entry(1).Completed {
		t.Fatal("lesson 1 not saved")
	}
}

func TestNewKeyLessonWalksThroughIntroAndDrills(t *testing.T) {
	h := newHarness(t)
	h.send(startLessonMsg{number: 2})

	h.send(runes("f"))
	h.send(runes("j"))
	h.typeDrill()
	h.typeDrill()

	if !strings.Contains(h.view(), "GOOD JOB!") || !strings.Contains(h.view(), "next lesson") {
		t.Fatalf("view:\n%s", h.view())
	}
	h.send(special(tea.KeyEnter))
	if !strings.Contains(h.view(), "Lesson 3") || !strings.Contains(h.view(), "space bar") {
		t.Fatalf("view:\n%s", h.view())
	}
}

func TestReviewEarnsStarsShownOnTheMap(t *testing.T) {
	h := newHarness(t)
	h.send(startLessonMsg{number: 4})

	h.typeDrill()
	if !strings.Contains(h.view(), "★★★★★") {
		t.Fatalf("view:\n%s", h.view())
	}
	h.send(special(tea.KeyEsc))

	if !strings.Contains(h.view(), "5 stars") {
		t.Fatalf("view:\n%s", h.view())
	}
}

func TestLeavingALessonReturnsToTheMap(t *testing.T) {
	h := newHarness(t)
	h.send(startLessonMsg{number: 4})

	h.send(special(tea.KeyEsc))

	if _, ok := h.current().(mapScreen); !ok {
		t.Fatalf("view:\n%s", h.view())
	}
}

func TestPlayLessonRunsTheGame(t *testing.T) {
	h := newHarness(t)
	h.send(startLessonMsg{number: 8})

	for range 2000 {
		g, ok := h.current().(gameScreen)
		if !ok {
			break
		}
		for _, d := range g.game.Drops() {
			h.send(runes(string(d.Char)))
		}
		h.clock.tick(frameEvery)
		h.send(frameMsg{})
	}

	if !strings.Contains(h.view(), "GOOD JOB!") {
		t.Fatalf("view:\n%s", h.view())
	}
}

func TestUnknownLessonKeepsTheMapWithAnError(t *testing.T) {
	h := newHarness(t)

	h.send(startLessonMsg{number: 99})

	if _, ok := h.current().(mapScreen); !ok || !strings.Contains(h.view(), "unknown lesson") {
		t.Fatalf("view:\n%s", h.view())
	}
}

func TestTheMapKeepsTheLessonYouJustLeft(t *testing.T) {
	h := newHarness(t)
	h.send(startLessonMsg{number: 6})

	h.send(special(tea.KeyEsc))

	if got := h.current().(mapScreen).selected(); got != 6 {
		t.Fatalf("selected = %d, want 6", got)
	}
}

func TestTheMapMovesOnAfterAPassedLesson(t *testing.T) {
	h := newHarness(t)
	h.send(startLessonMsg{number: 4})
	h.typeDrill()

	h.send(special(tea.KeyEsc))

	if got := h.current().(mapScreen).selected(); got != 5 {
		t.Fatalf("selected = %d, want 5", got)
	}
}

func TestABalloonStepOpensTheBalloonGame(t *testing.T) {
	m := newHarness(t).model.(Model)

	s := m.screenFor(curriculum.Lesson{}, curriculum.Step{Kind: curriculum.GameStep, Game: curriculum.BalloonWords, Words: []string{"sad"}})

	if _, ok := s.(balloonScreen); !ok {
		t.Fatalf("screen = %T", s)
	}
}

func TestAPopUpStepOpensTheWarmUp(t *testing.T) {
	m := newHarness(t).model.(Model)

	s := m.screenFor(curriculum.Lesson{}, curriculum.Step{Kind: curriculum.GameStep, Game: curriculum.PopUpWarmUp, Keys: []rune("fj")})

	if _, ok := s.(popUpScreen); !ok {
		t.Fatalf("screen = %T", s)
	}
}
