package exercise_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func TestTabbedPairsWordsAcrossATab(t *testing.T) {
	keys := "asdfghjkl;"
	text, err := exercise.Tabbed(seeded(4), []rune(keys), 6)
	if err != nil {
		t.Fatal(err)
	}

	pairs := strings.Split(text, " ")
	if len(pairs) != 6 {
		t.Fatalf("text = %q", text)
	}
	for _, p := range pairs {
		words := strings.Split(p, "\t")
		if len(words) != 2 || strings.Trim(words[0]+words[1], keys) != "" {
			t.Fatalf("pair %q", p)
		}
	}
}

func TestTabbedNeedsWords(t *testing.T) {
	if _, err := exercise.Tabbed(seeded(4), []rune("fj"), 3); !errors.Is(err, exercise.ErrNoWords) {
		t.Fatalf("err = %v, want ErrNoWords", err)
	}
}
