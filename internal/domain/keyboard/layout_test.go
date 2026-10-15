package keyboard_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func TestEveryRowKeyHasAFingerAndRow(t *testing.T) {
	cases := []struct {
		char   rune
		finger keyboard.Finger
		row    keyboard.Row
	}{
		{'1', keyboard.LeftPinky, keyboard.NumberRow},
		{'5', keyboard.LeftIndex, keyboard.NumberRow},
		{'6', keyboard.RightIndex, keyboard.NumberRow},
		{'=', keyboard.RightPinky, keyboard.NumberRow},
		{'q', keyboard.LeftPinky, keyboard.TopRow},
		{'t', keyboard.LeftIndex, keyboard.TopRow},
		{'y', keyboard.RightIndex, keyboard.TopRow},
		{'p', keyboard.RightPinky, keyboard.TopRow},
		{'\'', keyboard.RightPinky, keyboard.HomeRow},
		{'z', keyboard.LeftPinky, keyboard.BottomRow},
		{'b', keyboard.LeftIndex, keyboard.BottomRow},
		{'n', keyboard.RightIndex, keyboard.BottomRow},
		{'/', keyboard.RightPinky, keyboard.BottomRow},
		{' ', keyboard.LeftThumb, keyboard.SpaceRow},
	}
	for _, c := range cases {
		key, ok := keyboard.Lookup(c.char)
		if !ok {
			t.Fatalf("Lookup(%q) not found", c.char)
		}
		if key.Finger != c.finger || key.Row != c.row {
			t.Errorf("Lookup(%q) = %v on row %d, want %v on row %d", c.char, key.Finger, key.Row, c.finger, c.row)
		}
	}
}

func TestUppercaseLettersUseTheSameFinger(t *testing.T) {
	key, ok := keyboard.Lookup('F')
	if !ok || key.Finger != keyboard.LeftIndex {
		t.Errorf("Lookup('F') = %v, %v", key, ok)
	}
}

func TestRowsListTheKeysInOrder(t *testing.T) {
	rows := keyboard.Rows()
	want := []string{"1234567890-=", "qwertyuiop[]", "asdfghjkl;'", "zxcvbnm,./"}
	if len(rows) != len(want) {
		t.Fatalf("len(Rows()) = %d, want %d", len(rows), len(want))
	}
	for i := range want {
		if rows[i] != want[i] {
			t.Errorf("Rows()[%d] = %q, want %q", i, rows[i], want[i])
		}
	}
}

func TestHomeKeyIsWhereAFingerRests(t *testing.T) {
	cases := map[keyboard.Finger]rune{
		keyboard.LeftPinky:  'a',
		keyboard.LeftIndex:  'f',
		keyboard.RightIndex: 'j',
		keyboard.RightPinky: ';',
		keyboard.LeftThumb:  ' ',
	}
	for finger, want := range cases {
		if got := keyboard.HomeKey(finger); got != want {
			t.Errorf("HomeKey(%v) = %q, want %q", finger, got, want)
		}
	}
}
