package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestNumbersIsTheFifthModule(t *testing.T) {
	m := content.Course().Modules[4]

	if m.Number != 5 || m.Title != "Numbers" || m.Lessons[0].Number != 143 {
		t.Fatalf("module = %d %q starting at %d", m.Number, m.Title, m.Lessons[0].Number)
	}
}

func TestFourAndSevenThenThreeAndEight(t *testing.T) {
	introduces(t, map[int]string{143: "47", 146: "38"})
}

func TestNumberPracticeMixesDigitsAndWords(t *testing.T) {
	l, _, _ := content.Course().Lesson(148)

	text := l.Steps(rand.New(rand.NewPCG(1, 1)))[0].Text
	if l.Kind != curriculum.Practice || !strings.ContainsAny(text, "38") {
		t.Fatalf("lesson 148 = %v: %q", l.Kind, text)
	}
}

func TestTwoAndNineThenOneAndZeroThenFiveAndSix(t *testing.T) {
	introduces(t, map[int]string{149: "29", 152: "10", 155: "56"})

	last := content.Numbers().Lessons[len(content.Numbers().Lessons)-1]
	if last.Number != 157 || last.Kind != curriculum.Practice || !strings.Contains(string(last.Keys), "1234567890") {
		t.Fatalf("last lesson = %d %v %q", last.Number, last.Kind, string(last.Keys))
	}
}
