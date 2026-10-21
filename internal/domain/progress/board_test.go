package progress_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func TestUnplayedLessonHasAnEmptyEntry(t *testing.T) {
	b := progress.NewBoard(nil)

	if e := b.Entry(4); e.Completed || e.Stars != 0 {
		t.Fatalf("Entry(4) = %+v", e)
	}
}

func TestRecordKeepsTheBestOfEachAttempt(t *testing.T) {
	b := progress.NewBoard(nil)

	if !b.Record(4, progress.Entry{Completed: true, Stars: 3, WPM: 9, Accuracy: 95}) {
		t.Fatal("first attempt should improve the board")
	}
	if b.Record(4, progress.Entry{Completed: false, Stars: 0, WPM: 12, Accuracy: 70}) {
		t.Fatal("a failed attempt should not count as an improvement")
	}

	got := b.Entry(4)
	want := progress.Entry{Completed: true, Stars: 3, WPM: 12, Accuracy: 95}
	if got != want {
		t.Fatalf("Entry(4) = %+v, want %+v", got, want)
	}
}

func TestMoreStarsImprove(t *testing.T) {
	b := progress.NewBoard(map[int]progress.Entry{4: {Completed: true, Stars: 2}})

	if !b.Record(4, progress.Entry{Completed: true, Stars: 4}) || b.Entry(4).Stars != 4 {
		t.Fatalf("Entry(4) = %+v", b.Entry(4))
	}
}

func TestTotalsAcrossLessons(t *testing.T) {
	b := progress.NewBoard(map[int]progress.Entry{
		1: {Completed: true},
		2: {Completed: true},
		4: {Completed: true, Stars: 5},
		6: {Stars: 0},
	})

	if b.TotalStars() != 5 {
		t.Errorf("TotalStars() = %d, want 5", b.TotalStars())
	}
	if b.Percent(12) != 25 {
		t.Errorf("Percent(12) = %d, want 25", b.Percent(12))
	}
	if b.Percent(0) != 0 {
		t.Errorf("Percent(0) = %d, want 0", b.Percent(0))
	}
}

func TestEntriesAreACopy(t *testing.T) {
	b := progress.NewBoard(map[int]progress.Entry{1: {Completed: true}})

	b.Entries()[1] = progress.Entry{}

	if !b.Entry(1).Completed {
		t.Fatal("changing Entries() changed the board")
	}
}
