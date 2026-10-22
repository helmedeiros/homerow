package app_test

import (
	"errors"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

type fixedCatalog struct{ module curriculum.Module }

func (c fixedCatalog) Module() curriculum.Module { return c.module }

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
