package filestore_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/filestore"
	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/app/storetest"
	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func TestFileStoreHonoursTheContract(t *testing.T) {
	storetest.Contract(t, func(t *testing.T) app.ProgressStore {
		return filestore.New(filepath.Join(t.TempDir(), "nested", "progress.json"))
	})
}

func TestProgressSurvivesANewStoreOnTheSameFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	_ = filestore.New(path).Save(progress.NewBoard(map[int]progress.Entry{7: {Completed: true, Stars: 3}}))

	board, err := filestore.New(path).Load()
	if err != nil {
		t.Fatal(err)
	}
	if board.Entry(7).Stars != 3 {
		t.Fatalf("entry 7 = %+v", board.Entry(7))
	}
}

func TestFileIsReadableJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	_ = filestore.New(path).Save(progress.NewBoard(map[int]progress.Entry{2: {Completed: true}}))

	raw, _ := os.ReadFile(path)
	if !strings.Contains(string(raw), `"version": 1`) || !strings.Contains(string(raw), `"2": {`) {
		t.Fatalf("file = %s", raw)
	}
}

func TestCorruptFileIsReported(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	_ = os.WriteFile(path, []byte("{not json"), 0o600)

	if _, err := filestore.New(path).Load(); !errors.Is(err, filestore.ErrCorrupt) {
		t.Fatalf("err = %v, want ErrCorrupt", err)
	}
}

func TestSaveLeavesNoTemporaryFiles(t *testing.T) {
	dir := t.TempDir()
	_ = filestore.New(filepath.Join(dir, "progress.json")).Save(progress.NewBoard(nil))

	files, _ := os.ReadDir(dir)
	if len(files) != 1 {
		t.Fatalf("files = %v", files)
	}
}
