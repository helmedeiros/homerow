package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeTopVM = homeTop + "vm"

	bottomGoal = 12
)

func BottomRow() curriculum.Module {
	return curriculum.Module{Number: 3, Title: "Bottom Row", Lessons: bottomRowLessons()}
}

func bottomRowLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		newKeys(52, "Keys v and m", "vm", homeTopVM),
		review(53, "Review: v and m", homeTopVM, "vm", bottomGoal),
		practice(54, "Practice: v and m", homeTopVM, "vm", bottomGoal, true),
		travel(55, "Travel: right index", homeTopVM, "j", bottomGoal),
		travel(56, "Travel: left index", homeTopVM, "f", bottomGoal),
	}
}
