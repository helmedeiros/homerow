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
