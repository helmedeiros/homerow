package content_test

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestTopRowIsTheSecondModule(t *testing.T) {
	m := content.Course().Modules[1]

	if m.Number != 2 || m.Title != "Top Row" || m.Lessons[0].Number != 24 {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
}

func TestTopRowStartsWithRAndU(t *testing.T) {
	first, _, _ := content.Course().Lesson(24)

	if first.Kind != curriculum.NewKeys || string(first.NewKeys) != "ru" || !strings.Contains(string(first.Keys), "asdfghjkl;") {
		t.Fatalf("lesson 24 = %+v", first)
	}
}

func TestTopRowPractisesWithRealWords(t *testing.T) {
	for _, number := range []int{26, 29} {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Practice || !l.Words || l.TimeLimit == 0 {
			t.Errorf("lesson %d = %+v", number, l)
		}
	}
}

func TestTopRowKeysArriveInPairs(t *testing.T) {
	want := map[int]string{24: "ru", 27: "ei"}
	for number, keys := range want {
		l, _, _ := content.Course().Lesson(number)
		if string(l.NewKeys) != keys {
			t.Errorf("lesson %d introduces %q, want %q", number, string(l.NewKeys), keys)
		}
	}
}
