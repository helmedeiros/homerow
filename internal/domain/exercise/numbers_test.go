package exercise_test

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func TestNumberWordsMixNumbersWithWords(t *testing.T) {
	keys := "asdfghjkl;47"
	text := exercise.NumberWords(seeded(1), []rune(keys), []rune("47"), 40)

	items := strings.Fields(text)
	if len(items) != 40 {
		t.Fatalf("%d items: %q", len(items), text)
	}
	numbers, words := 0, 0
	for _, item := range items {
		switch {
		case strings.Trim(item, "47") == "":
			numbers++
		case strings.Trim(item, "asdfghjkl;") == "":
			words++
		default:
			t.Fatalf("item %q mixes digits and letters or uses other keys", item)
		}
	}
	if numbers == 0 || words == 0 {
		t.Fatalf("numbers = %d, words = %d in %q", numbers, words, text)
	}
}

func TestNumberWordsWithoutWordsUseOnlyNumbers(t *testing.T) {
	text := exercise.NumberWords(seeded(2), []rune("fj47"), []rune("47"), 10)

	if strings.Trim(strings.ReplaceAll(text, " ", ""), "47") != "" || len(strings.Fields(text)) != 10 {
		t.Fatalf("text = %q", text)
	}
}
