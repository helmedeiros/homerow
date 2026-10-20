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
