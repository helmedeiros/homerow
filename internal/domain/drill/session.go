package drill

import "time"

type Mark int

const (
	Pending Mark = iota
	Correct
	Wrong
	Corrected
)

const charactersPerWord = 5

type Session struct {
	target     []rune
	marks      []Mark
	missed     []bool
	pos        int
	keystrokes int
	errors     int
	first      time.Time
	last       time.Time
}

func NewSession(text string) *Session {
	target := []rune(text)
	return &Session{target: target, marks: make([]Mark, len(target)), missed: make([]bool, len(target))}
}

func (s *Session) Type(r rune, at time.Time) bool {
	if s.Done() {
		return false
	}
	if s.keystrokes == 0 {
		s.first = at
	}
	s.last = at
	s.keystrokes++
	correct := r == s.target[s.pos]
	switch {
	case !correct:
		s.marks[s.pos] = Wrong
		s.missed[s.pos] = true
		s.errors++
	case s.missed[s.pos]:
		s.marks[s.pos] = Corrected
	default:
		s.marks[s.pos] = Correct
	}
	s.pos++
	return correct
}

func (s *Session) Back() bool {
	if s.pos == 0 {
		return false
	}
	s.pos--
	s.marks[s.pos] = Pending
	return true
}

func (s *Session) Expected() (rune, bool) {
	if s.Done() {
		return 0, false
	}
	return s.target[s.pos], true
}

func (s *Session) Position() int { return s.pos }

func (s *Session) Done() bool { return s.pos >= len(s.target) }

func (s *Session) Marks() []Mark { return append([]Mark(nil), s.marks...) }

func (s *Session) Text() string { return string(s.target) }

func (s *Session) Keystrokes() int { return s.keystrokes }

func (s *Session) Errors() int { return s.errors }

func (s *Session) Accuracy() float64 {
	if s.keystrokes == 0 {
		return 100
	}
	return float64(s.keystrokes-s.errors) / float64(s.keystrokes) * 100
}

func (s *Session) Elapsed() time.Duration {
	return s.last.Sub(s.first)
}

func (s *Session) WPM() float64 {
	minutes := s.Elapsed().Minutes()
	if minutes <= 0 {
		return 0
	}
	return float64(s.typedRight()) / charactersPerWord / minutes
}

func (s *Session) typedRight() int {
	right := 0
	for _, m := range s.marks {
		if m == Correct || m == Corrected {
			right++
		}
	}
	return right
}
