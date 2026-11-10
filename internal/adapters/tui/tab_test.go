package tui

import (
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestATabShowsAsAnArrow(t *testing.T) {
	d := newDrillScreen("sad\tlad", 0, newClock().Now)

	view := d.View()

	if !strings.Contains(view, "→") || strings.Contains(view, "\t") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestTheTabKeyTypesATab(t *testing.T) {
	var s screen = newDrillScreen("a\tb", time.Duration(0), newClock().Now)
	s, _ = s.Update(runes("a"))
	if !strings.Contains(s.View(), "left pinky") {
		t.Fatalf("the hands should point at tab:\n%s", s.View())
	}

	s, _ = s.Update(special(tea.KeyTab))
	_, cmd := s.Update(runes("b"))

	if done, ok := message(cmd).(stepDoneMsg); !ok || done.result.Accuracy != 100 {
		t.Fatalf("message = %#v", message(cmd))
	}
}

func TestTheCursorOnATabIsAnArrow(t *testing.T) {
	var s screen = newDrillScreen("a\tb", 0, newClock().Now)
	s, _ = s.Update(runes("a"))

	if !strings.Contains(s.View(), cursorStyle.Render("→")) {
		t.Fatalf("view:\n%s", s.View())
	}
}
