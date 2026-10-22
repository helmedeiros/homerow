package app_test

import (
	"errors"
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestStartPreparesTheLessonSteps(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{}, fixedSeed)

	plan, err := trainer.Start(2)
	if err != nil {
		t.Fatal(err)
	}

	if plan.Lesson.Title != "Keys f and j" || len(plan.Steps) != 4 {
		t.Fatalf("plan = %q with %d steps", plan.Lesson.Title, len(plan.Steps))
	}
	if plan.Steps[0].Kind != curriculum.KeyIntroStep || plan.Steps[0].Key != 'f' {
		t.Fatalf("first step = %+v", plan.Steps[0])
	}
}

func TestStartIsRepeatableForTheSameSeed(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{}, fixedSeed)

	a, _ := trainer.Start(3)
	b, _ := trainer.Start(3)

	if a.Steps[0].Text != b.Steps[0].Text {
		t.Fatalf("drills differ: %q vs %q", a.Steps[0].Text, b.Steps[0].Text)
	}
}

func TestStartGeneratesNewDrillsForNewSeeds(t *testing.T) {
	seed := uint64(0)
	next := func() uint64 { seed++; return seed }
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{}, next)

	a, _ := trainer.Start(3)
	b, _ := trainer.Start(3)

	if a.Steps[0].Text == b.Steps[0].Text {
		t.Fatal("expected a fresh drill on each start")
	}
}

func TestStartRejectsUnknownLessons(t *testing.T) {
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, &memoryStore{}, fixedSeed)

	if _, err := trainer.Start(99); !errors.Is(err, app.ErrUnknownLesson) {
		t.Fatalf("err = %v, want ErrUnknownLesson", err)
	}
}
