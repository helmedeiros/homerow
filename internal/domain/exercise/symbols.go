package exercise

import (
	"math/rand/v2"
	"strings"
	"unicode"
)

const (
	beforeNumbers = "$#"
	afterNumbers  = "%"
	beforeWords   = "@("
	afterWords    = "!)"
)

type fillers struct {
	words   []string
	numbers []rune
}

func SymbolWords(rng *rand.Rand, keys, symbols []rune, count int) string {
	f := fillers{words: typeable(unshifted(keys)), numbers: digitsOf(keys)}
	items := make([]string, count)
	for i := range items {
		if len(f.words) == 0 || rng.IntN(focusedShare) > 0 {
			items[i] = f.around(rng, symbols[rng.IntN(len(symbols))])
			continue
		}
		items[i] = f.word(rng)
	}
	return strings.Join(items, " ")
}

func (f fillers) around(rng *rand.Rand, symbol rune) string {
	s := string(symbol)
	switch {
	case strings.ContainsRune(beforeNumbers, symbol):
		return s + f.number(rng)
	case strings.ContainsRune(afterNumbers, symbol):
		return f.number(rng) + s
	case strings.ContainsRune(beforeWords, symbol):
		return s + f.word(rng)
	case strings.ContainsRune(afterWords, symbol):
		return f.word(rng) + s
	}
	return strings.Repeat(s, 1+rng.IntN(2))
}

func (f fillers) number(rng *rand.Rand) string {
	if len(f.numbers) == 0 {
		return ""
	}
	return group(rng, Mix{Keys: f.numbers}, 1+rng.IntN(longestNumber))
}

func (f fillers) word(rng *rand.Rand) string {
	if len(f.words) == 0 {
		return ""
	}
	return f.words[rng.IntN(len(f.words))]
}

func digitsOf(keys []rune) []rune {
	var digits []rune
	for _, r := range keys {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}
	return digits
}
