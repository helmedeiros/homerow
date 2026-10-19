package exercise_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func TestIntroduceRepeatsEachKeyThenMixesThem(t *testing.T) {
	got := exercise.Introduce([]rune("fj"))

	want := "f f f ff ff fff j j j jj jj jjj fj jf ffjj jjff fjfj jfjf"
	if got != want {
		t.Fatalf("Introduce(fj) =\n%q\nwant\n%q", got, want)
	}
}

func TestIntroduceASingleKeyOnlyRepeatsIt(t *testing.T) {
	if got := exercise.Introduce([]rune("g")); got != "g g g gg gg ggg" {
		t.Fatalf("Introduce(g) = %q", got)
	}
}
