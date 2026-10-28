package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestPageShowsItsText(t *testing.T) {
	view := newPageScreen("Rest your fingers on the home row.").View()

	if !strings.Contains(view, "Rest your fingers on the home row.") || !strings.Contains(view, "enter") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestPageFinishesOnEnterOrSpace(t *testing.T) {
	for _, key := range []tea.KeyMsg{special(tea.KeyEnter), special(tea.KeySpace), special(tea.KeyRight)} {
		_, cmd := newPageScreen("hi").Update(key)
		if _, ok := message(cmd).(stepDoneMsg); !ok {
			t.Errorf("%v: message = %#v, want step done", key, message(cmd))
		}
	}
}

func TestPageCanBeLeft(t *testing.T) {
	_, cmd := newPageScreen("hi").Update(special(tea.KeyEsc))

	if _, ok := message(cmd).(leaveLessonMsg); !ok {
		t.Fatalf("message = %#v, want leave", message(cmd))
	}
}

func TestPageIgnoresOtherKeys(t *testing.T) {
	if _, cmd := newPageScreen("hi").Update(runes("x")); cmd != nil {
		t.Fatalf("message = %#v", message(cmd))
	}
}
