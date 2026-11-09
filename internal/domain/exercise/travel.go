package exercise

import (
	"math/rand/v2"
	"strings"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func Travel(rng *rand.Rand, homes, keys []rune, count int) string {
	groups := make([]string, count)
	for i := range groups {
		home := homes[rng.IntN(len(homes))]
		reach := home
		if reaches := reachesOf(home, keys); len(reaches) > 0 {
			reach = reaches[rng.IntN(len(reaches))]
		}
		groups[i] = string([]rune{home, reach, home})
	}
	return strings.Join(groups, " ")
}

func reachesOf(home rune, keys []rune) []rune {
	from, ok := keyboard.Lookup(home)
	if !ok {
		return nil
	}
	var reaches []rune
	for _, k := range keys {
		if to, ok := keyboard.Lookup(k); ok && k != home && !to.Shifted && to.Finger == from.Finger {
			reaches = append(reaches, k)
		}
	}
	return reaches
}
