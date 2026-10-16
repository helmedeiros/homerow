package drill_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

func TestAccuracyIsPerfectBeforeTyping(t *testing.T) {
	if got := drill.NewSession("fj").Accuracy(); got != 100 {
		t.Fatalf("Accuracy() = %v, want 100", got)
	}
}

func TestAccuracyIsTheShareOfCorrectKeystrokes(t *testing.T) {
	s := drill.NewSession("fjfj")
	typeAll(s, "fjdj")

	if got := s.Accuracy(); got != 75 {
		t.Fatalf("Accuracy() = %v, want 75", got)
	}
	if s.Keystrokes() != 4 || s.Errors() != 1 {
		t.Fatalf("Keystrokes() = %d, Errors() = %d", s.Keystrokes(), s.Errors())
	}
}

func TestKeysAfterTheEndDoNotCount(t *testing.T) {
	s := drill.NewSession("f")
	typeAll(s, "fxx")

	if s.Keystrokes() != 1 || s.Accuracy() != 100 {
		t.Fatalf("Keystrokes() = %d, Accuracy() = %v", s.Keystrokes(), s.Accuracy())
	}
}
