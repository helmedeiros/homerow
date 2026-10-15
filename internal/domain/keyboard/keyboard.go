package keyboard

import "unicode"

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
	NumberRow Row = iota
	TopRow
	HomeRow
	BottomRow
	SpaceRow
)

type Key struct {
	Char   rune
	Finger Finger
	Row    Row
}

var rows = []string{"1234567890-=", "qwertyuiop[]", "asdfghjkl;'", "zxcvbnm,./"}

var columnFingers = []Finger{
	LeftPinky, LeftRing, LeftMiddle, LeftIndex, LeftIndex,
	RightIndex, RightIndex, RightMiddle, RightRing, RightPinky,
}

var homeKeys = map[Finger]rune{
	LeftPinky: 'a', LeftRing: 's', LeftMiddle: 'd', LeftIndex: 'f',
	LeftThumb: ' ', RightThumb: ' ',
	RightIndex: 'j', RightMiddle: 'k', RightRing: 'l', RightPinky: ';',
}

var keys = buildKeys()

func buildKeys() map[rune]Key {
	built := map[rune]Key{' ': {Char: ' ', Finger: LeftThumb, Row: SpaceRow}}
	for row, chars := range rows {
		for column, char := range chars {
			built[char] = Key{Char: char, Finger: fingerForColumn(column), Row: Row(row)}
		}
	}
	return built
}

func fingerForColumn(column int) Finger {
	if column < len(columnFingers) {
		return columnFingers[column]
	}
	return RightPinky
}

func Lookup(char rune) (Key, bool) {
	key, ok := keys[unicode.ToLower(char)]
	if ok {
		key.Char = char
	}
	return key, ok
}

func Rows() []string {
	return append([]string(nil), rows...)
}

func HomeKey(f Finger) rune {
	return homeKeys[f]
}
