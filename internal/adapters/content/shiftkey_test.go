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

func TestSAndLThenAAndColon(t *testing.T) {
	want := map[int]string{97: "SL", 100: "A:"}
	for number, keys := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.NewKeys || string(l.NewKeys) != keys {
			t.Errorf("lesson %d = %v introducing %q", number, l.Kind, string(l.NewKeys))
		}
	}
}

func TestColonPracticeEndsWordsWithAColon(t *testing.T) {
	l, _, _ := content.Course().Lesson(102)

	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if !strings.Contains(text, ":") || !strings.Contains(text, "A") {
		t.Fatalf("lesson 102: %q", text)
	}
}

func TestHomeRowCapitalsWarmUpWithPopUps(t *testing.T) {
	l, _, _ := content.Course().Lesson(103)

	if l.Kind != curriculum.Play || l.Game != curriculum.PopUpWarmUp || !strings.Contains(string(l.Keys), "FJDKSLA:") {
		t.Fatalf("lesson 103 = %v %v %q", l.Kind, l.Game, string(l.Keys))
	}
}

func introduces(t *testing.T, want map[int]string) {
	t.Helper()
	for number, keys := range want {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.NewKeys || string(l.NewKeys) != keys {
			t.Errorf("lesson %d = %v introducing %q", number, l.Kind, string(l.NewKeys))
		}
	}
}

func plays(t *testing.T, number int, game curriculum.Game, keys string) {
	t.Helper()
	l, _, _ := content.Course().Lesson(number)
	if l.Kind != curriculum.Play || l.Game != game || !strings.Contains(string(l.Keys), keys) {
		t.Fatalf("lesson %d = %v %v %q", number, l.Kind, l.Game, string(l.Keys))
	}
}

func TestGAndHThenTAndY(t *testing.T) {
	introduces(t, map[int]string{104: "GH", 107: "TY"})
	plays(t, 110, curriculum.BalloonWords, "GHTY")
}

func TestRAndUThenEAndI(t *testing.T) {
	introduces(t, map[int]string{111: "RU", 114: "EI"})
	plays(t, 117, curriculum.PopUpWarmUp, "RUEI")
}

func TestWAndOThenQAndP(t *testing.T) {
	introduces(t, map[int]string{118: "WO", 121: "QP"})
	plays(t, 124, curriculum.BalloonWords, "WOQP")
}

func TestVAndMThenIndexTravelThenCAndQuestionMark(t *testing.T) {
	introduces(t, map[int]string{125: "VM", 129: "C?"})
	plays(t, 132, curriculum.BalloonWords, "VMC?")

	l, _, _ := content.Course().Lesson(128)
	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Travel || string(l.Focus) != "fj" || !strings.ContainsAny(text, "RTGVUYHM") {
		t.Fatalf("lesson 128 = %v from %q: %q", l.Kind, string(l.Focus), text)
	}
}

func TestZAndNThenXAndB(t *testing.T) {
	introduces(t, map[int]string{133: "ZN", 136: "XB"})
}

func TestShiftKeyEndsWithTravelAndDynamicPractice(t *testing.T) {
	for number, homes := range map[int]string{139: "dk", 140: "sl", 141: "a;"} {
		l, _, _ := content.Course().Lesson(number)
		if l.Kind != curriculum.Travel || string(l.Focus) != homes {
			t.Errorf("lesson %d = %v from %q", number, l.Kind, string(l.Focus))
		}
	}
	last := content.ShiftKey().Lessons[len(content.ShiftKey().Lessons)-1]
	if last.Number != 142 || last.Kind != curriculum.Dynamic || !strings.Contains(string(last.Keys), "XB") {
		t.Fatalf("last lesson = %d %v %q", last.Number, last.Kind, string(last.Keys))
	}
}
