package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestKeyIntroNamesTheKeyAndFinger(t *testing.T) {
	view := newKeyIntroScreen('f').View()

	for _, want := range []string{"NEW KEY", "Type", "left index", "‹f›", "▲"} {
		if !strings.Contains(view, want) {
			t.Errorf("missing %q:\n%s", want, view)
		}
	}
}

func TestKeyIntroNamesTheSpaceBar(t *testing.T) {
	view := newKeyIntroScreen(' ').View()

	if !strings.Contains(view, "space bar") || !strings.Contains(view, "thumb") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestKeyIntroFinishesWhenTheKeyIsPressed(t *testing.T) {
	_, cmd := newKeyIntroScreen('j').Update(runes("j"))

	if _, ok := message(cmd).(stepDoneMsg); !ok {
		t.Fatalf("message = %#v, want step done", message(cmd))
	}
}

func TestKeyIntroAcceptsTheSpaceBar(t *testing.T) {
	_, cmd := newKeyIntroScreen(' ').Update(special(tea.KeySpace))

	if _, ok := message(cmd).(stepDoneMsg); !ok {
		t.Fatalf("message = %#v, want step done", message(cmd))
	}
}

func TestKeyIntroHintsAfterAWrongKey(t *testing.T) {
	s, cmd := newKeyIntroScreen('f').Update(runes("d"))

	if cmd != nil {
		t.Fatalf("message = %#v", message(cmd))
	}
	if !strings.Contains(s.View(), "That was d") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestKeyIntroCanBeLeft(t *testing.T) {
	_, cmd := newKeyIntroScreen('f').Update(special(tea.KeyEsc))

	if _, ok := message(cmd).(leaveLessonMsg); !ok {
		t.Fatalf("message = %#v, want leave", message(cmd))
	}
}

func TestKeyIntroFinishesWhenABurstContainsTheKey(t *testing.T) {
	_, cmd := newKeyIntroScreen('j').Update(runes("dj"))

	if _, ok := message(cmd).(stepDoneMsg); !ok {
		t.Fatalf("message = %#v, want step done", message(cmd))
	}
}
