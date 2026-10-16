package drill

import "time"

type Mark int

const (
	Pending Mark = iota
	Correct
	Wrong
)

type Session struct {
	target []rune
	marks  []Mark
	pos    int
}

func NewSession(text string) *Session {
	target := []rune(text)
	return &Session{target: target, marks: make([]Mark, len(target))}
}

func (s *Session) Type(r rune, _ time.Time) bool {
	if s.Done() {
		return false
	}
	correct := r == s.target[s.pos]
	if correct {
		s.marks[s.pos] = Correct
	} else {
		s.marks[s.pos] = Wrong
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
