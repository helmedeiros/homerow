package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestTheEnterKeyTypesANewLine(t *testing.T) {
	var s screen = newDrillScreen("a;\nb", 0, newClock().Now)
	s, _ = s.Update(runes("a;"))
	if !strings.Contains(s.View(), "right pinky") {
		t.Fatalf("the hands should point at enter:\n%s", s.View())
	}

	s, _ = s.Update(special(tea.KeyEnter))
	_, cmd := s.Update(runes("b"))

	if done, ok := message(cmd).(stepDoneMsg); !ok || done.result.Accuracy != 100 {
		t.Fatalf("message = %#v", message(cmd))
	}
}

func TestCodeKeepsItsLinesAndMarksEnter(t *testing.T) {
	d := newDrillScreen("if ok {\n\treturn\n}", 0, newClock().Now)

	lines := strings.Split(d.renderText(), "\n")

	if len(lines) != 3 || lines[0] != "if ok {↵" || lines[1] != "    return↵" || lines[2] != "}" {
		t.Fatalf("lines = %q", lines)
	}
}

func TestNestedIndentationIsAllSpaces(t *testing.T) {
	d := newDrillScreen("a {\n\tb {\n\t\tc\n\t}\n}", 0, newClock().Now)

	lines := strings.Split(d.renderText(), "\n")

	if lines[2] != "        c↵" || lines[3] != "    }↵" {
		t.Fatalf("lines = %q", lines)
	}
}

func TestATabInsideALineIsStillAnArrow(t *testing.T) {
	d := newDrillScreen("a\tb\n\tc", 0, newClock().Now)

	lines := strings.Split(d.renderText(), "\n")

	if lines[0] != "a→b↵" || lines[1] != "    c" {
		t.Fatalf("lines = %q", lines)
	}
}

func TestACodeStepShowsItsNoteAboveTheCode(t *testing.T) {
	lesson := curriculum.Lesson{Kind: curriculum.Code}
	step := curriculum.Step{Kind: curriculum.DrillStep, Text: "a := 1", Page: "Declare and assign at once."}

	view := Model{}.screenFor(lesson, step).View()

	note, code := strings.Index(view, "Declare and assign at once."), strings.Index(view, "a := 1")
	if note < 0 || code < 0 || note > code {
		t.Fatalf("view:\n%s", view)
	}
}

func TestADrillWithoutANoteStartsWithTheText(t *testing.T) {
	view := newDrillScreen("fj", 0, newClock().Now).View()

	if !strings.HasPrefix(view, "fj") {
		t.Fatalf("view:\n%s", view)
	}
}
