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
