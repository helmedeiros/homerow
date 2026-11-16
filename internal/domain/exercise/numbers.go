package exercise

import (
	"math/rand/v2"
	"strings"
)

const longestNumber = 4

func NumberWords(rng *rand.Rand, keys, numbers []rune, count int) string {
	words := typeable(keys)
	items := make([]string, count)
	for i := range items {
		if len(words) == 0 || rng.IntN(focusedShare) > 0 {
			items[i] = group(rng, Mix{Keys: numbers}, 1+rng.IntN(longestNumber))
			continue
		}
		items[i] = words[rng.IntN(len(words))]
	}
	return strings.Join(items, " ")
}
