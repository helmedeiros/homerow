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
	return Plan{Lesson: lesson, Steps: lesson.Steps(rand.New(rand.NewPCG(seed, seed)))}, nil
}

func (t *Trainer) lesson(number int) (curriculum.Lesson, error) {
	lesson, _, ok := t.catalog.Course().Lesson(number)
	if !ok {
		return curriculum.Lesson{}, fmt.Errorf("%w: %d", ErrUnknownLesson, number)
	}
	return lesson, nil
}
