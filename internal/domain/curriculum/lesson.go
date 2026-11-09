package curriculum

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInvalidLesson = errors.New("invalid lesson")

type Kind int

const (
	Intro Kind = iota
	NewKeys
	Review
	Practice
	HandDrill
	Play
	Travel
)

var kindNames = [...]string{
	Intro:     "intro",
	NewKeys:   "new keys",
	Review:    "review",
	Practice:  "practice",
	HandDrill: "hand drill",
	Play:      "play",
	Travel:    "travel",
}

func (k Kind) String() string { return kindNames[k] }

func (k Kind) Scored() bool { return k >= Review }

type Game int

const (
	FallingLetters Game = iota
	BalloonWords
	PopUpWarmUp
)

type Lesson struct {
	Number    int
	Title     string
	Kind      Kind
	Pages     []string
	NewKeys   []rune
	Keys      []rune
	Focus     []rune
	Words     bool
	Game      Game
	GoalWPM   float64
	TimeLimit time.Duration
}

func (l Lesson) Validate() error {
	switch {
	case l.Title == "":
		return l.invalid("has no title")
	case l.Kind == Intro && len(l.Pages) == 0:
		return l.invalid("is an intro without pages")
	case l.Kind != Intro && len(l.Keys) == 0:
		return l.invalid("has no keys to practise")
	case !subset(l.NewKeys, l.Keys):
		return l.invalid("introduces keys it does not allow")
	case !subset(l.Focus, l.Keys):
		return l.invalid("focuses on keys it does not allow")
	case l.Kind == Travel && len(l.Focus) == 0:
		return l.invalid("travels without a home key to start from")
	case l.Kind.Scored() && l.GoalWPM <= 0:
		return l.invalid("is scored without a speed goal")
	}
	return nil
}

func (l Lesson) invalid(reason string) error {
	return fmt.Errorf("%w: lesson %d %s", ErrInvalidLesson, l.Number, reason)
}

func subset(part, whole []rune) bool {
	for _, r := range part {
		if !strings.ContainsRune(string(whole), r) {
			return false
		}
	}
	return true
}

type Module struct {
	Number  int
	Title   string
	Lessons []Lesson
}

func (m Module) Lesson(number int) (Lesson, bool) {
	for _, l := range m.Lessons {
		if l.Number == number {
			return l, true
		}
	}
	return Lesson{}, false
}
