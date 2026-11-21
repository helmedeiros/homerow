package content_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
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
		{2, curriculum.NewKeys, "fj "},
		{3, curriculum.Practice, ""},
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

func TestTheFirstKeysLessonIntroducesTheSpaceThatSeparatesItsGroups(t *testing.T) {
	l, _ := content.HomeRow().Lesson(2)

	if l.Title != "Keys f, j and space" || string(l.NewKeys) != "fj " || string(l.Keys) != "fj " {
		t.Fatalf("lesson 2 = %q new %q keys %q", l.Title, string(l.NewKeys), string(l.Keys))
	}
}

func TestTheFirstHandDrillsKeepToTheKeysLearnedSoFar(t *testing.T) {
	m := content.HomeRow()
	for number, keys := range map[int]string{17: "asdf", 18: "jkl;"} {
		l, _ := m.Lesson(number)
		if l.Kind != curriculum.HandDrill || string(l.Keys) != keys {
			t.Errorf("lesson %d = %v over %q, want %q", number, l.Kind, string(l.Keys), keys)
		}
	}
}
