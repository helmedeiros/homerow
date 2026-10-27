package tui

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func markerColumn(t *testing.T, view string) int {
	t.Helper()
	for _, line := range strings.Split(view, "\n") {
		if i := strings.IndexRune(line, '▲'); i >= 0 {
			return len([]rune(line[:i]))
		}
	}
	t.Fatalf("no marker in:\n%s", view)
	return -1
}

func labelColumn(t *testing.T, view string, label rune) int {
	t.Helper()
	lines := strings.Split(view, "\n")
	for _, line := range lines {
		if strings.Contains(line, "a s d f") {
			return len([]rune(line[:strings.IndexRune(line, label)]))
		}
	}
	t.Fatalf("no home keys in:\n%s", view)
	return -1
}

func TestHandsPointAtTheFingerToUse(t *testing.T) {
	for _, c := range []struct {
		finger keyboard.Finger
		label  rune
	}{
		{keyboard.LeftIndex, 'f'},
		{keyboard.RightPinky, ';'},
		{keyboard.LeftPinky, 'a'},
	} {
		view := renderHands(c.finger, true)
		if markerColumn(t, view) != labelColumn(t, view, c.label) {
			t.Errorf("%v: marker not under %q:\n%s", c.finger, c.label, view)
		}
		if !strings.Contains(view, c.finger.String()) {
			t.Errorf("caption %q missing:\n%s", c.finger, view)
		}
	}
}

func TestHandsWithoutAFingerHaveTheSameHeight(t *testing.T) {
	with := renderHands(keyboard.LeftThumb, true)
	without := renderHands(keyboard.LeftThumb, false)

	if strings.Contains(without, "▲") {
		t.Fatalf("unexpected marker:\n%s", without)
	}
	if strings.Count(with, "\n") != strings.Count(without, "\n") {
		t.Fatalf("heights differ:\n%s\n---\n%s", with, without)
	}
}
