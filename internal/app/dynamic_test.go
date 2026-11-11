package app_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func dynamicModule() curriculum.Module {
	return curriculum.Module{Number: 1, Title: "Home Row", Lessons: []curriculum.Lesson{
		{Number: 1, Title: "Dynamic practice", Kind: curriculum.Dynamic, Keys: []rune("asdfjkl;"), GoalWPM: 10},
	}}
}

func TestDynamicPracticeDrillsTheKeysYouMissMost(t *testing.T) {
	store := &memoryStore{mistakes: map[rune]int{'k': 9, 'a': 4, 'q': 20}}
	trainer := app.NewTrainer(fixedCatalog{dynamicModule()}, store, fixedSeed)

	plan, err := trainer.Start(1)
	if err != nil {
		t.Fatal(err)
	}

	for _, g := range strings.Fields(plan.Steps[0].Text) {
		if !strings.ContainsAny(g, "ka") || strings.Contains(g, "q") {
			t.Fatalf("group %q should hold k or a and never q", g)
		}
	}
}

func TestDynamicPracticeNeedsTheProgressFile(t *testing.T) {
	store := &memoryStore{failing: errors.New("disk gone")}
	trainer := app.NewTrainer(fixedCatalog{dynamicModule()}, store, fixedSeed)

	if _, err := trainer.Start(1); err == nil {
		t.Fatal("want an error")
	}
}
