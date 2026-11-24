package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestPunctuationIsTheSeventhModule(t *testing.T) {
	m := content.Course().Modules[6]

	if m.Number != 7 || m.Title != "Punctuation" || m.Lessons[0].Number != 188 {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
}

func TestPunctuationOpensWithThePinkyCorner(t *testing.T) {
	l, _, _ := content.Course().Lesson(188)

	if l.Kind != curriculum.Intro || len(l.Pages) == 0 || !strings.Contains(strings.Join(l.Pages, " "), "pinky") {
		t.Fatalf("lesson 188 = %v %q", l.Kind, l.Pages)
	}
}

func TestDashAndApostrophe(t *testing.T) {
	introduces(t, map[int]string{189: "-'"})

	l, _, _ := content.Course().Lesson(191)
	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Practice || !strings.ContainsAny(text, "-'") {
		t.Fatalf("lesson 191 = %v: %q", l.Kind, text)
	}
}

func TestBracketsThenEqualsAndBackslash(t *testing.T) {
	introduces(t, map[int]string{192: "[]", 195: "=\\"})

	l, _, _ := content.Course().Lesson(194)
	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Practice || !strings.Contains(text, "[") || !strings.Contains(text, "]") {
		t.Fatalf("lesson 194 = %v: %q", l.Kind, text)
	}
	l, _, _ = content.Course().Lesson(196)
	if l.Kind != curriculum.Review || string(l.Focus) != "=\\" {
		t.Fatalf("lesson 196 = %v focusing on %q", l.Kind, string(l.Focus))
	}
}

func TestUnderscoreAndPlusThenBraces(t *testing.T) {
	introduces(t, map[int]string{198: "_+", 201: "{}"})

	l, _, _ := content.Course().Lesson(200)
	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Practice || !strings.ContainsAny(text, "_+") {
		t.Fatalf("lesson 200 = %v: %q", l.Kind, text)
	}
}
