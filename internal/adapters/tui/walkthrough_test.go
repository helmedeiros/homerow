package tui

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/adapters/filestore"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
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

func (h *harness) playModule(m curriculum.Module) {
	for _, l := range m.Lessons {
		h.playLesson(l.Number)
	}
}

func moduleHeader(m curriculum.Module) string {
	scored := 0
	for _, l := range m.Lessons {
		if l.Kind.Scored() {
			scored++
		}
	}
	return fmt.Sprintf("Module %d · %s  100%% · %d stars", m.Number, m.Title, scored*scoring.MaxStars)
}

func (h *harness) showModule(m curriculum.Module) string {
	for range content.Course().Modules {
		if strings.Contains(h.view(), fmt.Sprintf("Module %d ·", m.Number)) {
			break
		}
		h.send(special(tea.KeyLeft))
	}
	return h.view()
}

func TestPlayingTheWholeHomeRowModule(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	home := content.HomeRow()

	newHarnessWith(t, filestore.New(path)).playModule(home)

	view := newHarnessWith(t, filestore.New(path)).showModule(home)
	if want := moduleHeader(home); !strings.Contains(view, want) {
		t.Fatalf("want %q after reopening:\n%s", want, view)
	}
}
