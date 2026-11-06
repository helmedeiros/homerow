package app_test

import (
	"errors"
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func TestFinishingAScoredLessonEarnsStars(t *testing.T) {
	store := &memoryStore{}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	out, err := trainer.Finish(3, scoring.Result{Accuracy: 98, WPM: 12})
	if err != nil {
		t.Fatal(err)
	}

	if out.Stars != 5 || !out.Passed || !out.Scored || !out.Improved || out.Next != 4 {
		t.Fatalf("outcome = %+v", out)
	}
	if e := store.entries[3]; !e.Completed || e.Stars != 5 || e.WPM != 12 {
		t.Fatalf("saved = %+v", e)
	}
}

func TestFailingAScoredLessonDoesNotCompleteIt(t *testing.T) {
	store := &memoryStore{}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	out, _ := trainer.Finish(3, scoring.Result{Accuracy: 60, WPM: 30})

	if out.Passed || out.Stars != 0 || store.entries[3].Completed {
		t.Fatalf("outcome = %+v, saved = %+v", out, store.entries[3])
	}
}

func TestUnscoredLessonsCompleteWithoutStars(t *testing.T) {
	store := &memoryStore{}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	out, _ := trainer.Finish(1, scoring.Result{})

	if !out.Passed || out.Scored || out.Stars != 0 || !store.entries[1].Completed {
		t.Fatalf("outcome = %+v, saved = %+v", out, store.entries[1])
	}
}

func TestWorseAttemptsKeepTheBestStars(t *testing.T) {
	store := &memoryStore{}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)
	_, _ = trainer.Finish(3, scoring.Result{Accuracy: 95, WPM: 5})

	out, _ := trainer.Finish(3, scoring.Result{Accuracy: 85, WPM: 5})

	if out.Stars != 1 || out.Improved || store.entries[3].Stars != 3 {
		t.Fatalf("outcome = %+v, saved = %+v", out, store.entries[3])
	}
}

func TestTheLastLessonHasNoNext(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{}, fixedSeed)

	out, _ := trainer.Finish(4, scoring.Result{Accuracy: 100, WPM: 20})

	if out.Next != 0 {
		t.Fatalf("Next = %d, want 0", out.Next)
	}
}

func TestFinishReportsErrors(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{}, fixedSeed)
	if _, err := trainer.Finish(99, scoring.Result{}); !errors.Is(err, app.ErrUnknownLesson) {
		t.Fatalf("err = %v, want ErrUnknownLesson", err)
	}

	broken := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{failing: errDisk}, fixedSeed)
	if _, err := broken.Finish(1, scoring.Result{}); !errors.Is(err, errDisk) {
		t.Fatalf("err = %v, want %v", err, errDisk)
	}
}

func TestNextLessonCrossesIntoTheNextModule(t *testing.T) {
	trainer := app.NewTrainer(courseCatalog{twoModules()}, &memoryStore{}, fixedSeed)

	out, err := trainer.Finish(4, scoring.Result{Accuracy: 98, WPM: 12})
	if err != nil || out.Next != 5 {
		t.Fatalf("out = %+v, err = %v", out, err)
	}
	if _, err := trainer.Start(6); err != nil {
		t.Fatalf("Start(6) = %v", err)
	}
}
