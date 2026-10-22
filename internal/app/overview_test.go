package app_test

import (
	"errors"
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func TestOverviewListsLessonsWithSavedProgress(t *testing.T) {
	store := &memoryStore{entries: map[int]progress.Entry{
		1: {Completed: true},
		3: {Completed: true, Stars: 4},
	}}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	o, err := trainer.Overview()
	if err != nil {
		t.Fatal(err)
	}

	if o.Module != "Home Row" || len(o.Lessons) != 4 {
		t.Fatalf("overview = %q with %d lessons", o.Module, len(o.Lessons))
	}
	card := o.Lessons[2]
	if card.Number != 3 || card.Kind != curriculum.Review || card.Entry.Stars != 4 {
		t.Fatalf("card 3 = %+v", card)
	}
	if o.Lessons[1].Keys != "fj" {
		t.Fatalf("card 2 keys = %q", o.Lessons[1].Keys)
	}
	if o.TotalStars != 4 || o.Percent != 50 {
		t.Fatalf("totals = %d stars, %d%%", o.TotalStars, o.Percent)
	}
}

func TestOverviewReportsStoreFailures(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{failing: errDisk}, fixedSeed)

	if _, err := trainer.Overview(); !errors.Is(err, errDisk) {
		t.Fatalf("err = %v, want %v", err, errDisk)
	}
}
