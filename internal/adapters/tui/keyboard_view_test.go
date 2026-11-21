package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func TestKeyboardShowsEveryRow(t *testing.T) {
	view := renderKeyboard(0)

	for _, key := range []string{"[1]", "[q]", "[a]", "[;]", "[z]", "[/]", "[`]", "[\\]", "space"} {
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

func TestTildeMarksTheBacktickKeyAndTheRightShift(t *testing.T) {
	view := renderKeyboard('~')

	if !strings.Contains(view, "‹`›") || !strings.HasSuffix(bottomRow(view), "‹shift›") {
		t.Fatalf("` and right shift should be marked:\n%s", view)
	}
}

func TestPipeMarksTheBackslashKey(t *testing.T) {
	if view := renderKeyboard('|'); !strings.Contains(view, "‹\\›") {
		t.Fatalf("\\ should be marked:\n%s", view)
	}
}

func TestBacktickSitsNextToOne(t *testing.T) {
	if view := renderKeyboard(0); !strings.Contains(view, "[`][1]") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestEveryKeyTakesTheColourOfItsFinger(t *testing.T) {
	for finger, colour := range fingerColors {
		idle, pressed := keyCapStyle(finger, false), keyCapStyle(finger, true)
		if idle.GetForeground() != colour || !idle.GetFaint() {
			t.Errorf("%v idle key = %v faint %v, want %v", finger, idle.GetForeground(), idle.GetFaint(), colour)
		}
		if pressed.GetBackground() != colour || !pressed.GetBold() {
			t.Errorf("%v pressed key background = %v, want %v", finger, pressed.GetBackground(), colour)
		}
	}
}

func TestTheKeyToPressShowsTheColourOfItsFinger(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	t.Cleanup(func() { lipgloss.SetColorProfile(termenv.Ascii) })

	for press, background := range map[rune]string{'g': "48;2;239;68;68", 'h': "48;2;239;68;68", 'a': "48;2;59;130;246", ' ': "48;2;147;163;184"} {
		if view := renderKeyboard(press); !strings.Contains(view, background) {
			t.Errorf("%q is not shown in its finger's colour", press)
		}
	}
}
