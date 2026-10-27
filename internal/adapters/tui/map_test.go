package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func sampleOverview() app.Overview {
	return app.Overview{
		Module:     "Home Row",
		TotalStars: 4,
		Percent:    50,
		Lessons: []app.Card{
			{Number: 1, Title: "Welcome", Kind: curriculum.Intro, Entry: progress.Entry{Completed: true}},
			{Number: 2, Title: "Keys f and j", Kind: curriculum.NewKeys, Keys: "fj", Entry: progress.Entry{Completed: true}},
			{Number: 3, Title: "Review: f and j", Kind: curriculum.Review, Entry: progress.Entry{Completed: true, Stars: 4}},
			{Number: 4, Title: "Play: f j", Kind: curriculum.Play},
		},
	}
}

func TestMapShowsProgressAndStars(t *testing.T) {
	view := newMapScreen(sampleOverview()).View()

	for _, want := range []string{"Home Row", "50% progress", "4 stars", "Review: f and j", "★★★★☆", "✓"} {
		if !strings.Contains(view, want) {
			t.Errorf("map is missing %q:\n%s", want, view)
		}
	}
}

func TestMapStartsOnTheFirstUnfinishedLesson(t *testing.T) {
	m := newMapScreen(sampleOverview())

	if m.selected() != 4 {
		t.Fatalf("selected = %d, want 4", m.selected())
	}
}

func TestMapMovesAndStartsTheSelectedLesson(t *testing.T) {
	var s screen = newMapScreen(sampleOverview())

	s, _ = s.Update(runes("k"))
	s, _ = s.Update(special(tea.KeyUp))
	s, _ = s.Update(special(tea.KeyDown))
	_, cmd := s.Update(special(tea.KeyEnter))

	if got := message(cmd); got != (startLessonMsg{number: 3}) {
		t.Fatalf("message = %#v, want start lesson 3", got)
	}
}

func TestMapStaysWithinTheList(t *testing.T) {
	var s screen = newMapScreen(sampleOverview())

	s, _ = s.Update(runes("j"))
	if s.(mapScreen).selected() != 4 {
		t.Fatal("moved past the last lesson")
	}
	for range 10 {
		s, _ = s.Update(runes("k"))
	}
	if s.(mapScreen).selected() != 1 {
		t.Fatal("moved past the first lesson")
	}
}

func TestMapQuits(t *testing.T) {
	_, cmd := newMapScreen(sampleOverview()).Update(runes("q"))

	if _, ok := message(cmd).(tea.QuitMsg); !ok {
		t.Fatalf("message = %#v, want quit", message(cmd))
	}
}
