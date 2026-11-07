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

func bottomRow(view string) string {
	for _, line := range strings.Split(view, "\n") {
		if strings.Contains(line, "shift") {
			return line
		}
	}
	return ""
}

func TestKeyboardShowsTabAndBothShifts(t *testing.T) {
	view := renderKeyboard(0)

	if !strings.Contains(view, "[tab]") || strings.Count(bottomRow(view), "[shift]") != 2 {
		t.Fatalf("view:\n%s", view)
	}
}

func TestCapitalsMarkTheirKeyAndTheOppositeShift(t *testing.T) {
	row := bottomRow(renderKeyboard('F'))

	if !strings.HasPrefix(strings.TrimSpace(row), "[shift]") || !strings.HasSuffix(row, "‹shift›") {
		t.Fatalf("right shift should be marked:\n%s", row)
	}
	if view := renderKeyboard('F'); !strings.Contains(view, "‹f›") {
		t.Fatalf("f should be marked:\n%s", view)
	}
}

func TestShiftedSymbolsMarkTheirBaseKeyAndShift(t *testing.T) {
	view := renderKeyboard(':')

	if !strings.Contains(view, "‹;›") || !strings.HasPrefix(strings.TrimSpace(bottomRow(view)), "‹shift›") {
		t.Fatalf("; and left shift should be marked:\n%s", view)
	}
}

func TestKeyboardMarksTab(t *testing.T) {
	if view := renderKeyboard('\t'); !strings.Contains(view, "‹tab›") {
		t.Fatalf("tab should be marked:\n%s", view)
	}
}
