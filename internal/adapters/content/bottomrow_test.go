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
		60: {"k", "i,"},
		61: {"d", "ec"},
		67: {"l", "o."},
		68: {"s", "wx"},
		72: {";", "p/"},
		73: {"a", "qz"},
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

func TestCAndCommaArriveBeforeTheBalloons(t *testing.T) {
	keys, _, _ := content.Course().Lesson(57)
	if keys.Kind != curriculum.NewKeys || string(keys.NewKeys) != "c," {
		t.Fatalf("lesson 57 = %v %q", keys.Kind, string(keys.NewKeys))
	}
	for _, number := range []int{58, 59} {
		l, _, _ := content.Course().Lesson(number)
		if string(l.Focus) != "c," {
			t.Errorf("lesson %d focus = %q", number, string(l.Focus))
		}
	}
	play, _, _ := content.Course().Lesson(62)
	if play.Kind != curriculum.Play || play.Game != curriculum.BalloonWords || !strings.ContainsAny(string(play.Keys), "vmc") {
		t.Fatalf("lesson 62 = %+v", play)
	}
	if step := play.Steps(rand.New(rand.NewPCG(1, 1)))[0]; step.Game != curriculum.BalloonWords {
		t.Fatalf("lesson 62 step = %+v", step)
	}
}

func TestQwertyHistoryComesBeforeXAndPeriod(t *testing.T) {
	intro, _, _ := content.Course().Lesson(63)
	if intro.Kind != curriculum.Intro || intro.Title != "History of QWERTY" || len(intro.Pages) == 0 {
		t.Fatalf("lesson 63 = %+v", intro)
	}
	keys, _, _ := content.Course().Lesson(64)
	if keys.Kind != curriculum.NewKeys || string(keys.NewKeys) != "x." {
		t.Fatalf("lesson 64 = %v %q", keys.Kind, string(keys.NewKeys))
	}
	for _, number := range []int{65, 66} {
		l, _, _ := content.Course().Lesson(number)
		if string(l.Focus) != "x." {
			t.Errorf("lesson %d focus = %q", number, string(l.Focus))
		}
	}
}

func TestZAndSlashLeadToTheBottomRowBalloons(t *testing.T) {
	keys, _, _ := content.Course().Lesson(69)
	if keys.Kind != curriculum.NewKeys || string(keys.NewKeys) != "z/" {
		t.Fatalf("lesson 69 = %v %q", keys.Kind, string(keys.NewKeys))
	}
	play, _, _ := content.Course().Lesson(74)
	if play.Game != curriculum.BalloonWords || !strings.Contains(string(play.Keys), "vmc,x.z/") {
		t.Fatalf("lesson 74 = %+v", play)
	}
}

func TestBAndNCompleteTheLettersBeforeABreak(t *testing.T) {
	keys, _, _ := content.Course().Lesson(75)
	if keys.Kind != curriculum.NewKeys || string(keys.NewKeys) != "bn" {
		t.Fatalf("lesson 75 = %v %q", keys.Kind, string(keys.NewKeys))
	}
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.ContainsRune(string(keys.Keys), r) {
			t.Fatalf("lesson 75 is missing %q", r)
		}
	}
	rest, _, _ := content.Course().Lesson(78)
	if rest.Kind != curriculum.Intro || rest.Title != "Take a break" {
		t.Fatalf("lesson 78 = %+v", rest)
	}
}

func TestBottomRowHandDrillsAndPractice(t *testing.T) {
	want := map[int]struct {
		kind curriculum.Kind
		keys string
	}{
		79: {curriculum.HandDrill, "asdfgzxcvb"},
		80: {curriculum.HandDrill, "hjkl;nm,./"},
		81: {curriculum.Practice, "asdfghjkl;rueiwoqyptvmc,x.z/bn"},
	}
	for number, w := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != w.kind || string(l.Keys) != w.keys {
			t.Errorf("lesson %d = %v %q", number, l.Kind, string(l.Keys))
		}
	}
}

func TestBottomRowPlaysThenTravelsWithEveryFinger(t *testing.T) {
	popUp, _, _ := content.Course().Lesson(82)
	if popUp.Game != curriculum.PopUpWarmUp || string(popUp.Keys) != "zxcvbnm,./" {
		t.Fatalf("lesson 82 = %+v", popUp)
	}
	falling, _, _ := content.Course().Lesson(83)
	if falling.Kind != curriculum.Play || falling.Game != curriculum.FallingLetters {
		t.Fatalf("lesson 83 = %+v", falling)
	}
	for number, homes := range map[int]string{84: "jkl;", 85: "asdf"} {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Travel || string(l.Focus) != homes {
			t.Fatalf("lesson %d = %v from %q", number, l.Kind, string(l.Focus))
		}
		starts := map[byte]bool{}
		for _, g := range strings.Fields(l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text) {
			if g[0] != g[2] || !strings.Contains(homes, g[:1]) {
				t.Fatalf("lesson %d group %q", number, g)
			}
			starts[g[0]] = true
		}
		if len(starts) != 4 {
			t.Fatalf("lesson %d travels from %d home keys", number, len(starts))
		}
	}
}

func TestTheBottomRowEndsWithTheTabKey(t *testing.T) {
	intro, _, _ := content.Course().Lesson(86)
	if intro.Kind != curriculum.Intro || len(intro.Pages) == 0 {
		t.Fatalf("lesson 86 = %+v", intro)
	}
	tab, _, _ := content.Course().Lesson(87)
	text := tab.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if tab.Kind != curriculum.Tab || !strings.Contains(text, "\t") {
		t.Fatalf("lesson 87 = %v %q", tab.Kind, text)
	}
}
