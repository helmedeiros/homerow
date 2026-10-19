package exercise

import (
	"math/rand/v2"
	"strings"
)

const (
	minGroup = 2
	maxGroup = 5
)

type Mix struct {
	Keys  []rune
	Focus []rune
}

func Groups(rng *rand.Rand, mix Mix, count int) string {
	groups := make([]string, count)
	for i := range groups {
		groups[i] = group(rng, mix, minGroup+rng.IntN(maxGroup-minGroup+1))
	}
	return strings.Join(groups, " ")
}

func group(rng *rand.Rand, mix Mix, length int) string {
	var b strings.Builder
	for range length {
		b.WriteRune(pick(rng, mix))
	}
	return b.String()
}

func pick(rng *rand.Rand, mix Mix) rune {
	if len(mix.Focus) > 0 && rng.IntN(2) == 0 {
		return mix.Focus[rng.IntN(len(mix.Focus))]
	}
	return mix.Keys[rng.IntN(len(mix.Keys))]
}
