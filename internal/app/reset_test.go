package app_test

import (
	"errors"
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func TestResetClearsAllProgress(t *testing.T) {
	store := &memoryStore{entries: map[int]progress.Entry{1: {Completed: true}, 3: {Completed: true, Stars: 4}}}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	if err := trainer.Reset(); err != nil {
		t.Fatal(err)
	}

	overview, _ := trainer.Overview()
	if len(store.entries) != 0 || overview.TotalStars != 0 || overview.Percent != 0 {
		t.Fatalf("entries = %v, overview = %+v", store.entries, overview)
	}
}

func TestResetReportsStoreFailures(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{failing: errDisk}, fixedSeed)

	if err := trainer.Reset(); !errors.Is(err, errDisk) {
		t.Fatalf("err = %v", err)
	}
}
