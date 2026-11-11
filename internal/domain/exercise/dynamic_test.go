package exercise_test

import (
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/exercise"
)

func TestDynamicDrillsPutAWeakKeyInEveryGroup(t *testing.T) {
	keys := []rune("asdfghjkl;")
	text := exercise.Dynamic(seeded(5), keys, []rune("gh"), 200)

	groups := strings.Fields(text)
	if len(groups) != 200 {
		t.Fatalf("groups = %d", len(groups))
	}
	weak := 0
	for _, g := range groups {
		if strings.Trim(g, string(keys)) != "" || !strings.ContainsAny(g, "gh") {
			t.Fatalf("group %q", g)
		}
		weak += strings.Count(g, "g") + strings.Count(g, "h")
	}
	if weak < len(strings.ReplaceAll(text, " ", ""))/2 {
		t.Fatalf("only %d weak keys in %q", weak, text)
	}
}

func TestDynamicWithoutWeakKeysMixesEveryKey(t *testing.T) {
	text := exercise.Dynamic(seeded(5), []rune("fj"), nil, 10)

	if len(strings.Fields(text)) != 10 || strings.Trim(text, "fj ") != "" {
		t.Fatalf("text = %q", text)
	}
}
