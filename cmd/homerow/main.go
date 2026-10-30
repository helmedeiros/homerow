package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/adapters/filestore"
	"github.com/helmedeiros/homerow/internal/adapters/tui"
	"github.com/helmedeiros/homerow/internal/app"
)

var version = "dev"

const usage = `usage: homerow [--data FILE] [command]

commands:
  (none)    open the typing tutor
  version   print the version

flags:
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

func run(args []string, in io.Reader, out, errOut io.Writer) int {
	flags := flag.NewFlagSet("homerow", flag.ContinueOnError)
	flags.SetOutput(errOut)
	flags.Usage = func() {
		_, _ = fmt.Fprint(errOut, usage)
		flags.PrintDefaults()
	}
	data := flags.String("data", "", "progress file (default $HOMEROW_DATA or the user config dir)")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	switch flags.Arg(0) {
	case "":
		return withTrainer(*data, errOut, func(t *app.Trainer) error { return play(t, in, out) })
	case "version":
		_, _ = fmt.Fprintf(out, "homerow %s\n", version)
		return 0
	default:
		flags.Usage()
		return 2
	}
}

func withTrainer(data string, errOut io.Writer, use func(*app.Trainer) error) int {
	path, err := progressPath(data, os.Getenv, os.UserConfigDir)
	if err == nil {
		err = use(app.NewTrainer(content.Catalog{}, filestore.New(path), rand.Uint64))
	}
	if err != nil {
		_, _ = fmt.Fprintf(errOut, "homerow: %v\n", err)
		return 1
	}
	return 0
}

func play(trainer *app.Trainer, in io.Reader, out io.Writer) error {
	model, err := tui.New(trainer, tui.Options{Now: time.Now, Seed: rand.Uint64})
	if err != nil {
		return err
	}
	_, err = tea.NewProgram(model, tea.WithAltScreen(), tea.WithInput(in), tea.WithOutput(out)).Run()
	return err
}

func progressPath(flagValue string, getenv func(string) string, configDir func() (string, error)) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}
	if env := getenv("HOMEROW_DATA"); env != "" {
		return env, nil
	}
	dir, err := configDir()
	if err != nil {
		return "", fmt.Errorf("finding a place for progress (use --data): %w", err)
	}
	return filepath.Join(dir, "homerow", "progress.json"), nil
}
