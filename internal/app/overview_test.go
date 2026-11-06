package app_test

import (
	"errors"
	"slices"
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

	if len(o.Modules) != 1 || o.Modules[0].Title != "Home Row" || len(o.Lessons) != 4 {
		t.Fatalf("overview = %+v with %d lessons", o.Modules, len(o.Lessons))
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

func TestOverviewSplitsTheCourseIntoModules(t *testing.T) {
	store := &memoryStore{entries: map[int]progress.Entry{
		1: {Completed: true},
		5: {Completed: true},
		6: {Completed: true, Stars: 3},
	}}
	trainer := app.NewTrainer(courseCatalog{twoModules()}, store, fixedSeed)

	o, _ := trainer.Overview()

	if len(o.Lessons) != 6 {
		t.Fatalf("%d lessons", len(o.Lessons))
	}
	want := []app.Section{
		{Number: 1, Title: "Home Row", First: 0, Last: 3, Stars: 0, Percent: 25},
		{Number: 2, Title: "Top Row", First: 4, Last: 5, Stars: 3, Percent: 100},
	}
	if !slices.Equal(o.Modules, want) {
		t.Fatalf("modules = %+v", o.Modules)
	}
	if o.Percent != 50 || o.TotalStars != 3 {
		t.Fatalf("totals = %d%%, %d stars", o.Percent, o.TotalStars)
	}
	if o.ModuleOf(4).Title != "Top Row" || o.ModuleOf(0).Title != "Home Row" {
		t.Fatal("ModuleOf picked the wrong module")
	}
}

func TestOverviewIgnoresProgressForLessonsOutsideTheCourse(t *testing.T) {
	store := &memoryStore{entries: map[int]progress.Entry{1: {Completed: true, Stars: 2}, 99: {Completed: true, Stars: 5}}}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	o, _ := trainer.Overview()

	if o.Percent != 25 || o.TotalStars != 2 {
		t.Fatalf("totals = %d%%, %d stars", o.Percent, o.TotalStars)
	}
}

func TestOverviewOfAnEmptyCourse(t *testing.T) {
	o, err := app.NewTrainer(courseCatalog{}, &memoryStore{}, fixedSeed).Overview()

	if err != nil || len(o.Lessons) != 0 || o.Percent != 0 || o.ModuleOf(0).Title != "" {
		t.Fatalf("overview = %+v, err = %v", o, err)
	}
}
