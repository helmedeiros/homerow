package curriculum_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func rng() *rand.Rand { return rand.New(rand.NewPCG(1, 1)) }

func kinds(steps []curriculum.Step) []curriculum.StepKind {
	out := make([]curriculum.StepKind, len(steps))
	for i, s := range steps {
		out[i] = s.Kind
	}
	return out
}

func sameKinds(t *testing.T, got []curriculum.Step, want ...curriculum.StepKind) {
	t.Helper()
	g := kinds(got)
	if len(g) != len(want) {
		t.Fatalf("step kinds = %v, want %v", g, want)
	}
	for i := range want {
		if g[i] != want[i] {
			t.Fatalf("step kinds = %v, want %v", g, want)
		}
	}
}

func TestIntroBecomesOnePagePerPage(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Intro, Pages: []string{"one", "two"}}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.PageStep, curriculum.PageStep)
	if steps[1].Page != "two" {
		t.Fatalf("second page = %q", steps[1].Page)
	}
}

func TestNewKeysShowEachKeyThenDrillThem(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.NewKeys, NewKeys: []rune("dk"), Keys: []rune("fjdk")}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.KeyIntroStep, curriculum.KeyIntroStep, curriculum.DrillStep, curriculum.DrillStep)
	if steps[0].Key != 'd' || steps[1].Key != 'k' {
		t.Fatalf("introduced %q and %q", steps[0].Key, steps[1].Key)
	}
	if !strings.HasPrefix(steps[2].Text, "d d d") {
		t.Fatalf("first drill = %q", steps[2].Text)
	}
}

func TestSpaceBarIsDrilledBetweenKnownKeys(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.NewKeys, NewKeys: []rune(" "), Keys: []rune("fj ")}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.KeyIntroStep, curriculum.DrillStep)
	if steps[0].Key != ' ' || strings.Trim(steps[1].Text, "fj ") != "" || !strings.Contains(steps[1].Text, " ") {
		t.Fatalf("space drill = %q", steps[1].Text)
	}
}

func TestReviewAndHandDrillAreOneDrill(t *testing.T) {
	for _, kind := range []curriculum.Kind{curriculum.Review, curriculum.HandDrill} {
		l := curriculum.Lesson{Kind: kind, Keys: []rune("asdfg")}
		steps := l.Steps(rng())
		sameKinds(t, steps, curriculum.DrillStep)
		if strings.Trim(steps[0].Text, "asdfg ") != "" {
			t.Fatalf("%v drill = %q", kind, steps[0].Text)
		}
	}
}

func TestPracticeCanUseRealWords(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Practice, Keys: []rune("asdfjkl;"), Words: true}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.DrillStep)
	if !strings.Contains(steps[0].Text, "a") {
		t.Fatalf("word drill = %q", steps[0].Text)
	}
}

func TestPracticeFallsBackToGroupsWhenNoWordFits(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Practice, Keys: []rune("fj"), Words: true}

	steps := l.Steps(rng())

	if strings.Trim(steps[0].Text, "fj ") != "" {
		t.Fatalf("fallback drill = %q", steps[0].Text)
	}
}

func TestPlayIsAGameOverTheLessonKeys(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Play, Keys: []rune("fjdk")}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.GameStep)
	if string(steps[0].Keys) != "fjdk" {
		t.Fatalf("game keys = %q", string(steps[0].Keys))
	}
}

func TestABalloonPlayRisesWordsMadeOfTheLessonKeys(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Play, Game: curriculum.BalloonWords, Keys: []rune("asdfjkl;")}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.GameStep)
	if steps[0].Game != curriculum.BalloonWords || len(steps[0].Words) != 40 {
		t.Fatalf("step = %+v", steps[0])
	}
	for _, w := range steps[0].Words {
		if strings.Trim(w, "asdfjkl;") != "" {
			t.Fatalf("word %q uses keys outside the lesson", w)
		}
	}
}

func TestABalloonPlayWithCapitalsRisesCapitalisedWords(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Play, Game: curriculum.BalloonWords, Keys: []rune("asdfghjkl;GH")}

	step := l.Steps(rng())[0]

	if step.Game != curriculum.BalloonWords || len(step.Words) != 40 {
		t.Fatalf("step = %+v", step)
	}
	capitals := 0
	for _, w := range step.Words {
		if strings.ContainsAny(w[:1], "GH") {
			capitals++
		}
	}
	if capitals == 0 {
		t.Fatalf("no capitalised words in %v", step.Words)
	}
}

func TestABalloonPlayWithoutWordsFallsBackToFallingLetters(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Play, Game: curriculum.BalloonWords, Keys: []rune("fj")}

	step := l.Steps(rng())[0]

	if step.Game != curriculum.FallingLetters || step.Words != nil || string(step.Keys) != "fj" {
		t.Fatalf("step = %+v", step)
	}
}

func TestAPopUpPlayWarmsUpBeforeTheScoredGame(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Play, Game: curriculum.PopUpWarmUp, Keys: []rune("fj")}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.GameStep, curriculum.GameStep)
	if steps[0].Game != curriculum.PopUpWarmUp || steps[1].Game != curriculum.FallingLetters {
		t.Fatalf("steps = %+v", steps)
	}
	if string(steps[0].Keys) != "fj" || string(steps[1].Keys) != "fj" {
		t.Fatalf("keys = %q, %q", string(steps[0].Keys), string(steps[1].Keys))
	}
}

func TestTravelIsOneDrillThatStartsFromTheHomeKeys(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Travel, Keys: []rune("jum"), Focus: []rune("j"), GoalWPM: 8}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.DrillStep)
	for _, g := range strings.Fields(steps[0].Text) {
		if g != "juj" && g != "jmj" {
			t.Fatalf("drill = %q", steps[0].Text)
		}
	}
}

func TestWordPracticePrefersTheFocusKeys(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Practice, Keys: []rune("asdfghjkl;qwertyuiopvm"), Focus: []rune("vm"), Words: true, GoalWPM: 8}

	words := strings.Fields(l.Steps(rng())[0].Text)

	focused := 0
	for _, w := range words {
		if strings.ContainsAny(w, "vm") {
			focused++
		}
	}
	if focused*2 < len(words) {
		t.Fatalf("%d of %d words use v or m", focused, len(words))
	}
}

func TestTabPracticeSplitsWordPairsWithTabs(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Tab, Keys: []rune("asdfghjkl;\t"), GoalWPM: 8}

	text := l.Steps(rng())[0].Text

	if strings.Count(text, "\t") != 16 || strings.Trim(text, "asdfghjkl;\t ") != "" {
		t.Fatalf("text = %q", text)
	}
}

func TestTabPracticeWithoutWordsUsesGroups(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Tab, Keys: []rune("fj\t"), GoalWPM: 8}

	if text := l.Steps(rng())[0].Text; strings.Trim(text, "fj\t ") != "" || text == "" {
		t.Fatalf("text = %q", text)
	}
}

func TestDynamicPracticeWithoutHistoryMixesEveryKey(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Dynamic, Keys: []rune("fj"), GoalWPM: 8}

	steps := l.Steps(rng())

	sameKinds(t, steps, curriculum.DrillStep)
	if strings.Trim(steps[0].Text, "fj ") != "" {
		t.Fatalf("drill = %q", steps[0].Text)
	}
}

func TestCapitalReviewStartsEachGroupWithACapital(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Review, Keys: []rune("asdfjkl;FJ"), Focus: []rune("FJ"), GoalWPM: 8}

	for _, g := range strings.Fields(l.Steps(rng())[0].Text) {
		if !strings.ContainsRune("FJ", rune(g[0])) {
			t.Fatalf("group %q does not start with a capital", g)
		}
	}
}

func TestCapitalPracticeCapitalisesRealWords(t *testing.T) {
	l := curriculum.Lesson{Kind: curriculum.Practice, Keys: []rune("asdfghjkl;FJ"), Focus: []rune("FJ"), Words: true, GoalWPM: 8}

	capitals := 0
	for _, w := range strings.Fields(l.Steps(rng())[0].Text) {
		if strings.ContainsAny(w[1:], "FJ") {
			t.Fatalf("%q has a capital inside", w)
		}
		if strings.ContainsRune("FJ", rune(w[0])) {
			capitals++
		}
	}
	if capitals == 0 {
		t.Fatal("no capitalised words")
	}
}
