package curriculum

import (
	"math/rand/v2"
	"strings"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
	"github.com/helmedeiros/homerow/internal/domain/keyboard"
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
		return l.playSteps(rng)
	case Dynamic:
		return l.DynamicSteps(rng, nil)
	case Tab:
		return []Step{drill(l.tabText(rng))}
	case Travel:
		return []Step{drill(exercise.Travel(rng, l.Focus, l.letters(), reviewGroups))}
	default:
		return []Step{drill(l.reviewText(rng))}
	}
}

func (l Lesson) playSteps(rng *rand.Rand) []Step {
	if l.Game == PopUpWarmUp {
		return []Step{{Kind: GameStep, Game: PopUpWarmUp, Keys: l.letters()}, {Kind: GameStep, Keys: l.letters()}}
	}
	return []Step{l.gameStep(rng)}
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

func (l Lesson) reviewText(rng *rand.Rand) string {
	if capitals := l.capitals(); len(capitals) > 0 {
		return exercise.CapitalGroups(rng, l.letters(), capitals, reviewGroups)
	}
	return exercise.Groups(rng, l.mix(), reviewGroups)
}

func (l Lesson) practiceText(rng *rand.Rand) string {
	if capitals := l.capitals(); len(capitals) > 0 && l.Words {
		return exercise.CapitalWords(rng, l.letters(), capitals, practiceGroups)
	}
	if l.Words {
		if text, err := exercise.FocusedWords(rng, l.letters(), l.Focus, practiceGroups); err == nil {
			return text
		}
	}
	return exercise.Groups(rng, l.mix(), practiceGroups)
}

func (l Lesson) DynamicSteps(rng *rand.Rand, weak []rune) []Step {
	return []Step{drill(exercise.Dynamic(rng, l.letters(), weak, practiceGroups))}
}

func (l Lesson) tabText(rng *rand.Rand) string {
	if text, err := exercise.Tabbed(rng, l.letters(), reviewGroups); err == nil {
		return text
	}
	return exercise.Groups(rng, l.mix(), reviewGroups)
}

func (l Lesson) mix() exercise.Mix {
	return exercise.Mix{Keys: l.letters(), Focus: l.Focus}
}

func (l Lesson) capitals() []rune {
	var capitals []rune
	for _, r := range l.Focus {
		if key, ok := keyboard.Lookup(r); ok && key.Shifted {
			capitals = append(capitals, r)
		}
	}
	return capitals
}

func (l Lesson) letters() []rune {
	return []rune(strings.ReplaceAll(string(l.Keys), " ", ""))
}

func drill(text string) Step {
	return Step{Kind: DrillStep, Text: text}
}
