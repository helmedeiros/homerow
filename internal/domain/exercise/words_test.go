package exercise_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func TestWordsUseOnlyLearnedKeys(t *testing.T) {
	keys := []rune("asdfjkl;")
	for seed := uint64(0); seed < 200; seed++ {
		text, err := exercise.Words(seeded(seed), keys, 12)
		if err != nil {
			t.Fatal(err)
		}
		words := strings.Split(text, " ")
		if len(words) != 12 {
			t.Fatalf("seed %d: %d words, want 12", seed, len(words))
		}
		for _, r := range strings.ReplaceAll(text, " ", "") {
			if !strings.ContainsRune(string(keys), r) {
				t.Fatalf("seed %d: %q uses %q", seed, text, r)
			}
		}
	}
}

func TestWordsIncludeTheWholeHomeRowOnceGAndHAreLearned(t *testing.T) {
	text, err := exercise.Words(seeded(3), []rune("asdfghjkl;"), 200)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.ContainsAny(text, "gh") {
		t.Fatalf("expected words with g or h in %q", text)
	}
}

func TestWordsFailWhenNoWordFits(t *testing.T) {
	_, err := exercise.Words(seeded(1), []rune("fj"), 5)
	if !errors.Is(err, exercise.ErrNoWords) {
		t.Fatalf("err = %v, want ErrNoWords", err)
	}
}

func TestWordsAreDeterministicForASeed(t *testing.T) {
	a, _ := exercise.Words(seeded(9), []rune("asdfjkl"), 10)
	b, _ := exercise.Words(seeded(9), []rune("asdfjkl"), 10)
	if a != b {
		t.Fatalf("same seed produced %q and %q", a, b)
	}
}

func TestEveryTopRowStageHasPlentyOfWordsWithItsNewKeys(t *testing.T) {
	home := "asdfghjkl;"
	stages := []struct{ learned, fresh string }{
		{home + "ru", "ru"},
		{home + "ruei", "ei"},
		{home + "rueiwo", "wo"},
		{home + "rueiwoqy", "qy"},
		{home + "rueiwoqypt", "pt"},
	}
	for _, s := range stages {
		text, err := exercise.Words(seeded(1), []rune(s.learned), 5000)
		if err != nil {
			t.Fatal(err)
		}
		distinct, fresh := map[string]bool{}, map[string]bool{}
		for _, w := range strings.Split(text, " ") {
			distinct[w] = true
			if strings.ContainsAny(w, s.fresh) {
				fresh[w] = true
			}
		}
		if len(distinct) < 100 || len(fresh) < 40 {
			t.Errorf("%s: %d words, %d with %s", s.learned, len(distinct), len(fresh), s.fresh)
		}
	}
}

func TestFocusedWordsMostlyUseTheNewKeys(t *testing.T) {
	keys := []rune("asdfghjkl;qwertyuiopvm")
	text, err := exercise.FocusedWords(seeded(3), keys, []rune("vm"), 300)
	if err != nil {
		t.Fatal(err)
	}

	words := strings.Split(text, " ")
	focused := 0
	for _, w := range words {
		if strings.Trim(w, string(keys)) != "" {
			t.Fatalf("%q uses keys outside the lesson", w)
		}
		if strings.ContainsAny(w, "vm") {
			focused++
		}
	}
	if focused < 180 || focused == len(words) {
		t.Fatalf("%d of %d words use v or m, want most but not all", focused, len(words))
	}
}

func TestFocusedWordsFallBackToAnyWordWhenNoneUseTheFocus(t *testing.T) {
	text, err := exercise.FocusedWords(seeded(3), []rune("asdfjkl;"), []rune("z"), 20)

	if err != nil || len(strings.Split(text, " ")) != 20 {
		t.Fatalf("text = %q, err = %v", text, err)
	}
	if _, err := exercise.FocusedWords(seeded(3), []rune("fj"), []rune("f"), 5); !errors.Is(err, exercise.ErrNoWords) {
		t.Fatalf("err = %v, want ErrNoWords", err)
	}
}
