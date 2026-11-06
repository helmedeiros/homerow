package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/memstore"
	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func TestLessonsListsTheModuleWithSavedProgress(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	saved := `{"version": 1, "lessons": {"1": {"completed": true}, "4": {"completed": true, "stars": 3}}}`
	if err := os.WriteFile(path, []byte(saved), 0o600); err != nil {
		t.Fatal(err)
	}
	var out, errOut bytes.Buffer

	code := run([]string{"--data", path, "lessons"}, nil, &out, &errOut)

	if code != 0 {
		t.Fatalf("exit code = %d: %s", code, errOut.String())
	}
	text := out.String()
	for _, want := range []string{
		"homerow · 8% progress · 3 stars",
		"Module 1 · Home Row · 8% · 3 stars",
		"  1  Welcome to touch typing",
		"  4  Review: f and j",
		"★★★☆☆",
		" 23  Play: home row",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q:\n%s", want, text)
		}
	}
	if strings.Count(text, "\n") != 26 {
		t.Errorf("expected two headers, a gap and 23 lessons:\n%s", text)
	}
}

func TestLessonsReportsACorruptProgressFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	_ = os.WriteFile(path, []byte("{"), 0o600)
	var out, errOut bytes.Buffer

	if code := run([]string{"--data", path, "lessons"}, nil, &out, &errOut); code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if !strings.Contains(errOut.String(), "corrupt") {
		t.Fatalf("stderr = %q", errOut.String())
	}
}

func TestLessonsSeparatesModules(t *testing.T) {
	course := curriculum.Course{Modules: []curriculum.Module{
		{Number: 1, Title: "Home Row", Lessons: []curriculum.Lesson{{Number: 1, Title: "Welcome", Kind: curriculum.Intro, Pages: []string{"hi"}}}},
		{Number: 2, Title: "Top Row", Lessons: []curriculum.Lesson{{Number: 2, Title: "Sit straight", Kind: curriculum.Intro, Pages: []string{"hi"}}}},
	}}
	trainer := app.NewTrainer(fixedCourse{course}, memstore.New(), func() uint64 { return 1 })
	var out bytes.Buffer

	if err := listLessons(trainer, &out); err != nil {
		t.Fatal(err)
	}

	want := "homerow · 0% progress · 0 stars\n\n" +
		"Module 1 · Home Row · 0% · 0 stars\n  1  Welcome                    ·\n\n" +
		"Module 2 · Top Row · 0% · 0 stars\n  2  Sit straight               ·\n"
	if out.String() != want {
		t.Fatalf("got:\n%s\nwant:\n%s", out.String(), want)
	}
}

type fixedCourse struct{ course curriculum.Course }

func (c fixedCourse) Course() curriculum.Course { return c.course }
