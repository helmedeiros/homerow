package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestBottomRowIsTheThirdModule(t *testing.T) {
	m := content.Course().Modules[2]

	if m.Number != 3 || m.Title != "Bottom Row" || m.Lessons[0].Number != 52 {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
}

func TestBottomRowStartsWithVAndM(t *testing.T) {
	want := map[int]curriculum.Kind{52: curriculum.NewKeys, 53: curriculum.Review, 54: curriculum.Practice}
	for number, kind := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != kind || !strings.Contains(string(l.Keys), "vm") {
			t.Errorf("lesson %d = %v %q", number, l.Kind, string(l.Keys))
		}
	}
	first, _, _ := content.Course().Lesson(52)
	if string(first.NewKeys) != "vm" {
		t.Fatalf("lesson 52 introduces %q", string(first.NewKeys))
	}
}

func TestIndexFingersTravelToTheirReaches(t *testing.T) {
	want := map[int]struct{ home, reaches string }{
		55: {"j", "uyhm"},
		56: {"f", "rtgv"},
	}
	for number, w := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Travel || string(l.Focus) != w.home {
			t.Fatalf("lesson %d = %v from %q", number, l.Kind, string(l.Focus))
		}
		for _, g := range strings.Fields(l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text) {
			if g[:1] != w.home || g[2:] != w.home || !strings.Contains(w.reaches, g[1:2]) {
				t.Fatalf("lesson %d group %q", number, g)
			}
		}
	}
}
