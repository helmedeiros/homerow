package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/game"
	"github.com/helmedeiros/homerow/internal/domain/keyboard"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func TestCourseIsValid(t *testing.T) {
	if err := content.Course().Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestCourseNumbersLessonsWithoutGaps(t *testing.T) {
	for i, l := range content.Course().Lessons() {
		if l.Number != i+1 {
			t.Fatalf("lesson at %d is numbered %d", i, l.Number)
		}
	}
}

func TestCourseNumbersModulesInOrder(t *testing.T) {
	for i, m := range content.Course().Modules {
		if m.Number != i+1 {
			t.Fatalf("module at %d is numbered %d", i, m.Number)
		}
	}
}

func TestCourseOpensWithTheHomeRow(t *testing.T) {
	first := content.Course().Modules[0]

	if first.Title != "Home Row" || len(first.Lessons) != 23 {
		t.Fatalf("first module = %q with %d lessons", first.Title, len(first.Lessons))
	}
}

func TestCatalogServesTheCourse(t *testing.T) {
	if got, want := len((content.Catalog{}).Course().Lessons()), len(content.Course().Lessons()); got != want {
		t.Fatalf("catalog serves %d lessons, want %d", got, want)
	}
}

func TestHandDrillsStayOnOneHand(t *testing.T) {
	for _, l := range content.Course().Lessons() {
		if l.Kind != curriculum.HandDrill {
			continue
		}
		first, _ := keyboard.Lookup(l.Keys[0])
		for _, r := range l.Keys {
			key, _ := keyboard.Lookup(r)
			if key.Finger.Hand() != first.Finger.Hand() {
				t.Errorf("lesson %d mixes hands with %q", l.Number, r)
			}
		}
	}
}

func TestEveryDrillUsesOnlyItsLessonKeys(t *testing.T) {
	for _, l := range content.Course().Lessons() {
		allowed := string(l.Keys) + " "
		for seed := uint64(0); seed < 50; seed++ {
			for _, step := range l.Steps(rand.New(rand.NewPCG(seed, seed))) {
				if step.Kind != curriculum.DrillStep {
					continue
				}
				if strings.Trim(step.Text, allowed) != "" || step.Text == "" {
					t.Fatalf("lesson %d seed %d drill %q leaves %q", l.Number, seed, step.Text, allowed)
				}
			}
		}
	}
}

func TestAPerfectGameEarnsFiveStarsInEveryPlayLesson(t *testing.T) {
	for _, l := range content.Course().Lessons() {
		if l.Kind != curriculum.Play {
			continue
		}
		g := game.New(game.DefaultConfig(l.Keys), rand.New(rand.NewPCG(1, 1)))
		for !g.Over() {
			g.Advance(100 * time.Millisecond)
			for _, d := range g.Drops() {
				g.Press(d.Char)
			}
		}
		if stars := scoring.Stars(g.Result(), l.GoalWPM); stars != scoring.MaxStars {
			t.Errorf("lesson %d: perfect play earns %d stars (%.1f wpm, goal %.0f)", l.Number, stars, g.Result().WPM, l.GoalWPM)
		}
	}
}

func TestPracticeWordsAreRealWordsOnceThereAreEnough(t *testing.T) {
	for _, l := range content.Course().Lessons() {
		if !l.Words {
			continue
		}
		for _, step := range l.Steps(rand.New(rand.NewPCG(3, 3))) {
			if step.Kind == curriculum.DrillStep && len(strings.Fields(step.Text)) < 10 {
				t.Errorf("lesson %d drill %q", l.Number, step.Text)
			}
		}
	}
}
