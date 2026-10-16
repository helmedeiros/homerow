package drill_test

import (
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

var start = time.Date(2026, 10, 16, 20, 0, 0, 0, time.UTC)

func typeAll(s *drill.Session, text string) {
	for _, r := range text {
		s.Type(r, start)
	}
}

func TestNewSessionStartsAtTheFirstCharacter(t *testing.T) {
	s := drill.NewSession("fj fj")

	want, ok := s.Expected()
	if !ok || want != 'f' {
		t.Fatalf("Expected() = %q, %v, want 'f', true", want, ok)
	}
	if s.Position() != 0 || s.Done() {
		t.Fatalf("Position() = %d, Done() = %v", s.Position(), s.Done())
	}
	for i, m := range s.Marks() {
		if m != drill.Pending {
			t.Errorf("Marks()[%d] = %v, want Pending", i, m)
		}
	}
}

func TestCorrectKeyMarksTheCharacterAndAdvances(t *testing.T) {
	s := drill.NewSession("fj")

	if !s.Type('f', start) {
		t.Fatal("Type('f') = false, want true")
	}

	if s.Marks()[0] != drill.Correct || s.Position() != 1 {
		t.Fatalf("mark = %v, position = %d", s.Marks()[0], s.Position())
	}
}

func TestWrongKeyMarksTheCharacterAndStillAdvances(t *testing.T) {
	s := drill.NewSession("fj")

	if s.Type('d', start) {
		t.Fatal("Type('d') = true, want false")
	}

	if s.Marks()[0] != drill.Wrong || s.Position() != 1 {
		t.Fatalf("mark = %v, position = %d", s.Marks()[0], s.Position())
	}
}

func TestSessionIsDoneAfterTheLastCharacter(t *testing.T) {
	s := drill.NewSession("fj")
	typeAll(s, "fj")

	if !s.Done() {
		t.Fatal("Done() = false after typing every character")
	}
	if _, ok := s.Expected(); ok {
		t.Fatal("Expected() should report nothing left")
	}
}

func TestKeysAfterTheEndAreIgnored(t *testing.T) {
	s := drill.NewSession("f")
	typeAll(s, "fx")

	if s.Position() != 1 || len(s.Marks()) != 1 {
		t.Fatalf("position = %d, marks = %d", s.Position(), len(s.Marks()))
	}
}

func TestTextIsKept(t *testing.T) {
	if got := drill.NewSession("asdf").Text(); got != "asdf" {
		t.Fatalf("Text() = %q", got)
	}
}
