package keyboard_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func TestShiftedSymbolsSitOnTheirBaseKey(t *testing.T) {
	cases := []struct {
		char   rune
		base   rune
		finger keyboard.Finger
		shift  keyboard.Finger
	}{
		{'!', '1', keyboard.LeftPinky, keyboard.RightPinky},
		{'@', '2', keyboard.LeftRing, keyboard.RightPinky},
		{'#', '3', keyboard.LeftMiddle, keyboard.RightPinky},
		{'$', '4', keyboard.LeftIndex, keyboard.RightPinky},
		{'%', '5', keyboard.LeftIndex, keyboard.RightPinky},
		{'^', '6', keyboard.RightIndex, keyboard.LeftPinky},
		{'&', '7', keyboard.RightIndex, keyboard.LeftPinky},
		{'*', '8', keyboard.RightMiddle, keyboard.LeftPinky},
		{'(', '9', keyboard.RightRing, keyboard.LeftPinky},
		{')', '0', keyboard.RightPinky, keyboard.LeftPinky},
		{':', ';', keyboard.RightPinky, keyboard.LeftPinky},
		{'?', '/', keyboard.RightPinky, keyboard.LeftPinky},
		{'<', ',', keyboard.RightMiddle, keyboard.LeftPinky},
		{'>', '.', keyboard.RightRing, keyboard.LeftPinky},
		{'"', '\'', keyboard.RightPinky, keyboard.LeftPinky},
	}
	for _, c := range cases {
		key, ok := keyboard.Lookup(c.char)
		shift, shifted := key.ShiftFinger()
		if !ok || key.Char != c.char || key.Base != c.base || key.Finger != c.finger || !shifted || shift != c.shift {
			t.Errorf("Lookup(%q) = %+v, %v; shift %v %v", c.char, key, ok, shift, shifted)
		}
	}
}

func TestTabBelongsToTheLeftPinky(t *testing.T) {
	key, ok := keyboard.Lookup('\t')

	if !ok || key.Finger != keyboard.LeftPinky || key.Row != keyboard.TopRow || key.Shifted {
		t.Fatalf("Lookup(tab) = %+v, %v", key, ok)
	}
}
