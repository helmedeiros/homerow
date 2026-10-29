package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

func passed() app.Outcome {
	return app.Outcome{
		Result: scoring.Result{Accuracy: 96.4, WPM: 13.2},
		Stars:  3, Scored: true, Passed: true, Improved: true, Next: 5,
	}
}

func TestResultsCelebrateAPass(t *testing.T) {
	view := newResultsScreen(4, passed()).View()

	for _, want := range []string{"GOOD JOB!", "★★★☆☆", "13 wpm", "96%", "new best", "next lesson"} {
		if !strings.Contains(view, want) {
			t.Errorf("missing %q:\n%s", want, view)
		}
	}
}

func TestResultsEncourageAfterAFail(t *testing.T) {
	out := app.Outcome{Result: scoring.Result{Accuracy: 70, WPM: 9}, Scored: true}
	view := newResultsScreen(4, out).View()

	if !strings.Contains(view, "KEEP PRACTISING") || !strings.Contains(view, "80%") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestUnscoredResultsShowNoStars(t *testing.T) {
	view := newResultsScreen(1, app.Outcome{Passed: true, Next: 2}).View()

	if strings.Contains(view, "★") || strings.Contains(view, "☆") || !strings.Contains(view, "GOOD JOB!") {
		t.Fatalf("view:\n%s", view)
	}
}

func TestEnterMovesToTheNextLessonAfterAPass(t *testing.T) {
	_, cmd := newResultsScreen(4, passed()).Update(special(tea.KeyEnter))

	if got := message(cmd); got != (startLessonMsg{number: 5}) {
		t.Fatalf("message = %#v, want start lesson 5", got)
	}
}

func TestEnterRetriesAfterAFail(t *testing.T) {
	_, cmd := newResultsScreen(4, app.Outcome{Scored: true}).Update(special(tea.KeyEnter))

	if got := message(cmd); got != (startLessonMsg{number: 4}) {
		t.Fatalf("message = %#v, want start lesson 4", got)
	}
}

func TestRetryAndLeave(t *testing.T) {
	_, cmd := newResultsScreen(4, passed()).Update(runes("r"))
	if got := message(cmd); got != (startLessonMsg{number: 4}) {
		t.Fatalf("message = %#v, want start lesson 4", got)
	}

	_, cmd = newResultsScreen(4, passed()).Update(special(tea.KeyEsc))
	if _, ok := message(cmd).(leaveLessonMsg); !ok {
		t.Fatalf("message = %#v, want leave", message(cmd))
	}
}

func TestTheLastLessonReturnsToTheMap(t *testing.T) {
	out := passed()
	out.Next = 0

	_, cmd := newResultsScreen(23, out).Update(special(tea.KeyEnter))

	if _, ok := message(cmd).(leaveLessonMsg); !ok {
		t.Fatalf("message = %#v, want leave", message(cmd))
	}
}
