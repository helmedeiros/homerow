package tui

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func TestAverageAddsUpTheMistakesOfEveryStep(t *testing.T) {
	got := average([]scoring.Result{
		{Accuracy: 90, WPM: 10, Mistakes: map[rune]int{'f': 1}},
		{Accuracy: 100, WPM: 20, Mistakes: map[rune]int{'f': 2, 'j': 1}},
		{Accuracy: 95, WPM: 15},
	})

	if got.Accuracy != 95 || got.WPM != 15 || got.Mistakes['f'] != 3 || got.Mistakes['j'] != 1 {
		t.Fatalf("average = %+v", got)
	}
}
