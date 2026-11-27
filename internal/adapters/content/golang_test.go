package content_test

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

const codeWidth = 48

func TestGoIsTheEighthModule(t *testing.T) {
	m := content.Course().Modules[7]

	if m.Number != 8 || m.Title != "Go" || m.Lessons[0].Number != 219 || m.Lessons[0].Kind != curriculum.Intro {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
}

func TestEveryCodeLessonExplainsItsPatternAndFitsTheScreen(t *testing.T) {
	for _, m := range content.Course().Modules {
		for _, l := range m.Lessons {
			if l.Kind != curriculum.Code {
				continue
			}
			if len(l.Pages) == 0 || len(l.Pages[0]) < 40 {
				t.Errorf("lesson %d has no note", l.Number)
			}
			for _, snippet := range l.Snippets {
				checkSnippet(t, l.Number, snippet)
			}
		}
	}
}

func checkSnippet(t *testing.T, number int, snippet string) {
	t.Helper()
	for _, r := range snippet {
		if _, ok := keyboard.Lookup(r); !ok {
			t.Errorf("lesson %d uses %q", number, r)
		}
	}
	for _, line := range strings.Split(snippet, "\n") {
		if len(strings.ReplaceAll(line, "\t", "    ")) > codeWidth || strings.HasSuffix(line, " ") {
			t.Errorf("lesson %d line %q", number, line)
		}
	}
}

func teaches(t *testing.T, want map[int]string) {
	t.Helper()
	for number, pattern := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Code || !strings.Contains(strings.Join(l.Snippets, "\n"), pattern) {
			t.Errorf("lesson %d = %v without %q", number, l.Kind, pattern)
		}
	}
}

func TestGoDeclarationsAndFunctions(t *testing.T) {
	teaches(t, map[int]string{
		220: ":=", 221: "const (", 222: "func add(a, b int) int {",
		223: "(float64, error)", 224: "return lo, hi",
	})
}

func TestTheFirstCodeLessonIntroducesEnterAndIndentationIsNotAKey(t *testing.T) {
	first, _, _ := content.Course().Lesson(220)
	if string(first.NewKeys) != "\n" {
		t.Fatalf("lesson 220 introduces %q", string(first.NewKeys))
	}
	for _, l := range content.Go().Lessons {
		if strings.ContainsRune(string(l.Keys), '\t') {
			t.Errorf("lesson %d asks for tab, but indentation is typed for you", l.Number)
		}
	}
}

func TestGoErrors(t *testing.T) {
	teaches(t, map[int]string{
		225: "if err != nil {", 226: "%w", 227: "errors.Is(err, ErrNotFound)", 228: "os.ReadFile",
	})
}
