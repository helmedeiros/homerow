package progress_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func TestMistakesAddUpPerKey(t *testing.T) {
	b := progress.NewBoard(nil)

	b.AddMistakes(map[rune]int{'f': 2, 'j': 1})
	b.AddMistakes(map[rune]int{'f': 3, 'k': 0})

	got := b.Mistakes()
	if len(got) != 2 || got['f'] != 5 || got['j'] != 1 {
		t.Fatalf("mistakes = %v", got)
	}
	got['f'] = 0
	if b.Mistakes()['f'] != 5 {
		t.Fatal("changing Mistakes() changed the board")
	}
}

func TestWeakestKeysComeFirst(t *testing.T) {
	b := progress.NewBoard(nil)
	b.AddMistakes(map[rune]int{'a': 1, 's': 4, 'd': 2, 'q': 9})

	if got := string(b.Weakest([]rune("asdfs"), 2)); got != "sd" {
		t.Fatalf("Weakest = %q, want sd", got)
	}
	if got := string(b.Weakest([]rune("asdf"), 5)); got != "sda" {
		t.Fatalf("Weakest = %q, want sda", got)
	}
	if got := b.Weakest([]rune("jkl"), 3); len(got) != 0 {
		t.Fatalf("Weakest = %q, want none", string(got))
	}
}
