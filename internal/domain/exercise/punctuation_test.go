package exercise_test

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

const (
	punctuationLetters = "asdfghjkl;qwertyuiopzxcvbnm,./"
	punctuationKeys    = punctuationLetters + "-'[]=\\_+{}\"|<>`~"
)

func punctuationItems(t *testing.T, marks string) []string {
	t.Helper()
	text := exercise.PunctuationWords(seeded(2), []rune(punctuationKeys), []rune(marks), 60)
	items := strings.Fields(text)
	if len(items) != 60 {
		t.Fatalf("%d items: %q", len(items), text)
	}
	for _, r := range strings.ReplaceAll(text, " ", "") {
		if !strings.ContainsRune(punctuationKeys, r) {
			t.Fatalf("%q uses %q", text, r)
		}
	}
	return items
}

func isWord(s string) bool {
	return s != "" && strings.Trim(s, punctuationLetters) == ""
}

func TestBracketsAndQuotesWrapAWord(t *testing.T) {
	pairs := map[byte]byte{'[': ']', '{': '}', '<': '>', '"': '"', '\'': '\'', '`': '`'}
	seen := 0
	for _, item := range punctuationItems(t, "[]{}<>\"'`") {
		closing, ok := pairs[item[0]]
		if !ok {
			continue
		}
		if item[len(item)-1] != closing || !isWord(item[1:len(item)-1]) {
			t.Fatalf("item %q", item)
		}
		seen++
	}
	if seen == 0 {
		t.Fatal("no wrapped words")
	}
}

func TestJoinersSitBetweenTwoWords(t *testing.T) {
	seen := map[rune]bool{}
	for _, item := range punctuationItems(t, "-_=+\\|") {
		i := strings.IndexAny(item, "-_=+\\|")
		if i < 0 {
			continue
		}
		if !isWord(item[:i]) || !isWord(item[i+1:]) {
			t.Fatalf("item %q", item)
		}
		seen[rune(item[i])] = true
	}
	if len(seen) != 6 {
		t.Fatalf("seen %v", seen)
	}
}

func TestTildeLeadsAWord(t *testing.T) {
	for _, item := range punctuationItems(t, "~") {
		if strings.Contains(item, "~") && (item[0] != '~' || !isWord(item[1:])) {
			t.Fatalf("item %q", item)
		}
	}
}

func TestABracketWithoutItsPartnerLeadsOrFollowsAWord(t *testing.T) {
	text := exercise.PunctuationWords(seeded(4), []rune("asdfjkl;["), []rune("["), 30)

	for _, item := range strings.Fields(text) {
		if strings.Contains(item, "[") && (item[0] != '[' || !isWord(item[1:])) {
			t.Fatalf("item %q in %q", item, text)
		}
	}
}

func TestAClosingBracketWithoutItsPartnerFollowsAWord(t *testing.T) {
	text := exercise.PunctuationWords(seeded(4), []rune("asdfjkl;]"), []rune("]"), 30)

	for _, item := range strings.Fields(text) {
		if strings.Contains(item, "]") && (item[len(item)-1] != ']' || !isWord(item[:len(item)-1])) {
			t.Fatalf("item %q in %q", item, text)
		}
	}
}

func TestPunctuationWithoutWordsStillAppears(t *testing.T) {
	text := exercise.PunctuationWords(seeded(3), []rune("fj-="), []rune("-="), 20)

	if !strings.ContainsAny(text, "-=") || len(strings.Fields(text)) != 20 {
		t.Fatalf("text = %q", text)
	}
}
