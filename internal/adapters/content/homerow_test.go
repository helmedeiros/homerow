package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func TestHomeRowLessonsAreValidAndNumberedInOrder(t *testing.T) {
	m := content.HomeRow()

	for i, l := range m.Lessons {
		if l.Number != i+1 {
			t.Errorf("lesson at %d is numbered %d", i, l.Number)
		}
		if err := l.Validate(); err != nil {
			t.Error(err)
		}
	}
}

func TestHomeRowStartsWithFAndJ(t *testing.T) {
	m := content.HomeRow()
	want := []struct {
		number int
		kind   curriculum.Kind
		keys   string
	}{
		{1, curriculum.Intro, ""},
		{2, curriculum.NewKeys, "fj"},
		{3, curriculum.NewKeys, " "},
		{4, curriculum.Review, ""},
		{5, curriculum.NewKeys, "dk"},
		{8, curriculum.Play, ""},
		{9, curriculum.NewKeys, "sl"},
		{12, curriculum.NewKeys, "a;"},
	}
	for _, w := range want {
		l, ok := m.Lesson(w.number)
		if !ok {
			t.Fatalf("lesson %d missing", w.number)
		}
		if l.Kind != w.kind || string(l.NewKeys) != w.keys {
			t.Errorf("lesson %d = %v %q, want %v %q", w.number, l.Kind, string(l.NewKeys), w.kind, w.keys)
		}
	}
}

func TestHomeRowHasTwentyThreeLessonsEndingInAGame(t *testing.T) {
	m := content.HomeRow()

	if len(m.Lessons) != 23 {
		t.Fatalf("lessons = %d, want 23", len(m.Lessons))
	}
	last := m.Lessons[22]
	if last.Kind != curriculum.Play || string(last.Keys) != "asdfghjkl;" {
		t.Fatalf("lesson 23 = %v over %q", last.Kind, string(last.Keys))
	}
}

func TestHandDrillsStayOnOneHand(t *testing.T) {
	m := content.HomeRow()
	for number, hand := range map[int]keyboard.Hand{17: keyboard.Left, 18: keyboard.Right} {
		l, _ := m.Lesson(number)
		for _, r := range l.Keys {
			key, _ := keyboard.Lookup(r)
			if key.Finger.Hand() != hand {
				t.Errorf("lesson %d uses %q from the %v hand", number, r, key.Finger.Hand())
			}
		}
	}
}

func TestEveryDrillUsesOnlyItsLessonKeys(t *testing.T) {
	for _, l := range content.HomeRow().Lessons {
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

func TestCatalogServesTheHomeRow(t *testing.T) {
	if got := (content.Catalog{}).Module(); got.Title != "Home Row" || len(got.Lessons) != 23 {
		t.Fatalf("module = %q with %d lessons", got.Title, len(got.Lessons))
	}
}
