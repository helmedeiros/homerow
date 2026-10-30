package main

import (
	"bytes"
	"errors"
	"path/filepath"
	"testing"
)

func TestRunPrintsVersion(t *testing.T) {
	var out, errOut bytes.Buffer

	code := run([]string{"version"}, nil, &out, &errOut)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if got, want := out.String(), "homerow dev\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestRunRejectsUnknownCommandsAndFlags(t *testing.T) {
	for _, args := range [][]string{{"bogus"}, {"--nope"}} {
		var out, errOut bytes.Buffer
		if code := run(args, nil, &out, &errOut); code != 2 {
			t.Errorf("%v: exit code = %d, want 2", args, code)
		}
		if errOut.Len() == 0 {
			t.Errorf("%v: no usage printed", args)
		}
	}
}

func TestProgressPathPrefersTheFlagThenTheEnvironment(t *testing.T) {
	env := func(key string) string {
		if key == "HOMEROW_DATA" {
			return "/env/progress.json"
		}
		return ""
	}
	noEnv := func(string) string { return "" }
	configDir := func() (string, error) { return "/config", nil }

	cases := []struct {
		flag string
		env  func(string) string
		want string
	}{
		{"/flag/p.json", env, "/flag/p.json"},
		{"", env, "/env/progress.json"},
		{"", noEnv, filepath.Join("/config", "homerow", "progress.json")},
	}
	for _, c := range cases {
		got, err := progressPath(c.flag, c.env, configDir)
		if err != nil || got != c.want {
			t.Errorf("progressPath(%q) = %q, %v, want %q", c.flag, got, err, c.want)
		}
	}
}

func TestProgressPathReportsAMissingConfigDir(t *testing.T) {
	broken := func() (string, error) { return "", errors.New("no home") }

	if _, err := progressPath("", func(string) string { return "" }, broken); err == nil {
		t.Fatal("expected an error")
	}
}
