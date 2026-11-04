package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const someProgress = `{"version": 1, "lessons": {"1": {"completed": true}, "4": {"completed": true, "stars": 3}}}`

func resetWith(t *testing.T, answer string) (int, string, string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "progress.json")
	if err := os.WriteFile(path, []byte(someProgress), 0o600); err != nil {
		t.Fatal(err)
	}
	var out, errOut bytes.Buffer
	code := run([]string{"--data", path, "reset"}, strings.NewReader(answer), &out, &errOut)
	var list bytes.Buffer
	run([]string{"--data", path, "lessons"}, nil, &list, &errOut)
	return code, out.String(), list.String()
}

func TestResetErasesProgressOnceConfirmed(t *testing.T) {
	for _, answer := range []string{"y\n", "yes\n", "Y"} {
		code, out, list := resetWith(t, answer)

		if code != 0 || !strings.Contains(out, "Erase all progress? [y/N] ") || !strings.Contains(out, "progress erased") {
			t.Errorf("%q: code = %d, out = %q", answer, code, out)
		}
		if !strings.Contains(list, "0% progress · 0 stars") {
			t.Errorf("%q: progress kept:\n%s", answer, list)
		}
	}
}

func TestResetKeepsProgressUnlessConfirmed(t *testing.T) {
	for _, answer := range []string{"\n", "n\n", "nope\n", ""} {
		code, out, list := resetWith(t, answer)

		if code != 0 || !strings.Contains(out, "nothing changed") {
			t.Errorf("%q: code = %d, out = %q", answer, code, out)
		}
		if !strings.Contains(list, "8% progress · 3 stars") {
			t.Errorf("%q: progress lost:\n%s", answer, list)
		}
	}
}
