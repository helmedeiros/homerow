package memstore

import (
	"sync"

	"github.com/helmedeiros/homerow/internal/domain/progress"
)

type Store struct {
	mu      sync.Mutex
	entries map[int]progress.Entry
}

func New() *Store {
	return &Store{}
}

func (s *Store) Load() (*progress.Board, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return progress.NewBoard(s.entries), nil
}

func (s *Store) Save(b *progress.Board) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries = b.Entries()
	return nil
}
