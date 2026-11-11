package app

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

var ErrUnknownLesson = errors.New("unknown lesson")

type Plan struct {
	Lesson curriculum.Lesson
	Steps  []curriculum.Step
}

func (t *Trainer) Start(number int) (Plan, error) {
	lesson, err := t.lesson(number)
	if err != nil {
		return Plan{}, err
	}
	seed := t.seeds()
	rng := rand.New(rand.NewPCG(seed, seed))
	if lesson.Kind != curriculum.Dynamic {
		return Plan{Lesson: lesson, Steps: lesson.Steps(rng)}, nil
	}
	board, err := t.store.Load()
	if err != nil {
		return Plan{}, fmt.Errorf("loading progress: %w", err)
	}
	return Plan{Lesson: lesson, Steps: lesson.DynamicSteps(rng, board.Weakest(lesson.Keys, weakKeys))}, nil
}

const weakKeys = 3

func (t *Trainer) lesson(number int) (curriculum.Lesson, error) {
	lesson, _, ok := t.catalog.Course().Lesson(number)
	if !ok {
		return curriculum.Lesson{}, fmt.Errorf("%w: %d", ErrUnknownLesson, number)
	}
	return lesson, nil
}
