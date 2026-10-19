package scoring_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func TestStarsFollowAccuracyBands(t *testing.T) {
	const goal = 20
	cases := []struct {
		accuracy float64
		wpm      float64
		want     int
	}{
		{79.9, 40, 0},
		{80, 5, 1},
		{87.9, 5, 1},
		{88, 5, 2},
		{94, 5, 3},
		{97, 19.9, 4},
		{97, 20, 5},
		{100, 35, 5},
	}
	for _, c := range cases {
		got := scoring.Stars(scoring.Result{Accuracy: c.accuracy, WPM: c.wpm}, goal)
		if got != c.want {
			t.Errorf("Stars(%v%%, %v wpm) = %d, want %d", c.accuracy, c.wpm, got, c.want)
		}
	}
}

func TestAnyStarPasses(t *testing.T) {
	if scoring.Passed(0) || !scoring.Passed(1) {
		t.Fatal("Passed should be true from one star")
	}
}
