package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeTopVM = homeTop + "vm"
	homeTopC  = homeTopVM + "c,"

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
		newKeys(57, "Keys c and comma", "c,", homeTopC),
		review(58, "Review: c and comma", homeTopC, "c,", bottomGoal),
		practice(59, "Practice: c and comma", homeTopC, "c,", bottomGoal, true),
		travel(60, "Travel: right middle", homeTopC, "k", bottomGoal),
		travel(61, "Travel: left middle", homeTopC, "d", bottomGoal),
		balloons(62, "Play: v m c ,", homeTopC),
	}
}
