package drill_test

import (
	"slices"
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

func TestBackspaceClearsThePreviousCharacter(t *testing.T) {
	s := drill.NewSession("fjf")
	typeAll(s, "fx")

	if !s.Back() {
		t.Fatal("Back() = false, want true")
	}

	if s.Position() != 1 || !slices.Equal(s.Marks(), []drill.Mark{drill.Correct, drill.Pending, drill.Pending}) {
		t.Fatalf("Position() = %d, Marks() = %v", s.Position(), s.Marks())
	}
	if want, _ := s.Expected(); want != 'j' {
		t.Fatalf("Expected() = %q, want 'j'", want)
	}
}

func TestBackspaceAtTheStartDoesNothing(t *testing.T) {
	s := drill.NewSession("fj")

	if s.Back() || s.Position() != 0 {
		t.Fatalf("Back() moved to %d", s.Position())
	}
}

func TestFixedMistakeIsMarkedCorrected(t *testing.T) {
	s := drill.NewSession("fj")
	typeAll(s, "fx")
	s.Back()
	typeAll(s, "j")

	if !s.Done() || !slices.Equal(s.Marks(), []drill.Mark{drill.Correct, drill.Corrected}) {
		t.Fatalf("Done() = %v, Marks() = %v", s.Done(), s.Marks())
	}
}

func TestMistakeRetypedWrongStaysWrong(t *testing.T) {
	s := drill.NewSession("fj")
	typeAll(s, "fx")
	s.Back()
	typeAll(s, "y")

	if s.Marks()[1] != drill.Wrong {
		t.Fatalf("Marks() = %v", s.Marks())
	}
}

func TestErasingACorrectCharacterKeepsItCorrectWhenRetyped(t *testing.T) {
	s := drill.NewSession("fj")
	typeAll(s, "f")
	s.Back()
	typeAll(s, "f")

	if s.Marks()[0] != drill.Correct {
		t.Fatalf("Marks() = %v", s.Marks())
	}
}

func TestFixedMistakesStillCountAgainstAccuracy(t *testing.T) {
	s := drill.NewSession("fj")
	typeAll(s, "fx")
	s.Back()
	typeAll(s, "j")

	if s.Keystrokes() != 3 || s.Errors() != 1 {
		t.Fatalf("Keystrokes() = %d, Errors() = %d", s.Keystrokes(), s.Errors())
	}
	if got := s.Accuracy(); got < 66.6 || got > 66.7 {
		t.Fatalf("Accuracy() = %v, want about 66.7", got)
	}
}

func TestFixedMistakesCountTowardsSpeed(t *testing.T) {
	s := drill.NewSession("fjfjfjfjfj")
	s.Type('x', start)
	s.Back()
	for _, r := range "fjfjfjfjfj" {
		s.Type(r, start.Add(6*time.Second))
	}

	if got := s.WPM(); got != 20 {
		t.Fatalf("WPM() = %v, want 20", got)
	}
}
