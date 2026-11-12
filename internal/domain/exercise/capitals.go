package exercise

import (
	"math/rand/v2"
	"strings"
	"unicode"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func CapitalWords(rng *rand.Rand, keys, capitals []rune, count int) string {
	lower := unshifted(keys)
	words := typeable(lower)
	if len(words) == 0 {
		return CapitalGroups(rng, keys, capitals, count)
	}
	out := make([]string, count)
	for i := range out {
		out[i] = words[rng.IntN(len(words))]
		if rng.IntN(focusedShare) > 0 {
			out[i] = capitalWord(rng, words, lower, capitals[rng.IntN(len(capitals))])
		}
	}
	return strings.Join(out, " ")
}

func CapitalGroups(rng *rand.Rand, keys, capitals []rune, count int) string {
	lower := unshifted(keys)
	groups := make([]string, count)
	for i := range groups {
		groups[i] = capitalGroup(rng, lower, capitals[rng.IntN(len(capitals))])
	}
	return strings.Join(groups, " ")
}

func capitalWord(rng *rand.Rand, words []string, lower []rune, capital rune) string {
	if !unicode.IsLetter(capital) {
		return words[rng.IntN(len(words))] + string(capital)
	}
	var starting []string
	for _, w := range words {
		if strings.HasPrefix(w, string(unicode.ToLower(capital))) {
			starting = append(starting, w)
		}
	}
	if len(starting) == 0 {
		return capitalGroup(rng, lower, capital)
	}
	word := starting[rng.IntN(len(starting))]
	return string(capital) + word[1:]
}

func capitalGroup(rng *rand.Rand, lower []rune, capital rune) string {
	rest := group(rng, Mix{Keys: lower}, minGroup-1+rng.IntN(maxGroup-minGroup+1))
	if unicode.IsLetter(capital) {
		return string(capital) + rest
	}
	return rest + string(capital)
}

func unshifted(keys []rune) []rune {
	var lower []rune
	for _, r := range keys {
		if key, ok := keyboard.Lookup(r); !ok || !key.Shifted {
			lower = append(lower, r)
		}
	}
	return lower
}
