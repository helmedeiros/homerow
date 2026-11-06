package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func listLessons(trainer *app.Trainer, out io.Writer) error {
	o, err := trainer.Overview()
	if err != nil {
		return err
	}
	for _, m := range o.Modules {
		_, _ = fmt.Fprintf(out, "%s · %d%% progress · %d stars\n", m.Title, m.Percent, m.Stars)
		for _, c := range o.Lessons[m.First : m.Last+1] {
			_, _ = fmt.Fprintf(out, "%3d  %-26s %s\n", c.Number, c.Title, plainStatus(c))
		}
	}
	return nil
}

func plainStatus(c app.Card) string {
	switch {
	case c.Kind.Scored() && (c.Entry.Completed || c.Entry.Stars > 0):
		return strings.Repeat("★", c.Entry.Stars) + strings.Repeat("☆", scoring.MaxStars-c.Entry.Stars)
	case c.Entry.Completed:
		return "✓"
	default:
		return "·"
	}
}
