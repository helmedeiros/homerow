package tui

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func markerColumns(view string) []int {
	var columns []int
	for _, line := range strings.Split(view, "\n") {
		if !strings.ContainsRune(line, '▲') {
			continue
		}
		for i, r := range []rune(line) {
			if r == '▲' {
				columns = append(columns, i)
			}
		}
	}
	return columns
}

func markerColumn(t *testing.T, view string) int {
	t.Helper()
	columns := markerColumns(view)
	if len(columns) == 0 {
		t.Fatalf("no marker in:\n%s", view)
	}
	return columns[0]
}

func key(r rune) keyboard.Key {
	k, _ := keyboard.Lookup(r)
	return k
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
		view := renderHands(key(c.label), true)
		if markerColumn(t, view) != labelColumn(t, view, c.label) {
			t.Errorf("%v: marker not under %q:\n%s", c.finger, c.label, view)
		}
		if !strings.Contains(view, c.finger.String()) {
			t.Errorf("caption %q missing:\n%s", c.finger, view)
		}
	}
}

func TestHandsWithoutAFingerHaveTheSameHeight(t *testing.T) {
	with := renderHands(key('F'), true)
	without := renderHands(key(' '), false)

	if strings.Contains(without, "▲") {
		t.Fatalf("unexpected marker:\n%s", without)
	}
	if strings.Count(with, "\n") != strings.Count(without, "\n") {
		t.Fatalf("heights differ:\n%s\n---\n%s", with, without)
	}
}

func TestCapitalsPointAtTheKeyAndTheOppositeShift(t *testing.T) {
	view := renderHands(key('F'), true)

	columns := markerColumns(view)
	want := []int{labelColumn(t, view, 'f'), labelColumn(t, view, ';')}
	if len(columns) != 2 || columns[0] != want[0] || columns[1] != want[1] {
		t.Fatalf("markers at %v, want %v:\n%s", columns, want, view)
	}
	if !strings.Contains(view, "left index + right pinky on shift") {
		t.Fatalf("caption missing:\n%s", view)
	}
}
