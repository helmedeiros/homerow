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

func Words(rng *rand.Rand, keys []rune, count int) (string, error) {
	candidates := typeable(keys)
	if len(candidates) == 0 {
		return "", ErrNoWords
	}
	words := make([]string, count)
	for i := range words {
		words[i] = candidates[rng.IntN(len(candidates))]
	}
	return strings.Join(words, " "), nil
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
