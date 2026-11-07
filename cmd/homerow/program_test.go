package main

import (
	"bytes"
	"path/filepath"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/adapters/filestore"
	"github.com/helmedeiros/homerow/internal/app"
)

func startProgram(t *testing.T, path string) *teatest.TestModel {
	t.Helper()
	trainer := app.NewTrainer(content.Catalog{}, filestore.New(path), func() uint64 { return 1 })
	model, err := newModel(trainer)
	if err != nil {
		t.Fatal(err)
	}
	return teatest.NewTestModel(t, model, teatest.WithInitialTermSize(100, 40))
}

func waitFor(t *testing.T, tm *teatest.TestModel, text string) {
	t.Helper()
	teatest.WaitFor(t, tm.Output(), func(out []byte) bool {
		return bytes.Contains(out, []byte(text))
	}, teatest.WithDuration(5*time.Second))
}

func TestTheProgramRunsAnIntroLessonAndQuits(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	tm := startProgram(t, path)

	waitFor(t, tm, "Welcome to touch typing")
	tm.Send(tea.KeyMsg{Type: tea.KeyEnter})
	waitFor(t, tm, "Lesson 1")
	for range 4 {
		tm.Send(tea.KeyMsg{Type: tea.KeyEnter})
	}
	waitFor(t, tm, "GOOD JOB!")
	tm.Send(tea.KeyMsg{Type: tea.KeyEsc})
	waitFor(t, tm, "Home Row  4% · 0 stars")
	tm.Type("q")
	tm.WaitFinished(t, teatest.WithFinalTimeout(5*time.Second))

	board, err := filestore.New(path).Load()
	if err != nil || !board.Entry(1).Completed {
		t.Fatalf("lesson 1 not saved: %+v, %v", board.Entry(1), err)
	}
}

func TestTheProgramIntroducesKeysAndStartsTheDrill(t *testing.T) {
	tm := startProgram(t, filepath.Join(t.TempDir(), "progress.json"))
	waitFor(t, tm, "Welcome to touch typing")

	tm.Type("j")
	tm.Send(tea.KeyMsg{Type: tea.KeyEnter})
	waitFor(t, tm, "left index finger")
	tm.Type("f")
	waitFor(t, tm, "right index finger")
	tm.Type("j")
	waitFor(t, tm, "f f f ff ff fff")

	tm.Send(tea.KeyMsg{Type: tea.KeyCtrlC})
	tm.WaitFinished(t, teatest.WithFinalTimeout(5*time.Second))
}
