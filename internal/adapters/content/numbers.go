package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	digits47 = capsXB + "47"
	digits38 = digits47 + "38"
	digits29 = digits38 + "29"
	digits10 = digits29 + "10"
	allKeys  = capsXB + "1234567890"

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
		newKeys(149, "Keys 2 and 9", "29", digits29),
		review(150, "Review: 2 and 9", digits29, "29", numberGoal),
		practice(151, "Practice: 2 and 9", digits29, "29", numberGoal, true),
		newKeys(152, "Keys 1 and 0", "10", digits10),
		review(153, "Review: 1 and 0", digits10, "10", numberGoal),
		practice(154, "Practice: 1 and 0", digits10, "10", numberGoal, true),
		newKeys(155, "Keys 5 and 6", "56", allKeys),
		review(156, "Review: 5 and 6", allKeys, "56", numberGoal),
		practice(157, "Practice: 5 and 6", allKeys, "56", numberGoal, true),
	}
}
