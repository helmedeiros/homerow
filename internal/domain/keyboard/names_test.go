package keyboard_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

func TestFingerNamesReadAsPlainWords(t *testing.T) {
	cases := map[keyboard.Finger]string{
		keyboard.LeftPinky:   "left pinky",
		keyboard.LeftRing:    "left ring",
		keyboard.LeftMiddle:  "left middle",
		keyboard.LeftIndex:   "left index",
		keyboard.LeftThumb:   "left thumb",
		keyboard.RightThumb:  "right thumb",
		keyboard.RightIndex:  "right index",
		keyboard.RightMiddle: "right middle",
		keyboard.RightRing:   "right ring",
		keyboard.RightPinky:  "right pinky",
	}
	for finger, want := range cases {
		if got := finger.String(); got != want {
			t.Errorf("String() = %q, want %q", got, want)
		}
	}
}

func TestHandNames(t *testing.T) {
	if keyboard.Left.String() != "left" || keyboard.Right.String() != "right" {
		t.Errorf("hand names = %q, %q", keyboard.Left, keyboard.Right)
	}
}
