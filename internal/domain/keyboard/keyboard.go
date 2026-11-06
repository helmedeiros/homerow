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
	Char    rune
	Base    rune
	Shifted bool
	Finger  Finger
	Row     Row
}

func (k Key) ShiftFinger() (Finger, bool) {
	if !k.Shifted {
		return 0, false
	}
	if k.Finger.Hand() == Left {
		return RightPinky, true
	}
	return LeftPinky, true
}

var rows = []string{"1234567890-=", "qwertyuiop[]", "asdfghjkl;'", "zxcvbnm,./"}

var shiftedRows = []string{"!@#$%^&*()_+", "          {}", "         :\"", "       <>?"}

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
	built := map[rune]Key{' ': {Char: ' ', Base: ' ', Finger: LeftThumb, Row: SpaceRow}}
	built['\t'] = Key{Char: '\t', Base: '\t', Finger: LeftPinky, Row: TopRow}
	for row, chars := range rows {
		for column, char := range chars {
			built[char] = Key{Char: char, Base: char, Finger: fingerForColumn(column), Row: Row(row)}
		}
	}
	for row, chars := range shiftedRows {
		for column, char := range chars {
			if char != ' ' {
				base := rune(rows[row][column])
				built[char] = Key{Char: char, Base: base, Shifted: true, Finger: fingerForColumn(column), Row: Row(row)}
			}
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
	if ok && char != key.Char {
		key.Char, key.Shifted = char, true
	}
	return key, ok
}

func Rows() []string {
	return append([]string(nil), rows...)
}

func HomeKey(f Finger) rune {
	return homeKeys[f]
}
