package exercise

import (
	_ "embed"
	"errors"
	"math/rand/v2"
	"strings"
)

var ErrNoWords = errors.New("no words can be typed with these keys")

//go:embed words.txt
var wordList string

var dictionary = uniqueWords(wordList)

const focusedShare = 3

func Words(rng *rand.Rand, keys []rune, count int) (string, error) {
	return FocusedWords(rng, keys, nil, count)
}

func FocusedWords(rng *rand.Rand, keys, focus []rune, count int) (string, error) {
	candidates := typeable(keys)
	if len(candidates) == 0 {
		return "", ErrNoWords
	}
	focused := using(candidates, focus)
	words := make([]string, count)
	for i := range words {
		pool := candidates
		if len(focused) > 0 && rng.IntN(focusedShare) > 0 {
			pool = focused
		}
		words[i] = pool[rng.IntN(len(pool))]
	}
	return strings.Join(words, " "), nil
}

func using(words []string, focus []rune) []string {
	var fits []string
	for _, w := range words {
		if strings.ContainsAny(w, string(focus)) {
			fits = append(fits, w)
		}
	}
	return fits
}

func typeable(keys []rune) []string {
	allowed := string(keys)
	var fits []string
	for _, w := range dictionary {
		if len(w) > 1 && onlyUses(w, allowed) {
			fits = append(fits, w)
		}
	}
	return fits
}

func onlyUses(word, allowed string) bool {
	for _, r := range word {
		if !strings.ContainsRune(allowed, r) {
			return false
		}
	}
	return true
}

func uniqueWords(list string) []string {
	seen := map[string]bool{}
	var words []string
	for _, w := range strings.Fields(list) {
		if !seen[w] {
			seen[w] = true
			words = append(words, w)
		}
	}
	return words
}
