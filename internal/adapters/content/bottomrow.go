package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeTopVM = homeTop + "vm"
	homeTopC  = homeTopVM + "c,"
	homeTopX  = homeTopC + "x."

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
		{Number: 63, Title: "History of QWERTY", Kind: curriculum.Intro, Pages: qwertyPages},
		newKeys(64, "Keys x and period", "x.", homeTopX),
		review(65, "Review: x and period", homeTopX, "x.", bottomGoal),
		practice(66, "Practice: x and period", homeTopX, "x.", bottomGoal, true),
		travel(67, "Travel: right ring", homeTopX, "l", bottomGoal),
		travel(68, "Travel: left ring", homeTopX, "s", bottomGoal),
	}
}

var qwertyPages = []string{
	"The QWERTY layout comes from typewriters of the 1870s. Its letters are named after the first six keys of the top row.",
	"Early typewriters jammed when neighbouring type bars swung up together, so common letter pairs were spread across the keyboard.",
	"Jams are long gone, but the layout stayed. Millions of people already knew it, and every keyboard since has followed.",
	"Your fingers do not care where the letters came from. Learn each reach once and QWERTY becomes second nature.",
}
