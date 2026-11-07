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

func TestPostureIntroThenAReviewOfEverythingSoFar(t *testing.T) {
	intro, _, _ := content.Course().Lesson(30)
	review, _, _ := content.Course().Lesson(31)

	if intro.Kind != curriculum.Intro || len(intro.Pages) < 3 {
		t.Fatalf("lesson 30 = %+v", intro)
	}
	if review.Kind != curriculum.Practice || !review.Words || string(review.Focus) != "ruei" {
		t.Fatalf("lesson 31 = %+v", review)
	}
}

func TestTopRowHandDrillsAddTheNewKeysOfEachHand(t *testing.T) {
	left, _, _ := content.Course().Lesson(32)
	right, _, _ := content.Course().Lesson(33)
	play, _, _ := content.Course().Lesson(34)

	if left.Kind != curriculum.HandDrill || string(left.Keys) != "asdfgre" {
		t.Errorf("lesson 32 = %+v", left)
	}
	if right.Kind != curriculum.HandDrill || string(right.Keys) != "hjkl;ui" {
		t.Errorf("lesson 33 = %+v", right)
	}
	if play.Kind != curriculum.Play || string(play.Keys) != "ruei" {
		t.Errorf("lesson 34 = %+v", play)
	}
}

func TestTopRowAddsWOThenQY(t *testing.T) {
	want := map[int]struct {
		kind curriculum.Kind
		keys string
	}{
		35: {curriculum.NewKeys, "wo"},
		36: {curriculum.Review, "asdfghjkl;rueiwo"},
		37: {curriculum.Practice, "asdfghjkl;rueiwo"},
		38: {curriculum.Play, "rueiwo"},
		39: {curriculum.NewKeys, "qy"},
		40: {curriculum.Review, "asdfghjkl;rueiwoqy"},
		41: {curriculum.Practice, "asdfghjkl;rueiwoqy"},
	}
	for number, w := range want {
		l, _, _ := content.Course().Lesson(number)
		keys := string(l.Keys)
		if l.Kind == curriculum.NewKeys {
			keys = string(l.NewKeys)
		}
		if l.Kind != w.kind || keys != w.keys {
			t.Errorf("lesson %d = %v %q, want %v %q", number, l.Kind, keys, w.kind, w.keys)
		}
	}
}
