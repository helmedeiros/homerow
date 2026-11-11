package app_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func TestFinishingALessonRecordsItsMistakesPerKey(t *testing.T) {
	store := &memoryStore{}
	trainer := app.NewTrainer(fixedCatalog{sampleModule()}, store, fixedSeed)

	_, _ = trainer.Finish(3, scoring.Result{Accuracy: 90, WPM: 10, Mistakes: map[rune]int{'f': 2}})
	_, _ = trainer.Finish(3, scoring.Result{Accuracy: 95, WPM: 10, Mistakes: map[rune]int{'f': 1, 'j': 1}})

	if store.mistakes['f'] != 3 || store.mistakes['j'] != 1 {
		t.Fatalf("mistakes = %v", store.mistakes)
	}
}
