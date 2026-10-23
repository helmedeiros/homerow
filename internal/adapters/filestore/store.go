package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"

	"github.com/helmedeiros/homerow/internal/domain/progress"
)

const formatVersion = 1

var ErrCorrupt = errors.New("progress file is corrupt")

type Store struct {
	path string
}

func New(path string) *Store {
	return &Store{path: path}
}

type document struct {
	Version int               `json:"version"`
	Lessons map[string]record `json:"lessons"`
}

type record struct {
	Completed bool    `json:"completed"`
	Stars     int     `json:"stars"`
	WPM       float64 `json:"wpm"`
	Accuracy  float64 `json:"accuracy"`
}

func (s *Store) Load() (*progress.Board, error) {
	raw, err := os.ReadFile(s.path)
	if errors.Is(err, fs.ErrNotExist) {
		return progress.NewBoard(nil), nil
	}
	if err != nil {
		return nil, err
	}
	var doc document
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrCorrupt, s.path, err)
	}
	entries := map[int]progress.Entry{}
	for key, r := range doc.Lessons {
		number, err := strconv.Atoi(key)
		if err != nil {
			return nil, fmt.Errorf("%w: %s: lesson %q", ErrCorrupt, s.path, key)
		}
		entries[number] = progress.Entry(r)
	}
	return progress.NewBoard(entries), nil
}

func (s *Store) Save(b *progress.Board) error {
	doc := document{Version: formatVersion, Lessons: map[string]record{}}
	for number, e := range b.Entries() {
		doc.Lessons[strconv.Itoa(number)] = record(e)
	}
	raw, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	return writeAtomically(s.path, raw)
}

func writeAtomically(path string, raw []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".progress-*")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmp.Name()) }()
	if _, err := tmp.Write(raw); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}
