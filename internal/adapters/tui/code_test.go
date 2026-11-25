package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
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
