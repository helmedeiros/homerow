package app

import (
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

type Catalog interface {
	Course() curriculum.Course
}

type ProgressStore interface {
	Load() (*progress.Board, error)
	Save(*progress.Board) error
}

type SeedSource func() uint64
