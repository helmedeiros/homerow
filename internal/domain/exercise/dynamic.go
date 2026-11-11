package exercise

import (
	"math/rand/v2"
	"strings"
)

func Dynamic(rng *rand.Rand, keys, weak []rune, count int) string {
	if len(weak) == 0 {
		return Groups(rng, Mix{Keys: keys}, count)
	}
	mix := Mix{Keys: keys, Focus: weak}
	groups := make([]string, count)
	for i := range groups {
		group := []rune(group(rng, mix, minGroup+rng.IntN(maxGroup-minGroup+1)))
		if !strings.ContainsAny(string(group), string(weak)) {
			group[rng.IntN(len(group))] = weak[rng.IntN(len(weak))]
		}
		groups[i] = string(group)
	}
	return strings.Join(groups, " ")
}
