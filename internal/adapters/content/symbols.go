package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	symbolsDA = allKeys + "$&"
	symbolsHS = symbolsDA + "#*"

	symbolGoal = 8
)

func Symbols() curriculum.Module {
	return curriculum.Module{Number: 6, Title: "Symbols", Lessons: symbolLessons()}
}

func symbolLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		newKeys(158, "Keys $ and &", "$&", symbolsDA),
		review(159, "Review: $ and &", symbolsDA, "$&", symbolGoal),
		practice(160, "Practice: $ and &", symbolsDA, "$&", symbolGoal, true),
		newKeys(161, "Keys # and *", "#*", symbolsHS),
		review(162, "Review: # and *", symbolsHS, "#*", symbolGoal),
		practice(163, "Practice: # and *", symbolsHS, "#*", symbolGoal, true),
	}
}
