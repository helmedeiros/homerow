package content_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func shownKeys(l curriculum.Lesson, seed uint64) string {
	shown := string(l.Keys)
	for _, step := range l.Steps(rand.New(rand.NewPCG(seed, seed))) {
		shown += step.Text + string(step.Keys) + strings.Join(step.Words, "")
		if step.Kind == curriculum.KeyIntroStep {
			shown += string(step.Key)
		}
	}
	return shown
}

func TestEveryLessonShowsOnlyKeysIntroducedSoFar(t *testing.T) {
	introduced := ""
	for _, l := range content.Course().Lessons() {
		introduced += string(l.NewKeys)
		for seed := range uint64(5) {
			for _, r := range shownKeys(l, seed) {
				if !strings.ContainsRune(introduced, r) {
					t.Fatalf("lesson %d %q shows %q before it is introduced", l.Number, l.Title, r)
				}
			}
		}
	}
}
