package keyboard_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func TestCapitalsNeedShiftFromTheOtherHand(t *testing.T) {
	cases := []struct {
		char  rune
		base  rune
		shift keyboard.Finger
	}{
		{'F', 'f', keyboard.RightPinky},
		{'A', 'a', keyboard.RightPinky},
		{'T', 't', keyboard.RightPinky},
		{'J', 'j', keyboard.LeftPinky},
		{'P', 'p', keyboard.LeftPinky},
		{'N', 'n', keyboard.LeftPinky},
	}
	for _, c := range cases {
		key, _ := keyboard.Lookup(c.char)
		shift, ok := key.ShiftFinger()
		if !key.Shifted || key.Base != c.base || !ok || shift != c.shift {
			t.Errorf("Lookup(%q) = %+v, shift %v %v", c.char, key, shift, ok)
		}
	}
}

func TestLowercaseNeedsNoShift(t *testing.T) {
	key, _ := keyboard.Lookup('f')

	if _, ok := key.ShiftFinger(); key.Shifted || ok || key.Base != 'f' {
		t.Errorf("Lookup('f') = %+v", key)
	}
}
