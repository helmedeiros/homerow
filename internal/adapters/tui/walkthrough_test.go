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
	case gameScreen, balloonScreen, popUpScreen:
		h.playPerfectly()
	default:
		h.t.Fatalf("unexpected screen %T:\n%s", s, h.view())
	}
}

func (h *harness) playPerfectly() {
	for {
		switch s := h.current().(type) {
		case gameScreen:
			for _, d := range s.game.Drops() {
				h.send(runes(string(d.Char)))
			}
		case popUpScreen:
			for _, p := range s.game.Shown() {
				h.send(runes(string(p.Char)))
			}
		case balloonScreen:
			for _, b := range s.game.Balloons() {
				h.send(runes(b.Word[b.Typed:]))
			}
		default:
			return
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
		h.send(special(tea.KeyLeft))
	}
	for range m.Number - 1 {
		h.send(special(tea.KeyRight))
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

func TestPlayingTheWholeTopRowModule(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	top := content.TopRow()

	h := newHarnessWith(t, filestore.New(path))
	h.send(special(tea.KeyRight))
	h.playModule(top)

	reopened := newHarnessWith(t, filestore.New(path))
	if view, want := reopened.showModule(top), moduleHeader(top); !strings.Contains(view, want) {
		t.Fatalf("want %q after reopening:\n%s", want, view)
	}
	if view := reopened.showModule(content.HomeRow()); !strings.Contains(view, "Module 1 · Home Row  0% · 0 stars") {
		t.Fatalf("home row should be untouched:\n%s", view)
	}
}

func TestPlayingTheWholeBottomRowModule(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	bottom := content.BottomRow()

	h := newHarnessWith(t, filestore.New(path))
	h.send(special(tea.KeyRight))
	h.send(special(tea.KeyRight))
	h.playModule(bottom)

	reopened := newHarnessWith(t, filestore.New(path))
	if view, want := reopened.showModule(bottom), moduleHeader(bottom); !strings.Contains(view, want) {
		t.Fatalf("want %q after reopening:\n%s", want, view)
	}
	if view := reopened.showModule(content.TopRow()); !strings.Contains(view, "Module 2 · Top Row  0% · 0 stars") {
		t.Fatalf("top row should be untouched:\n%s", view)
	}
}
