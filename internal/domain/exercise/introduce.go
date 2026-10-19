package exercise

import "strings"

func Introduce(keys []rune) string {
	var groups []string
	for _, k := range keys {
		groups = append(groups, repeatPattern(k)...)
	}
	if len(keys) == 2 {
		groups = append(groups, pairPattern(keys[0], keys[1])...)
	}
	return strings.Join(groups, " ")
}

func repeatPattern(k rune) []string {
	one, two, three := string(k), strings.Repeat(string(k), 2), strings.Repeat(string(k), 3)
	return []string{one, one, one, two, two, three}
}

func pairPattern(a, b rune) []string {
	x, y := string(a), string(b)
	return []string{
		x + y,
		y + x,
		x + x + y + y,
		y + y + x + x,
		x + y + x + y,
		y + x + y + x,
	}
}
