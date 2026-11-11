package curriculum_test

import (
	"errors"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestOnlyPractisingKindsAreScored(t *testing.T) {
	scored := map[curriculum.Kind]bool{
		curriculum.Intro:     false,
		curriculum.NewKeys:   false,
		curriculum.Review:    true,
		curriculum.Practice:  true,
		curriculum.HandDrill: true,
		curriculum.Play:      true,
		curriculum.Travel:    true,
		curriculum.Tab:       true,
	}
	for kind, want := range scored {
		if kind.Scored() != want {
			t.Errorf("%v.Scored() = %v, want %v", kind, kind.Scored(), want)
		}
	}
}

func TestKindNames(t *testing.T) {
	if curriculum.HandDrill.String() != "hand drill" || curriculum.Play.String() != "play" || curriculum.Travel.String() != "travel" {
		t.Fatalf("names = %q, %q", curriculum.HandDrill, curriculum.Play)
	}
}

func TestValidLessons(t *testing.T) {
	lessons := []curriculum.Lesson{
		{Number: 1, Title: "Welcome", Kind: curriculum.Intro, Pages: []string{"hi"}},
		{Number: 2, Title: "Keys f and j", Kind: curriculum.NewKeys, NewKeys: []rune("fj"), Keys: []rune("fj")},
		{Number: 4, Title: "Review", Kind: curriculum.Review, Keys: []rune("fj"), Focus: []rune("f"), GoalWPM: 8},
	}
	for _, l := range lessons {
		if err := l.Validate(); err != nil {
			t.Errorf("lesson %d: %v", l.Number, err)
		}
	}
}

func TestInvalidLessons(t *testing.T) {
	cases := map[string]curriculum.Lesson{
		"intro without pages":   {Number: 1, Title: "x", Kind: curriculum.Intro},
		"drill without keys":    {Number: 2, Title: "x", Kind: curriculum.Review, GoalWPM: 8},
		"new key not allowed":   {Number: 3, Title: "x", Kind: curriculum.NewKeys, NewKeys: []rune("d"), Keys: []rune("fj")},
		"focus key not allowed": {Number: 4, Title: "x", Kind: curriculum.Review, Keys: []rune("fj"), Focus: []rune("k"), GoalWPM: 8},
		"scored without goal":   {Number: 5, Title: "x", Kind: curriculum.Practice, Keys: []rune("fj")},
		"no title":              {Number: 6, Kind: curriculum.Intro, Pages: []string{"hi"}},
		"travel without a home": {Number: 7, Title: "x", Kind: curriculum.Travel, Keys: []rune("jum"), GoalWPM: 8},
	}
	for name, l := range cases {
		if err := l.Validate(); !errors.Is(err, curriculum.ErrInvalidLesson) {
			t.Errorf("%s: err = %v, want ErrInvalidLesson", name, err)
		}
	}
}

func TestModuleFindsLessonsByNumber(t *testing.T) {
	m := curriculum.Module{Title: "Home Row", Lessons: []curriculum.Lesson{
		{Number: 1, Title: "Welcome"},
		{Number: 2, Title: "Keys f and j"},
	}}

	l, ok := m.Lesson(2)
	if !ok || l.Title != "Keys f and j" {
		t.Fatalf("Lesson(2) = %v, %v", l.Title, ok)
	}
	if _, ok := m.Lesson(99); ok {
		t.Fatal("Lesson(99) should not exist")
	}
}
