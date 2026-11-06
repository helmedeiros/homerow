package app_test

import (
	"errors"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

type fixedCatalog struct{ module curriculum.Module }

func (c fixedCatalog) Course() curriculum.Course {
	return curriculum.Course{Modules: []curriculum.Module{c.module}}
}

type courseCatalog struct{ course curriculum.Course }

func (c courseCatalog) Course() curriculum.Course { return c.course }

func twoModules() curriculum.Course {
	home := sampleModule()
	top := curriculum.Module{Number: 2, Title: "Top Row", Lessons: []curriculum.Lesson{
		{Number: 5, Title: "Keys r and u", Kind: curriculum.NewKeys, NewKeys: []rune("ru"), Keys: []rune("fjru")},
		{Number: 6, Title: "Review: r and u", Kind: curriculum.Review, Keys: []rune("fjru"), GoalWPM: 10},
	}}
	return curriculum.Course{Modules: []curriculum.Module{home, top}}
}

type memoryStore struct {
	entries map[int]progress.Entry
	saves   int
	failing error
}

func (s *memoryStore) Load() (*progress.Board, error) {
	if s.failing != nil {
		return nil, s.failing
	}
	return progress.NewBoard(s.entries), nil
}

func (s *memoryStore) Save(b *progress.Board) error {
	if s.failing != nil {
		return s.failing
	}
	s.saves++
	s.entries = b.Entries()
	return nil
}

var errDisk = errors.New("disk on fire")

func sampleModule() curriculum.Module {
	return curriculum.Module{Number: 1, Title: "Home Row", Lessons: []curriculum.Lesson{
		{Number: 1, Title: "Welcome", Kind: curriculum.Intro, Pages: []string{"hello"}},
		{Number: 2, Title: "Keys f and j", Kind: curriculum.NewKeys, NewKeys: []rune("fj"), Keys: []rune("fj")},
		{Number: 3, Title: "Review: f and j", Kind: curriculum.Review, Keys: []rune("fj"), GoalWPM: 10},
		{Number: 4, Title: "Play: f j", Kind: curriculum.Play, Keys: []rune("fj"), GoalWPM: 10},
	}}
}

func fixedSeed() uint64 { return 42 }
