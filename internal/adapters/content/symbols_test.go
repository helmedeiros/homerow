package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestSymbolsIsTheSixthModule(t *testing.T) {
	m := content.Course().Modules[5]

	if m.Number != 6 || m.Title != "Symbols" || m.Lessons[0].Number != 158 {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
}

func TestDollarAndAmpersandThenHashAndStar(t *testing.T) {
	introduces(t, map[int]string{158: "$&", 161: "#*"})

	l, _, _ := content.Course().Lesson(160)
	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Practice || !strings.Contains(text, "$") {
		t.Fatalf("lesson 160 = %v: %q", l.Kind, text)
	}
}

func TestAtAndOpenBracketThenBangAndCloseBracket(t *testing.T) {
	introduces(t, map[int]string{164: "@(", 167: "!)"})
}
