package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeRU   = homeRow + "ru"
	homeRUEI = homeRU + "ei"

	topGoal = 12
)

func TopRow() curriculum.Module {
	return curriculum.Module{Number: 2, Title: "Top Row", Lessons: topRowLessons()}
}

func topRowLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		newKeys(24, "Keys r and u", "ru", homeRU),
		review(25, "Review: r and u", homeRU, "ru", topGoal),
		practice(26, "Practice: r and u", homeRU, "ru", topGoal, true),
		newKeys(27, "Keys e and i", "ei", homeRUEI),
		review(28, "Review: e and i", homeRUEI, "ei", topGoal),
		practice(29, "Practice: e and i", homeRUEI, "ei", topGoal, true),
	}
}
