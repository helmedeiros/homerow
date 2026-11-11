package storetest

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func Contract(t *testing.T, open func(t *testing.T) app.ProgressStore) {
	t.Run("starts empty", func(t *testing.T) {
		board, err := open(t).Load()
		if err != nil {
			t.Fatal(err)
		}
		if len(board.Entries()) != 0 {
			t.Fatalf("entries = %v", board.Entries())
		}
	})

	t.Run("returns what was saved", func(t *testing.T) {
		store := open(t)
		want := map[int]progress.Entry{
			1: {Completed: true},
			4: {Completed: true, Stars: 4, WPM: 11.5, Accuracy: 96.25},
		}
		if err := store.Save(progress.NewBoard(want)); err != nil {
			t.Fatal(err)
		}

		board, err := store.Load()
		if err != nil {
			t.Fatal(err)
		}
		got := board.Entries()
		if len(got) != len(want) || got[1] != want[1] || got[4] != want[4] {
			t.Fatalf("loaded %v, want %v", got, want)
		}
	})

	t.Run("later saves replace earlier ones", func(t *testing.T) {
		store := open(t)
		_ = store.Save(progress.NewBoard(map[int]progress.Entry{1: {Completed: true}}))
		_ = store.Save(progress.NewBoard(map[int]progress.Entry{2: {Stars: 1}}))

		board, _ := store.Load()
		if got := board.Entries(); len(got) != 1 || got[2].Stars != 1 {
			t.Fatalf("loaded %v", got)
		}
	})

	t.Run("keeps mistakes per key", func(t *testing.T) { keepsMistakes(t, open(t)) })
}

func keepsMistakes(t *testing.T, store app.ProgressStore) {
	t.Helper()
	board := progress.NewBoard(nil)
	board.AddMistakes(map[rune]int{'f': 2, '\t': 1})
	if err := store.Save(board); err != nil {
		t.Fatal(err)
	}

	loaded, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if got := loaded.Mistakes(); len(got) != 2 || got['f'] != 2 || got['\t'] != 1 {
		t.Fatalf("mistakes = %v", got)
	}
}
