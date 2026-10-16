package drill

import "time"

type Mark int

const (
	Pending Mark = iota
	Correct
	Wrong
)

const charactersPerWord = 5

type Session struct {
	target []rune
	marks  []Mark
	pos    int
	errors int
	first  time.Time
	last   time.Time
}

func NewSession(text string) *Session {
	target := []rune(text)
	return &Session{target: target, marks: make([]Mark, len(target))}
}

func (s *Session) Type(r rune, at time.Time) bool {
	if s.Done() {
		return false
	}
	if s.pos == 0 {
		s.first = at
	}
	s.last = at
	correct := r == s.target[s.pos]
	if correct {
		s.marks[s.pos] = Correct
	} else {
		s.marks[s.pos] = Wrong
		s.errors++
	}
	s.pos++
	return correct
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

func (s *Session) Keystrokes() int { return s.pos }

func (s *Session) Errors() int { return s.errors }

func (s *Session) Accuracy() float64 {
	if s.pos == 0 {
		return 100
	}
	return float64(s.pos-s.errors) / float64(s.pos) * 100
}

func (s *Session) Elapsed() time.Duration {
	return s.last.Sub(s.first)
}

func (s *Session) WPM() float64 {
	minutes := s.Elapsed().Minutes()
	if minutes <= 0 {
		return 0
	}
	return float64(s.pos-s.errors) / charactersPerWord / minutes
}
