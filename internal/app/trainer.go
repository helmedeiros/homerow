package app

import (
	"fmt"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

type Trainer struct {
	catalog Catalog
	store   ProgressStore
	seeds   SeedSource
}

func NewTrainer(catalog Catalog, store ProgressStore, seeds SeedSource) *Trainer {
	return &Trainer{catalog: catalog, store: store, seeds: seeds}
}

type Card struct {
	Number int
	Title  string
	Kind   curriculum.Kind
	Keys   string
	Entry  progress.Entry
}

type Overview struct {
	Module     string
	Lessons    []Card
	TotalStars int
	Percent    int
}

func (t *Trainer) Overview() (Overview, error) {
	board, err := t.store.Load()
	if err != nil {
		return Overview{}, fmt.Errorf("loading progress: %w", err)
	}
	module := t.catalog.Module()
	cards := make([]Card, len(module.Lessons))
	for i, l := range module.Lessons {
		cards[i] = Card{Number: l.Number, Title: l.Title, Kind: l.Kind, Keys: string(l.NewKeys), Entry: board.Entry(l.Number)}
	}
	return Overview{
		Module:     module.Title,
		Lessons:    cards,
		TotalStars: board.TotalStars(),
		Percent:    board.Percent(len(module.Lessons)),
	}, nil
}
