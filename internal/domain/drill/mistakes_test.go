package drill_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

func TestMistakesCountTheKeysThatWereMissed(t *testing.T) {
	s := drill.NewSession("fjfj")
	typeAll(s, "x")
	s.Back()
	typeAll(s, "fk")
	s.Back()
	typeAll(s, "jfj")

	got := s.Mistakes()
	if len(got) != 2 || got['j'] != 1 || got['f'] != 1 {
		t.Fatalf("mistakes = %v", got)
	}
}

func TestACleanRunHasNoMistakes(t *testing.T) {
	s := drill.NewSession("fj")
	typeAll(s, "fj")

	if len(s.Mistakes()) != 0 {
		t.Fatalf("mistakes = %v", s.Mistakes())
	}
}
