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
	if !strings.Contains(string(raw), `"version": 2`) || !strings.Contains(string(raw), `"2": {`) {
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

func TestUnknownLessonKeyIsCorrupt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	_ = os.WriteFile(path, []byte(`{"version":1,"lessons":{"one":{"completed":true}}}`), 0o600)

	if _, err := filestore.New(path).Load(); !errors.Is(err, filestore.ErrCorrupt) {
		t.Fatalf("err = %v, want ErrCorrupt", err)
	}
}

func TestUnreadableFileIsReportedAsIs(t *testing.T) {
	path := t.TempDir()

	_, err := filestore.New(path).Load()
	if err == nil || errors.Is(err, filestore.ErrCorrupt) {
		t.Fatalf("err = %v, want a read error", err)
	}
}

func TestSaveFailsWhenTheFolderCannotBeMade(t *testing.T) {
	blocker := filepath.Join(t.TempDir(), "file")
	_ = os.WriteFile(blocker, nil, 0o600)

	if err := filestore.New(filepath.Join(blocker, "progress.json")).Save(progress.NewBoard(nil)); err == nil {
		t.Fatal("want an error")
	}
}

func TestSaveFailsWhenTheTargetIsAFolder(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "progress.json")
	_ = os.MkdirAll(filepath.Join(target, "inside"), 0o700)

	if err := filestore.New(target).Save(progress.NewBoard(nil)); err == nil {
		t.Fatal("want an error")
	}
	files, _ := os.ReadDir(dir)
	if len(files) != 1 {
		t.Fatalf("files = %v", files)
	}
}

func TestSaveFailsInAReadOnlyFolder(t *testing.T) {
	dir := t.TempDir()
	_ = os.Chmod(dir, 0o500)
	t.Cleanup(func() { _ = os.Chmod(dir, 0o700) })
	if f, err := os.CreateTemp(dir, "probe"); err == nil {
		_ = f.Close()
		t.Skip("folder is still writable, running as root")
	}

	if err := filestore.New(filepath.Join(dir, "progress.json")).Save(progress.NewBoard(nil)); err == nil {
		t.Fatal("want an error")
	}
}

func TestMistakesAreStoredByKey(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	board := progress.NewBoard(nil)
	board.AddMistakes(map[rune]int{'f': 3, ';': 1})
	_ = filestore.New(path).Save(board)

	raw, _ := os.ReadFile(path)
	if !strings.Contains(string(raw), `"f": 3`) || !strings.Contains(string(raw), `";": 1`) {
		t.Fatalf("file = %s", raw)
	}
}

func TestVersionOneFilesStillLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	_ = os.WriteFile(path, []byte(`{"version":1,"lessons":{"4":{"completed":true,"stars":3,"wpm":11,"accuracy":96}}}`), 0o600)

	board, err := filestore.New(path).Load()
	if err != nil {
		t.Fatal(err)
	}
	if board.Entry(4).Stars != 3 || len(board.Mistakes()) != 0 {
		t.Fatalf("entry = %+v, mistakes = %v", board.Entry(4), board.Mistakes())
	}
}

func TestNewerFilesAndBadMistakeKeysAreCorrupt(t *testing.T) {
	for name, raw := range map[string]string{
		"newer version": `{"version":3,"lessons":{}}`,
		"long key":      `{"version":2,"lessons":{},"mistakes":{"fj":1}}`,
	} {
		path := filepath.Join(t.TempDir(), "progress.json")
		_ = os.WriteFile(path, []byte(raw), 0o600)

		if _, err := filestore.New(path).Load(); !errors.Is(err, filestore.ErrCorrupt) {
			t.Errorf("%s: err = %v, want ErrCorrupt", name, err)
		}
	}
}
