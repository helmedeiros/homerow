package drill_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

func TestANewLineSkipsTheIndentationLikeAnEditor(t *testing.T) {
	s := drill.NewSession("if {\n\t\tx\n}")

	typeAll(s, "if {\n")

	if want, _ := s.Expected(); want != 'x' || s.Position() != 7 {
		t.Fatalf("Expected() = %q at %d", want, s.Position())
	}
	marks := s.Marks()
	if marks[5] != drill.Skipped || marks[6] != drill.Skipped {
		t.Fatalf("marks = %v", marks)
	}
	typeAll(s, "x\n}")
	if !s.Done() || s.Accuracy() != 100 || s.Keystrokes() != 8 {
		t.Fatalf("done %v, accuracy %v, keystrokes %d", s.Done(), s.Accuracy(), s.Keystrokes())
	}
}

func TestSpacesAfterANewLineAreSkippedToo(t *testing.T) {
	s := drill.NewSession("a\n    b")

	typeAll(s, "a\n")

	if want, _ := s.Expected(); want != 'b' {
		t.Fatalf("Expected() = %q", want)
	}
}

func TestSpacesInsideALineAreTyped(t *testing.T) {
	s := drill.NewSession("a b")

	typeAll(s, "a")

	if want, _ := s.Expected(); want != ' ' {
		t.Fatalf("Expected() = %q", want)
	}
}

func TestBackspaceStepsBackOverTheIndentationToTheNewLine(t *testing.T) {
	s := drill.NewSession("a\n\tb")
	typeAll(s, "ax")

	s.Back()

	if want, _ := s.Expected(); want != '\n' || s.Position() != 1 {
		t.Fatalf("Expected() = %q at %d", want, s.Position())
	}
	for i, m := range s.Marks()[1:] {
		if m != drill.Pending {
			t.Fatalf("Marks()[%d] = %v", i+1, m)
		}
	}
}
