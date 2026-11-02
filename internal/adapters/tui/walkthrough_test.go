package tui

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/adapters/filestore"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func (h *harness) finishStep() {
	h.t.Helper()
	switch s := h.current().(type) {
	case pageScreen:
		h.send(special(tea.KeyEnter))
	case keyIntroScreen:
		if s.key == ' ' {
			h.send(special(tea.KeySpace))
		} else {
			h.send(runes(string(s.key)))
		}
	case drillScreen:
		h.typeDrill()
	case gameScreen:
		h.playPerfectly()
	default:
		h.t.Fatalf("unexpected screen %T:\n%s", s, h.view())
	}
}

func (h *harness) playPerfectly() {
	for {
		g, ok := h.current().(gameScreen)
		if !ok {
			return
		}
		for _, d := range g.game.Drops() {
			h.send(runes(string(d.Char)))
		}
		h.clock.tick(frameEvery)
		h.send(frameMsg{})
	}
}

func (h *harness) playLesson(number int) {
	h.t.Helper()
	h.send(special(tea.KeyEnter))
	if !strings.Contains(h.view(), "Lesson "+strconv.Itoa(number)+" ") {
		h.t.Fatalf("expected lesson %d to open:\n%s", number, h.view())
	}
	for range 20 {
		if _, done := h.current().(resultsScreen); done {
			break
		}
		h.finishStep()
	}
	if !strings.Contains(h.view(), "GOOD JOB!") {
		h.t.Fatalf("lesson %d not passed:\n%s", number, h.view())
	}
	h.send(special(tea.KeyEsc))
}

func TestPlayingTheWholeHomeRowModule(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	h := newHarnessWith(t, filestore.New(path))

	for number := 1; number <= 23; number++ {
		h.playLesson(number)
	}

	scored := 0
	for _, l := range content.HomeRow().Lessons {
		if l.Kind.Scored() {
			scored++
		}
	}
	reopened := newHarnessWith(t, filestore.New(path))
	view := reopened.view()
	want := "100% progress · " + strconv.Itoa(scored*scoring.MaxStars) + " stars"
	if !strings.Contains(view, want) {
		t.Fatalf("want %q after reopening:\n%s", want, view)
	}
}
