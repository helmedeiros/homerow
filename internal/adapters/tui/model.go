package tui

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/game"
	"github.com/helmedeiros/homerow/internal/domain/scoring"
)

type Trainer interface {
	Overview() (app.Overview, error)
	Start(number int) (app.Plan, error)
	Finish(number int, result scoring.Result) (app.Outcome, error)
}

type Options struct {
	Now  func() time.Time
	Seed func() uint64
	tick ticker
}

func (o Options) ticker() ticker {
	if o.tick == nil {
		return realTicker
	}
	return o.tick
}

type Model struct {
	trainer Trainer
	opts    Options
	current screen
	run     *lessonRun
	err     error
}

type lessonRun struct {
	plan    app.Plan
	step    int
	results []scoring.Result
}

func New(trainer Trainer, opts Options) (Model, error) {
	m := Model{trainer: trainer, opts: opts}
	overview, err := trainer.Overview()
	if err != nil {
		return Model{}, err
	}
	m.current = newMapScreen(overview)
	return m, nil
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case startLessonMsg:
		return m.start(msg.number)
	case stepDoneMsg:
		return m.stepDone(msg)
	case leaveLessonMsg:
		return m.showMap()
	}
	next, cmd := m.current.Update(msg)
	m.current = next
	return m, cmd
}

func (m Model) start(number int) (tea.Model, tea.Cmd) {
	plan, err := m.trainer.Start(number)
	if err != nil {
		m.err = err
		return m, nil
	}
	m.err = nil
	m.run = &lessonRun{plan: plan}
	return m.enterStep()
}

func (m Model) enterStep() (tea.Model, tea.Cmd) {
	m.current = m.screenFor(m.run.plan.Lesson, m.run.plan.Steps[m.run.step])
	if s, ok := m.current.(interface{ Init() tea.Cmd }); ok {
		return m, s.Init()
	}
	return m, nil
}

func (m Model) screenFor(lesson curriculum.Lesson, step curriculum.Step) screen {
	switch step.Kind {
	case curriculum.PageStep:
		return newPageScreen(step.Page)
	case curriculum.KeyIntroStep:
		return newKeyIntroScreen(step.Key)
	case curriculum.GameStep:
		return m.gameFor(step)
	default:
		d := newDrillScreen(step.Text, lesson.TimeLimit, m.opts.Now)
		d.tick = m.opts.ticker()
		return d
	}
}

func (m Model) gameFor(step curriculum.Step) screen {
	seed := m.opts.Seed()
	rng := rand.New(rand.NewPCG(seed, seed))
	if step.Game == curriculum.BalloonWords {
		b := newBalloonScreen(game.DefaultBalloons(step.Words), rng, m.opts.Now)
		b.tick = m.opts.ticker()
		return b
	}
	g := newGameScreen(game.DefaultConfig(step.Keys), rng, m.opts.Now)
	g.tick = m.opts.ticker()
	return g
}

func (m Model) stepDone(msg stepDoneMsg) (tea.Model, tea.Cmd) {
	if m.run == nil {
		return m, nil
	}
	if msg.scored {
		m.run.results = append(m.run.results, msg.result)
	}
	m.run.step++
	if m.run.step < len(m.run.plan.Steps) {
		return m.enterStep()
	}
	number := m.run.plan.Lesson.Number
	outcome, err := m.trainer.Finish(number, average(m.run.results))
	if err != nil {
		m.err = err
		return m.showMap()
	}
	m.current = newResultsScreen(number, outcome)
	return m, nil
}

func average(results []scoring.Result) scoring.Result {
	if len(results) == 0 {
		return scoring.Result{}
	}
	var sum scoring.Result
	for _, r := range results {
		sum.Accuracy += r.Accuracy
		sum.WPM += r.WPM
	}
	n := float64(len(results))
	return scoring.Result{Accuracy: sum.Accuracy / n, WPM: sum.WPM / n}
}

func (m Model) showMap() (tea.Model, tea.Cmd) {
	focus := m.mapFocus()
	m.run = nil
	overview, err := m.trainer.Overview()
	if err != nil {
		m.err = err
		return m, nil
	}
	m.current = newMapScreen(overview).focus(focus)
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	if m.run != nil {
		b.WriteString(m.header() + "\n\n")
	}
	b.WriteString(m.current.View())
	if m.err != nil {
		b.WriteString("\n\n" + wrongTextStyle.Render(m.err.Error()))
	}
	return b.String() + "\n"
}

func (m Model) header() string {
	lesson := m.run.plan.Lesson
	total := len(m.run.plan.Steps)
	done := min(m.run.step, total)
	dots := accentStyle.Render(strings.Repeat("●", done)) + mutedStyle.Render(strings.Repeat("○", total-done))
	return fmt.Sprintf("%s  %s", titleStyle.Render(fmt.Sprintf("Lesson %d · %s", lesson.Number, lesson.Title)), dots)
}

func (m Model) mapFocus() int {
	if r, ok := m.current.(resultsScreen); ok && r.outcome.Passed && r.outcome.Next != 0 {
		return r.outcome.Next
	}
	if m.run != nil {
		return m.run.plan.Lesson.Number
	}
	return 0
}
