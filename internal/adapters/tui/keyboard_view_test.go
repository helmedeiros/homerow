package tui

import (
	"strings"
	"testing"
)

func TestKeyboardShowsEveryRow(t *testing.T) {
	view := renderKeyboard(0)

	for _, key := range []string{"[1]", "[q]", "[a]", "[;]", "[z]", "[/]", "space"} {
		if !strings.Contains(view, key) {
			t.Errorf("keyboard is missing %s:\n%s", key, view)
		}
	}
}

func TestKeyboardMarksTheKeyToPress(t *testing.T) {
	view := renderKeyboard('f')

	if !strings.Contains(view, "‹f›") || strings.Contains(view, "[f]") {
		t.Fatalf("f should be marked:\n%s", view)
	}
}

func TestKeyboardMarksTheSpaceBar(t *testing.T) {
	view := renderKeyboard(' ')

	if !strings.Contains(view, "‹") || !strings.Contains(view, "space") {
		t.Fatalf("space should be marked:\n%s", view)
	}
}
