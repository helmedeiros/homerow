package exercise_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func TestTravelGoesOutFromAHomeKeyAndBack(t *testing.T) {
	text := exercise.Travel(rand.New(rand.NewPCG(1, 1)), []rune("j"), []rune("asdfjkl;um"), 30)

	groups := strings.Fields(text)
	if len(groups) != 30 {
		t.Fatalf("groups = %d", len(groups))
	}
	seen := map[string]bool{}
	for _, g := range groups {
		if g != "juj" && g != "jmj" {
			t.Fatalf("group %q does not travel from j to one of its reaches", g)
		}
		seen[g] = true
	}
	if len(seen) != 2 {
		t.Fatalf("only saw %v", seen)
	}
}

func TestTravelMixesSeveralHomeKeys(t *testing.T) {
	text := exercise.Travel(rand.New(rand.NewPCG(2, 2)), []rune("fj"), []rune("fjrvum"), 40)

	for _, g := range strings.Fields(text) {
		if g[0] != g[2] || !strings.ContainsRune("fj", rune(g[0])) {
			t.Fatalf("group %q", g)
		}
		if strings.ContainsRune("rv", rune(g[1])) != (g[0] == 'f') {
			t.Fatalf("group %q reaches with the wrong finger", g)
		}
	}
}

func TestTravelWithoutReachesStaysHome(t *testing.T) {
	if text := exercise.Travel(rand.New(rand.NewPCG(1, 1)), []rune("j"), []rune("fj"), 3); text != "jjj jjj jjj" {
		t.Fatalf("text = %q", text)
	}
	if text := exercise.Travel(rand.New(rand.NewPCG(1, 1)), []rune("€"), []rune("fj€"), 1); text != "€€€" {
		t.Fatalf("text = %q", text)
	}
}
