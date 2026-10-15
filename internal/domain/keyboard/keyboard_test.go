package keyboard_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func TestHomeRowKeysBelongToTheirRestingFinger(t *testing.T) {
	cases := map[rune]keyboard.Finger{
		'a': keyboard.LeftPinky,
		's': keyboard.LeftRing,
		'd': keyboard.LeftMiddle,
		'f': keyboard.LeftIndex,
		'g': keyboard.LeftIndex,
		'h': keyboard.RightIndex,
		'j': keyboard.RightIndex,
		'k': keyboard.RightMiddle,
		'l': keyboard.RightRing,
		';': keyboard.RightPinky,
	}
	for char, want := range cases {
		key, ok := keyboard.Lookup(char)
		if !ok {
			t.Fatalf("Lookup(%q) not found", char)
		}
		if key.Finger != want {
			t.Errorf("Lookup(%q).Finger = %v, want %v", char, key.Finger, want)
		}
		if key.Row != keyboard.HomeRow {
			t.Errorf("Lookup(%q).Row = %v, want home row", char, key.Row)
		}
	}
}

func TestFingersBelongToAHand(t *testing.T) {
	if keyboard.LeftIndex.Hand() != keyboard.Left {
		t.Error("left index should be on the left hand")
	}
	if keyboard.RightPinky.Hand() != keyboard.Right {
		t.Error("right pinky should be on the right hand")
	}
}

func TestLookupRejectsUnknownCharacters(t *testing.T) {
	if _, ok := keyboard.Lookup('€'); ok {
		t.Error("Lookup('€') should not be found")
	}
}
