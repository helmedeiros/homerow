package app

import (
	"fmt"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type Outcome struct {
	Result   scoring.Result
	Stars    int
	Scored   bool
	Passed   bool
	Improved bool
	Next     int
}

func (t *Trainer) Finish(number int, result scoring.Result) (Outcome, error) {
	lesson, err := t.lesson(number)
	if err != nil {
		return Outcome{}, err
	}
	out := grade(lesson, result)
	out.Next = t.catalog.Course().Next(number)

	board, err := t.store.Load()
	if err != nil {
		return Outcome{}, fmt.Errorf("loading progress: %w", err)
	}
	out.Improved = board.Record(number, progress.Entry{
		Completed: out.Passed,
		Stars:     out.Stars,
		WPM:       result.WPM,
		Accuracy:  result.Accuracy,
	})
	board.AddMistakes(result.Mistakes)
	if err := t.store.Save(board); err != nil {
		return Outcome{}, fmt.Errorf("saving progress: %w", err)
	}
	return out, nil
}

func grade(lesson curriculum.Lesson, result scoring.Result) Outcome {
	if !lesson.Kind.Scored() {
		return Outcome{Result: result, Passed: true}
	}
	stars := scoring.Stars(result, lesson.GoalWPM)
	return Outcome{Result: result, Stars: stars, Scored: true, Passed: scoring.Passed(stars)}
}
