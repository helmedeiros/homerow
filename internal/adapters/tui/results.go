package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/app"
)

type resultsScreen struct {
	lesson  int
	outcome app.Outcome
}

func newResultsScreen(lesson int, outcome app.Outcome) resultsScreen {
	return resultsScreen{lesson: lesson, outcome: outcome}
}

func (r resultsScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return r, nil
	}
	switch key.String() {
	case "enter", "right":
		return r, send(r.forward())
	case "r":
		return r, send(startLessonMsg{number: r.lesson})
	case "esc", "q":
		return r, send(leaveLessonMsg{})
	case "ctrl+c":
		return r, tea.Quit
	}
	return r, nil
}

func (r resultsScreen) forward() tea.Msg {
	switch {
	case !r.outcome.Passed:
		return startLessonMsg{number: r.lesson}
	case r.outcome.Next == 0:
		return leaveLessonMsg{}
	default:
		return startLessonMsg{number: r.outcome.Next}
	}
}

func (r resultsScreen) View() string {
	var b strings.Builder
	if r.outcome.Passed {
		b.WriteString(correctStyle.Bold(true).Render("✔ GOOD JOB!") + "\n\n")
	} else {
		b.WriteString(wrongTextStyle.Bold(true).Render("KEEP PRACTISING") + "\n\n")
	}
	if r.outcome.Scored {
		b.WriteString(r.scoreLines())
	}
	b.WriteString(mutedStyle.Render(r.help()))
	return b.String()
}

func (r resultsScreen) scoreLines() string {
	var b strings.Builder
	b.WriteString(renderStars(r.outcome.Stars) + "\n\n")
	fmt.Fprintf(&b, "speed %.0f wpm   accuracy %.0f%%\n", r.outcome.Result.WPM, r.outcome.Result.Accuracy)
	if r.outcome.Improved && r.outcome.Passed {
		b.WriteString(starStyle.Render("new best!") + "\n")
	}
	if !r.outcome.Passed {
		b.WriteString("You need 80% accuracy to pass. Slow down and aim for every key.\n")
	}
	return b.String() + "\n"
}

func (r resultsScreen) help() string {
	switch {
	case !r.outcome.Passed:
		return "enter try again · esc lessons"
	case r.outcome.Next == 0:
		return "enter lessons · r try again"
	default:
		return "enter next lesson · r try again · esc lessons"
	}
}
