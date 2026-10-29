package tui

import (
	"strings"
	"testing"
	"time"
)

func TestTimedDrillWaitsForTheFirstKeyBeforeCounting(t *testing.T) {
	clock := newClock()
	d := newDrillScreen("fj fj fj", time.Minute, clock.Now)

	if d.Init() == nil {
		t.Fatal("a timed drill should start ticking")
	}
	clock.tick(5 * time.Minute)
	s, cmd := d.Update(tickMsg{})

	if _, done := message(cmd).(stepDoneMsg); done {
		t.Fatal("finished before the first key")
	}
	if !strings.Contains(s.View(), "1:00") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestTimedDrillShowsTimeLeft(t *testing.T) {
	clock := newClock()
	s, _ := typeText(newDrillScreen("fj fj fj", time.Minute, clock.Now), clock, "fj")

	clock.tick(15 * time.Second)
	s, _ = s.Update(tickMsg{})

	if !strings.Contains(s.View(), "0:44") {
		t.Fatalf("view:\n%s", s.View())
	}
}

func TestTimedDrillEndsWhenTimeRunsOut(t *testing.T) {
	clock := newClock()
	s, _ := typeText(newDrillScreen("fj fj fj", time.Minute, clock.Now), clock, "fj f")

	clock.tick(time.Minute)
	_, cmd := s.Update(tickMsg{})

	done, ok := message(cmd).(stepDoneMsg)
	if !ok || done.result.Accuracy != 100 {
		t.Fatalf("message = %#v", message(cmd))
	}
}

func TestUntimedDrillHasNoClock(t *testing.T) {
	d := newDrillScreen("fj", 0, newClock().Now)

	if d.Init() != nil || strings.Contains(d.View(), "time left") {
		t.Fatal("untimed drill should not tick or show a clock")
	}
}
