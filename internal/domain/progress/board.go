package progress

import (
	"maps"
	"slices"
)

type Entry struct {
	Completed bool
	Stars     int
	WPM       float64
	Accuracy  float64
}

type Board struct {
	entries  map[int]Entry
	mistakes map[rune]int
}

func NewBoard(entries map[int]Entry) *Board {
	b := &Board{entries: map[int]Entry{}, mistakes: map[rune]int{}}
	maps.Copy(b.entries, entries)
	return b
}

func (b *Board) Entry(lesson int) Entry {
	return b.entries[lesson]
}

func (b *Board) Record(lesson int, attempt Entry) bool {
	old := b.entries[lesson]
	best := Entry{
		Completed: old.Completed || attempt.Completed,
		Stars:     max(old.Stars, attempt.Stars),
		WPM:       max(old.WPM, attempt.WPM),
		Accuracy:  max(old.Accuracy, attempt.Accuracy),
	}
	b.entries[lesson] = best
	return best.Completed != old.Completed || best.Stars != old.Stars
}

func (b *Board) Entries() map[int]Entry {
	return maps.Clone(b.entries)
}

func (b *Board) AddMistakes(mistakes map[rune]int) {
	for key, n := range mistakes {
		if n > 0 {
			b.mistakes[key] += n
		}
	}
}

func (b *Board) Mistakes() map[rune]int {
	return maps.Clone(b.mistakes)
}

func (b *Board) Weakest(keys []rune, n int) []rune {
	var missed []rune
	for _, k := range keys {
		if b.mistakes[k] > 0 && !slices.Contains(missed, k) {
			missed = append(missed, k)
		}
	}
	slices.SortStableFunc(missed, func(x, y rune) int { return b.mistakes[y] - b.mistakes[x] })
	return missed[:min(n, len(missed))]
}
