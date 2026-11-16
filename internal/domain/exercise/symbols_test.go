package exercise_test

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

const symbolKeys = "asdfghjkl;qwertyuiop1234567890$#@(%!)&*^"

func symbolItems(t *testing.T, symbols string) []string {
	t.Helper()
	text := exercise.SymbolWords(seeded(1), []rune(symbolKeys), []rune(symbols), 60)
	items := strings.Fields(text)
	if len(items) != 60 {
		t.Fatalf("%d items: %q", len(items), text)
	}
	for _, r := range strings.ReplaceAll(text, " ", "") {
		if !strings.ContainsRune(symbolKeys, r) {
			t.Fatalf("%q uses %q", text, r)
		}
	}
	return items
}

func TestDollarAndHashLeadNumbers(t *testing.T) {
	for _, item := range symbolItems(t, "$#") {
		if strings.ContainsAny(item, "$#") && (!strings.ContainsAny(item[:1], "$#") || strings.Trim(item[1:], "1234567890") != "") {
			t.Fatalf("item %q", item)
		}
	}
}

func TestPercentFollowsANumberAndBangFollowsAWord(t *testing.T) {
	seen := map[byte]bool{}
	for _, item := range symbolItems(t, "%!") {
		last := item[len(item)-1]
		switch last {
		case '%':
			if strings.Trim(item[:len(item)-1], "1234567890") != "" {
				t.Fatalf("item %q", item)
			}
		case '!':
			if strings.Trim(item[:len(item)-1], "asdfghjkl;qwertyuiop") != "" {
				t.Fatalf("item %q", item)
			}
		}
		seen[last] = true
	}
	if !seen['%'] || !seen['!'] {
		t.Fatalf("seen %v", seen)
	}
}

func TestOtherSymbolsStandAlone(t *testing.T) {
	for _, item := range symbolItems(t, "&*^") {
		if strings.ContainsAny(item, "&*^") && strings.Trim(item, "&*^") != "" {
			t.Fatalf("item %q", item)
		}
	}
}

func TestSymbolsWithoutWordsOrNumbersStillAppear(t *testing.T) {
	text := exercise.SymbolWords(seeded(3), []rune("fj@!"), []rune("@!"), 20)

	if !strings.ContainsAny(text, "@!") || len(strings.Fields(text)) != 20 {
		t.Fatalf("text = %q", text)
	}
}
