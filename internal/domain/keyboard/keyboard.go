package keyboard

type Hand int

const (
	Left Hand = iota
	Right
)

type Finger int

const (
	LeftPinky Finger = iota
	LeftRing
	LeftMiddle
	LeftIndex
	LeftThumb
	RightThumb
	RightIndex
	RightMiddle
	RightRing
	RightPinky
)

func (f Finger) Hand() Hand {
	if f <= LeftThumb {
		return Left
	}
	return Right
}

type Row int

const (
	HomeRow Row = iota
)

type Key struct {
	Char   rune
	Finger Finger
	Row    Row
}

var keys = map[rune]Key{}

func init() {
	homeFingers := []Finger{LeftPinky, LeftRing, LeftMiddle, LeftIndex, LeftIndex, RightIndex, RightIndex, RightMiddle, RightRing, RightPinky}
	for i, char := range "asdfghjkl;" {
		keys[char] = Key{Char: char, Finger: homeFingers[i], Row: HomeRow}
	}
}

func Lookup(char rune) (Key, bool) {
	key, ok := keys[char]
	return key, ok
}
