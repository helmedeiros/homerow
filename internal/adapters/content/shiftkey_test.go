package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestShiftKeyIsTheFourthModule(t *testing.T) {
	m := content.Course().Modules[3]

	if m.Number != 4 || m.Title != "Shift Key" || m.Lessons[0].Number != 89 {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
	if m.Lessons[0].Kind != curriculum.Intro {
		t.Fatalf("lesson 89 = %v, want an intro", m.Lessons[0].Kind)
	}
}

func TestCapitalsArriveInPairs(t *testing.T) {
	want := map[int]string{90: "FJ", 93: "DK"}
	for number, keys := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.NewKeys || string(l.NewKeys) != keys {
			t.Errorf("lesson %d = %v introducing %q", number, l.Kind, string(l.NewKeys))
		}
	}
}

func TestCapitalPracticeUsesCapitalisedWords(t *testing.T) {
	l, _, _ := content.Course().Lesson(95)

	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Practice || !strings.ContainsAny(text, "DK") {
		t.Fatalf("lesson 95 = %v: %q", l.Kind, text)
	}
}

func TestFirstCapitalsArePlayedAsFallingLetters(t *testing.T) {
	l, _, _ := content.Course().Lesson(96)

	if l.Kind != curriculum.Play || l.Game != curriculum.FallingLetters || string(l.Keys) != "fjdkFJDK" {
		t.Fatalf("lesson 96 = %v %v %q", l.Kind, l.Game, string(l.Keys))
	}
}
