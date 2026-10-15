package keyboard

var fingerNames = [...]string{
	LeftPinky:   "left pinky",
	LeftRing:    "left ring",
	LeftMiddle:  "left middle",
	LeftIndex:   "left index",
	LeftThumb:   "left thumb",
	RightThumb:  "right thumb",
	RightIndex:  "right index",
	RightMiddle: "right middle",
	RightRing:   "right ring",
	RightPinky:  "right pinky",
}

func (f Finger) String() string {
	return fingerNames[f]
}

func (h Hand) String() string {
	if h == Left {
		return "left"
	}
	return "right"
}
