package exercise_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func checkCapitalised(t *testing.T, text, keys string, count int) {
	t.Helper()
	words := strings.Fields(text)
	if len(words) != count {
		t.Fatalf("%d words, want %d: %q", len(words), count, text)
	}
	for _, w := range words {
		for i, r := range w {
			if !strings.ContainsRune(keys, r) {
				t.Fatalf("%q uses %q", w, r)
			}
			if unicode.IsUpper(r) && i > 0 {
				t.Fatalf("%q has a capital inside", w)
			}
			if r == ':' && i != len(w)-1 {
				t.Fatalf("%q has a colon inside", w)
			}
		}
	}
}

func TestCapitalWordsStartWithTheNewCapitals(t *testing.T) {
	keys := "asdfghjkl;FJ"
	for seed := uint64(0); seed < 50; seed++ {
		text := exercise.CapitalWords(seeded(seed), []rune(keys), []rune("FJ"), 30)

		checkCapitalised(t, text, keys, 30)
		if !strings.ContainsAny(text, "FJ") {
			t.Fatalf("seed %d: no capitals in %q", seed, text)
		}
	}
}

func TestCapitalWordsEndWithAColon(t *testing.T) {
	keys := "asdfghjkl;:"
	text := exercise.CapitalWords(seeded(1), []rune(keys), []rune(":"), 30)

	checkCapitalised(t, text, keys, 30)
	if !strings.Contains(text, ":") {
		t.Fatalf("no colon in %q", text)
	}
}

func TestCapitalWordsFallBackToGroupsWhenNoWordFits(t *testing.T) {
	keys := "fjF"
	text := exercise.CapitalWords(seeded(2), []rune(keys), []rune("F"), 20)

	checkCapitalised(t, text, keys, 20)
	if !strings.Contains(text, "F") {
		t.Fatalf("no capital in %q", text)
	}
}

func TestCapitalGroupsStartWithACapital(t *testing.T) {
	keys := "asdfjkl;DK"
	text := exercise.CapitalGroups(seeded(4), []rune(keys), []rune("DK"), 16)

	checkCapitalised(t, text, keys, 16)
	for _, g := range strings.Fields(text) {
		if !strings.ContainsRune("DK", rune(g[0])) {
			t.Fatalf("group %q does not start with a capital", g)
		}
	}
}
