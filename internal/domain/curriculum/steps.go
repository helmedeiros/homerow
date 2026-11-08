package curriculum

import (
	"math/rand/v2"
	"strings"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

type StepKind int

const (
	PageStep StepKind = iota
	KeyIntroStep
	DrillStep
	GameStep
)

const (
	warmUpGroups   = 10
	reviewGroups   = 16
	practiceGroups = 22
	balloonWords   = 40
)

type Step struct {
	Kind  StepKind
	Page  string
	Key   rune
	Text  string
	Keys  []rune
	Game  Game
	Words []string
}

func (l Lesson) Steps(rng *rand.Rand) []Step {
	switch l.Kind {
	case Intro:
		return l.pageSteps()
	case NewKeys:
		return l.newKeySteps(rng)
	case Practice:
		return []Step{drill(l.practiceText(rng))}
	case Play:
		return []Step{l.gameStep(rng)}
	default:
		return []Step{drill(exercise.Groups(rng, l.mix(), reviewGroups))}
	}
}

func (l Lesson) gameStep(rng *rand.Rand) Step {
	step := Step{Kind: GameStep, Keys: l.letters()}
	if l.Game != BalloonWords {
		return step
	}
	text, err := exercise.Words(rng, l.letters(), balloonWords)
	if err != nil {
		return step
	}
	step.Game = BalloonWords
	step.Words = strings.Fields(text)
	return step
}

func (l Lesson) pageSteps() []Step {
	steps := make([]Step, len(l.Pages))
	for i, p := range l.Pages {
		steps[i] = Step{Kind: PageStep, Page: p}
	}
	return steps
}

func (l Lesson) newKeySteps(rng *rand.Rand) []Step {
	var steps []Step
	for _, k := range l.NewKeys {
		steps = append(steps, Step{Kind: KeyIntroStep, Key: k})
	}
	newLetters := strings.ReplaceAll(string(l.NewKeys), " ", "")
	if newLetters != "" {
		steps = append(steps, drill(exercise.Introduce([]rune(newLetters))))
	}
	warmUp := exercise.Mix{Keys: l.letters(), Focus: []rune(newLetters)}
	return append(steps, drill(exercise.Groups(rng, warmUp, warmUpGroups)))
}

func (l Lesson) practiceText(rng *rand.Rand) string {
	if l.Words {
		if text, err := exercise.Words(rng, l.letters(), practiceGroups); err == nil {
			return text
		}
	}
	return exercise.Groups(rng, l.mix(), practiceGroups)
}

func (l Lesson) mix() exercise.Mix {
	return exercise.Mix{Keys: l.letters(), Focus: l.Focus}
}

func (l Lesson) letters() []rune {
	return []rune(strings.ReplaceAll(string(l.Keys), " ", ""))
}

func drill(text string) Step {
	return Step{Kind: DrillStep, Text: text}
}
