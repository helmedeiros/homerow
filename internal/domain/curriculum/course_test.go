package curriculum_test

import (
	"errors"
	"testing"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

func intro(number int) curriculum.Lesson {
	return curriculum.Lesson{Number: number, Title: "Lesson", Kind: curriculum.Intro, Pages: []string{"page"}}
}

func sampleCourse() curriculum.Course {
	return curriculum.Course{Modules: []curriculum.Module{
		{Number: 1, Title: "Home Row", Lessons: []curriculum.Lesson{intro(1), intro(2)}},
		{Number: 2, Title: "Top Row", Lessons: []curriculum.Lesson{intro(3)}},
	}}
}

func TestCourseFindsALessonAndItsModule(t *testing.T) {
	lesson, module, ok := sampleCourse().Lesson(3)

	if !ok || lesson.Number != 3 || module.Title != "Top Row" {
		t.Fatalf("Lesson(3) = %d, %q, %v", lesson.Number, module.Title, ok)
	}
	if _, _, ok := sampleCourse().Lesson(9); ok {
		t.Fatal("Lesson(9) should not exist")
	}
}

func TestCourseListsEveryLessonInOrder(t *testing.T) {
	var numbers []int
	for _, l := range sampleCourse().Lessons() {
		numbers = append(numbers, l.Number)
	}

	if len(numbers) != 3 || numbers[0] != 1 || numbers[2] != 3 {
		t.Fatalf("Lessons() = %v", numbers)
	}
}

func TestNextLessonCrossesIntoTheNextModule(t *testing.T) {
	c := sampleCourse()

	if c.Next(1) != 2 || c.Next(2) != 3 || c.Next(3) != 0 || c.Next(9) != 0 {
		t.Fatalf("Next = %d %d %d %d", c.Next(1), c.Next(2), c.Next(3), c.Next(9))
	}
}

func TestValidCourse(t *testing.T) {
	if err := sampleCourse().Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidCourses(t *testing.T) {
	cases := map[string]curriculum.Course{
		"no modules":      {},
		"empty module":    {Modules: []curriculum.Module{{Number: 1, Title: "Home Row"}}},
		"untitled module": {Modules: []curriculum.Module{{Number: 1, Lessons: []curriculum.Lesson{intro(1)}}}},
		"repeated number": {Modules: []curriculum.Module{
			{Number: 1, Title: "A", Lessons: []curriculum.Lesson{intro(1)}},
			{Number: 2, Title: "B", Lessons: []curriculum.Lesson{intro(1)}},
		}},
		"numbers out of order": {Modules: []curriculum.Module{
			{Number: 1, Title: "A", Lessons: []curriculum.Lesson{intro(2), intro(1)}},
		}},
		"invalid lesson": {Modules: []curriculum.Module{
			{Number: 1, Title: "A", Lessons: []curriculum.Lesson{{Number: 1}}},
		}},
	}
	for name, c := range cases {
		if err := c.Validate(); !errors.Is(err, curriculum.ErrInvalidCourse) {
			t.Errorf("%s: err = %v, want ErrInvalidCourse", name, err)
		}
	}
}
