package progress

import "maps"

type Entry struct {
	Completed bool
	Stars     int
	WPM       float64
	Accuracy  float64
}

type Board struct {
	entries map[int]Entry
}

func NewBoard(entries map[int]Entry) *Board {
	b := &Board{entries: map[int]Entry{}}
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
