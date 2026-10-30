package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
		"Home Row · 8% progress · 3 stars",
		"  1  Welcome to touch typing",
		"  4  Review: f and j",
		"★★★☆☆",
		" 23  Play: home row",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q:\n%s", want, text)
		}
	}
	if strings.Count(text, "\n") != 24 {
		t.Errorf("expected a header and 23 lessons:\n%s", text)
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
