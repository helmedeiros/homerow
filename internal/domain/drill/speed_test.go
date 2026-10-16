package drill_test

import (
	"testing"
	"time"

	"github.com/helmedeiros/homerow/internal/domain/drill"
)

func typeEvery(s *drill.Session, text string, gap time.Duration) {
	at := start
	for _, r := range text {
		s.Type(r, at)
		at = at.Add(gap)
	}
}

func TestSpeedIsZeroUntilTimeHasPassed(t *testing.T) {
	s := drill.NewSession("fj")
	s.Type('f', start)

	if s.WPM() != 0 || s.Elapsed() != 0 {
		t.Fatalf("WPM() = %v, Elapsed() = %v", s.WPM(), s.Elapsed())
	}
}

func TestSpeedCountsFiveCorrectCharactersAsOneWord(t *testing.T) {
	s := drill.NewSession("fjfjf fjfjf ")
	typeEvery(s, "fjfjf fjfjf ", 5*time.Second)

	if s.Elapsed() != 55*time.Second {
		t.Fatalf("Elapsed() = %v, want 55s", s.Elapsed())
	}
	want := 12.0 / 5 / (55.0 / 60)
	if got := s.WPM(); got < want-0.001 || got > want+0.001 {
		t.Fatalf("WPM() = %v, want %v", got, want)
	}
}

func TestWrongCharactersDoNotAddSpeed(t *testing.T) {
	s := drill.NewSession("fffff fffff")
	typeEvery(s, "fffff jjjjj", 6*time.Second)

	want := 6.0 / 5 / 1.0
	if got := s.WPM(); got < want-0.001 || got > want+0.001 {
		t.Fatalf("WPM() = %v, want %v", got, want)
	}
}
