package exercise_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func seeded(seed uint64) *rand.Rand {
	return rand.New(rand.NewPCG(seed, seed))
}

func TestGroupsAreDeterministicForASeed(t *testing.T) {
	mix := exercise.Mix{Keys: []rune("fjdk")}

	a := exercise.Groups(seeded(7), mix, 12)
	b := exercise.Groups(seeded(7), mix, 12)

	if a != b {
		t.Fatalf("same seed produced %q and %q", a, b)
	}
}

func TestGroupsUseOnlyTheGivenKeysInShortGroups(t *testing.T) {
	mix := exercise.Mix{Keys: []rune("asdfjkl;"), Focus: []rune("a;")}
	for seed := uint64(0); seed < 300; seed++ {
		text := exercise.Groups(seeded(seed), mix, 15)
		groups := strings.Split(text, " ")
		if len(groups) != 15 {
			t.Fatalf("seed %d: %d groups, want 15", seed, len(groups))
		}
		for _, g := range groups {
			if len(g) < 2 || len(g) > 5 {
				t.Fatalf("seed %d: group %q has length %d", seed, g, len(g))
			}
			for _, r := range g {
				if !strings.ContainsRune("asdfjkl;", r) {
					t.Fatalf("seed %d: unexpected key %q in %q", seed, r, text)
				}
			}
		}
	}
}

func TestFocusKeysAppearAboutHalfTheTime(t *testing.T) {
	mix := exercise.Mix{Keys: []rune("fjdkslgh"), Focus: []rune("gh")}
	text := strings.ReplaceAll(exercise.Groups(seeded(1), mix, 400), " ", "")

	focus := strings.Count(text, "g") + strings.Count(text, "h")
	share := float64(focus) / float64(len(text))
	if share < 0.45 || share > 0.70 {
		t.Fatalf("focus share = %.2f, want roughly half plus their share of the rest", share)
	}
}
