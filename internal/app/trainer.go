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

type Section struct {
	Number  int
	Title   string
	First   int
	Last    int
	Stars   int
	Percent int
}

type Overview struct {
	Modules    []Section
	Lessons    []Card
	TotalStars int
	Percent    int
}

func (o Overview) ModuleOf(index int) Section {
	for _, s := range o.Modules {
		if index >= s.First && index <= s.Last {
			return s
		}
	}
	return Section{}
}

func (t *Trainer) Overview() (Overview, error) {
	board, err := t.store.Load()
	if err != nil {
		return Overview{}, fmt.Errorf("loading progress: %w", err)
	}
	var o Overview
	for _, m := range t.catalog.Course().Modules {
		section := Section{Number: m.Number, Title: m.Title, First: len(o.Lessons)}
		for _, l := range m.Lessons {
			o.Lessons = append(o.Lessons, Card{Number: l.Number, Title: l.Title, Kind: l.Kind, Keys: string(l.NewKeys), Entry: board.Entry(l.Number)})
		}
		section.Last = len(o.Lessons) - 1
		section.Stars, section.Percent = tally(o.Lessons[section.First:])
		o.Modules = append(o.Modules, section)
	}
	o.TotalStars, o.Percent = tally(o.Lessons)
	return o, nil
}

func tally(cards []Card) (stars, percent int) {
	if len(cards) == 0 {
		return 0, 0
	}
	done := 0
	for _, c := range cards {
		stars += c.Entry.Stars
		if c.Entry.Completed {
			done++
		}
	}
	return stars, done * 100 / len(cards)
}
