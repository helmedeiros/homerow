package exercise

import (
	"math/rand/v2"
	"strings"
)

func Tabbed(rng *rand.Rand, keys []rune, pairs int) (string, error) {
	text, err := Words(rng, keys, 2*pairs)
	if err != nil {
		return "", err
	}
	words := strings.Fields(text)
	rows := make([]string, pairs)
	for i := range rows {
		rows[i] = words[2*i] + "\t" + words[2*i+1]
	}
	return strings.Join(rows, " "), nil
}
