package curriculum

import (
	"errors"
	"fmt"
)

var ErrInvalidCourse = errors.New("invalid course")

type Course struct {
	Modules []Module
}

func (c Course) Lesson(number int) (Lesson, Module, bool) {
	for _, m := range c.Modules {
		if l, ok := m.Lesson(number); ok {
			return l, m, true
		}
	}
	return Lesson{}, Module{}, false
}

func (c Course) Lessons() []Lesson {
	var all []Lesson
	for _, m := range c.Modules {
		all = append(all, m.Lessons...)
	}
	return all
}

func (c Course) Next(number int) int {
	all := c.Lessons()
	for i, l := range all {
		if l.Number == number && i+1 < len(all) {
			return all[i+1].Number
		}
	}
	return 0
}

func (c Course) Validate() error {
	if len(c.Modules) == 0 {
		return fmt.Errorf("%w: no modules", ErrInvalidCourse)
	}
	last := 0
	for _, m := range c.Modules {
		if m.Title == "" || len(m.Lessons) == 0 {
			return fmt.Errorf("%w: module %d needs a title and lessons", ErrInvalidCourse, m.Number)
		}
		for _, l := range m.Lessons {
			if err := l.Validate(); err != nil {
				return fmt.Errorf("%w: %w", ErrInvalidCourse, err)
			}
			if l.Number <= last {
				return fmt.Errorf("%w: lesson %d comes after %d", ErrInvalidCourse, l.Number, last)
			}
			last = l.Number
		}
	}
	return nil
}
