package exercise

import (
	"math/rand/v2"
	"strings"
)

const (
	joiners = "-_=+\\|"
	leaders = "~"
)

var partners = map[rune]rune{
	'[': ']', ']': '[', '{': '}', '}': '{', '<': '>', '>': '<',
	'"': '"', '\'': '\'', '`': '`',
}

func PunctuationWords(rng *rand.Rand, keys, marks []rune, count int) string {
	f := fillers{words: typeable(unshifted(keys))}
	allowed := string(keys)
	items := make([]string, count)
	for i := range items {
		if len(f.words) == 0 || rng.IntN(focusedShare) > 0 {
			items[i] = f.punctuate(rng, marks[rng.IntN(len(marks))], allowed)
			continue
		}
		items[i] = f.word(rng)
	}
	return strings.Join(items, " ")
}

func (f fillers) punctuate(rng *rand.Rand, mark rune, allowed string) string {
	s := string(mark)
	if len(f.words) == 0 {
		return strings.Repeat(s, 1+rng.IntN(2))
	}
	switch {
	case strings.ContainsRune(joiners, mark):
		return f.word(rng) + s + f.word(rng)
	case strings.ContainsRune(leaders, mark):
		return s + f.word(rng)
	}
	return wrap(mark, f.word(rng), allowed)
}

func wrap(mark rune, word, allowed string) string {
	partner := partners[mark]
	opens := strings.ContainsRune("[{<\"'`", mark)
	switch {
	case !strings.ContainsRune(allowed, partner) && opens:
		return string(mark) + word
	case !strings.ContainsRune(allowed, partner):
		return word + string(mark)
	case opens:
		return string(mark) + word + string(partner)
	}
	return string(partner) + word + string(mark)
}
