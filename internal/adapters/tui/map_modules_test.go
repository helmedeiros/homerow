package tui

import (
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func twoModuleOverview() app.Overview {
	o := sampleOverview()
	o.Lessons = append(o.Lessons,
		app.Card{Number: 5, Title: "Keys r and u", Kind: curriculum.NewKeys, Keys: "ru"},
		app.Card{Number: 6, Title: "Review: r and u", Kind: curriculum.Review},
	)
	o.Modules = append(o.Modules, app.Section{Number: 2, Title: "Top Row", First: 4, Last: 5})
	return o
}

func longOverview(lessons int) app.Overview {
	o := app.Overview{Modules: []app.Section{{Number: 1, Title: "Long Row", First: 0, Last: lessons - 1}}}
	for i := range lessons {
		o.Lessons = append(o.Lessons, app.Card{Number: i + 1, Title: fmt.Sprintf("Lesson %d", i+1), Kind: curriculum.Intro, Entry: progress.Entry{Completed: i < lessons-1}})
	}
	return o
}

func TestMapShowsOnlyTheModuleOfTheSelectedLesson(t *testing.T) {
	view := newMapScreen(twoModuleOverview()).View()

	if !strings.Contains(view, "Module 1 · Home Row") || strings.Contains(view, "Keys r and u") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestMovingPastTheLastLessonOpensTheNextModule(t *testing.T) {
	s, _ := newMapScreen(twoModuleOverview()).Update(special(tea.KeyDown))

	view := s.View()
	if !strings.Contains(view, "Module 2 · Top Row") || !strings.Contains(view, "Keys r and u") || strings.Contains(view, "Welcome") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestModuleHeaderShowsItsOwnProgress(t *testing.T) {
	o := twoModuleOverview()
	o.Modules[1].Percent, o.Modules[1].Stars = 50, 3

	view := newMapScreen(o).focus(6).View()

	if !strings.Contains(view, "50% · 3 stars") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestLongModulesScrollToKeepTheSelectionVisible(t *testing.T) {
	m := newMapScreen(longOverview(60))

	view := m.View()
	if !strings.Contains(view, "Lesson 60") || strings.Contains(view, "Lesson 1 ") || !strings.Contains(view, "↑ more") {
		t.Fatalf("view:\n%s", view)
	}
	if lines := strings.Count(view, "Lesson "); lines != mapRows {
		t.Fatalf("%d lessons shown, want %d", lines, mapRows)
	}
}

func TestScrollingShowsWhatIsBelow(t *testing.T) {
	view := newMapScreen(longOverview(60)).focus(1).View()

	if !strings.Contains(view, "Lesson 1 ") || !strings.Contains(view, "↓ more") || strings.Contains(view, "↑ more") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestRightJumpsToTheNextModuleAndLeftComesBack(t *testing.T) {
	var s screen = newMapScreen(twoModuleOverview()).focus(2)

	s, _ = s.Update(special(tea.KeyRight))
	if got := s.(mapScreen).selected(); got != 5 {
		t.Fatalf("after right selected = %d, want 5", got)
	}
	s, _ = s.Update(special(tea.KeyLeft))
	if got := s.(mapScreen).selected(); got != 4 {
		t.Fatalf("after left selected = %d, want the first unfinished lesson 4", got)
	}
}

func TestModuleJumpsStopAtTheEnds(t *testing.T) {
	var s screen = newMapScreen(twoModuleOverview()).focus(2)

	s, _ = s.Update(special(tea.KeyLeft))
	if got := s.(mapScreen).selected(); got != 2 {
		t.Fatalf("left on the first module moved to %d", got)
	}
	s, _ = s.Update(runes("ll"))
	if got := s.(mapScreen).selected(); got != 5 {
		t.Fatalf("right past the last module moved to %d", got)
	}
	s, _ = s.Update(runes("h"))
	if got := s.(mapScreen).selected(); got != 4 {
		t.Fatalf("h moved to %d, want 4", got)
	}
}

func TestMapHintMentionsModules(t *testing.T) {
	if view := newMapScreen(twoModuleOverview()).View(); !strings.Contains(view, "←/→ module") {
		t.Fatalf("view:\n%s", view)
	}
}
