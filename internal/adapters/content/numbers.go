package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	digits47 = capsXB + "47"
	digits38 = digits47 + "38"

	numberGoal = 10
)

func Numbers() curriculum.Module {
	return curriculum.Module{Number: 5, Title: "Numbers", Lessons: numberLessons()}
}

func numberLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		newKeys(143, "Keys 4 and 7", "47", digits47),
		review(144, "Review: 4 and 7", digits47, "47", numberGoal),
		practice(145, "Practice: 4 and 7", digits47, "47", numberGoal, true),
		newKeys(146, "Keys 3 and 8", "38", digits38),
		review(147, "Review: 3 and 8", digits38, "38", numberGoal),
		practice(148, "Practice: 3 and 8", digits38, "38", numberGoal, true),
	}
}
