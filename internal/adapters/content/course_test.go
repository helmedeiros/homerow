package content_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/content"
)

func TestCourseIsValid(t *testing.T) {
	if err := content.Course().Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestCourseNumbersLessonsWithoutGaps(t *testing.T) {
	for i, l := range content.Course().Lessons() {
		if l.Number != i+1 {
			t.Fatalf("lesson at %d is numbered %d", i, l.Number)
		}
	}
}

func TestCourseNumbersModulesInOrder(t *testing.T) {
	for i, m := range content.Course().Modules {
		if m.Number != i+1 {
			t.Fatalf("module at %d is numbered %d", i, m.Number)
		}
	}
}

func TestCourseOpensWithTheHomeRow(t *testing.T) {
	first := content.Course().Modules[0]

	if first.Title != "Home Row" || len(first.Lessons) != 23 {
		t.Fatalf("first module = %q with %d lessons", first.Title, len(first.Lessons))
	}
}

func TestCatalogServesTheCourse(t *testing.T) {
	if got, want := len((content.Catalog{}).Course().Lessons()), len(content.Course().Lessons()); got != want {
		t.Fatalf("catalog serves %d lessons, want %d", got, want)
	}
}
