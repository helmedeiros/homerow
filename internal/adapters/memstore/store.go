package memstore

import (
	"sync"

	"github.com/helmedeiros/homerow/internal/domain/progress"
)

type Store struct {
	mu       sync.Mutex
	entries  map[int]progress.Entry
	mistakes map[rune]int
}

func New() *Store {
	return &Store{}
}

func (s *Store) Load() (*progress.Board, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	board := progress.NewBoard(s.entries)
	board.AddMistakes(s.mistakes)
	return board, nil
}

func (s *Store) Save(b *progress.Board) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries = b.Entries()
	s.mistakes = b.Mistakes()
	return nil
}
