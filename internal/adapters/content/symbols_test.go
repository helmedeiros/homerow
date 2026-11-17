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

func TestPercentAndCaretThenPlayEverySymbol(t *testing.T) {
	introduces(t, map[int]string{170: "%^"})
	plays(t, 173, curriculum.BalloonWords, "$&#*@(!)%^")
}

func TestEachSymbolHasATravelLessonFromItsHomeKey(t *testing.T) {
	want := []struct{ home, reaches string }{
		{"f", "4$"}, {"d", "3#"}, {"s", "2@"}, {"a", "1!"}, {"f", "5%"},
		{"j", "7&"}, {"k", "8*"}, {"l", "9("}, {";", "0)"}, {"j", "6^"},
	}
	for i, w := range want {
		number := 174 + i
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Travel || string(l.Focus) != w.home {
			t.Fatalf("lesson %d = %v from %q", number, l.Kind, string(l.Focus))
		}
		text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
		if !strings.Contains(text, w.home+w.reaches[1:]+w.home) {
			t.Fatalf("lesson %d never reaches %q: %q", number, w.reaches[1:], text)
		}
		for _, g := range strings.Fields(text) {
			if g[:1] != w.home || !strings.Contains(w.reaches, g[1:2]) {
				t.Fatalf("lesson %d group %q", number, g)
			}
		}
	}
}

func TestSymbolsEndWithBothHandsThenReviewAndPractice(t *testing.T) {
	for number, homes := range map[int]string{184: "asdf", 185: "jkl;"} {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Travel || string(l.Focus) != homes {
			t.Errorf("lesson %d = %v from %q", number, l.Kind, string(l.Focus))
		}
	}
	for number, kind := range map[int]curriculum.Kind{186: curriculum.Review, 187: curriculum.Practice} {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != kind || string(l.Focus) != "$&#*@(!)%^" {
			t.Errorf("lesson %d = %v focusing on %q", number, l.Kind, string(l.Focus))
		}
	}
	if n := len(content.Course().Lessons()); n != 187 {
		t.Fatalf("course has %d lessons, want 187", n)
	}
}
