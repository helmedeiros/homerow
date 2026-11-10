package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeTopVM = homeTop + "vm"
	homeTopC  = homeTopVM + "c,"
	homeTopX  = homeTopC + "x."
	homeTopZ  = homeTopX + "z/"
	allRows   = homeTopZ + "bn"
	bottomRow = "zxcvbnm,./"

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
		newKeys(69, "Keys z and slash", "z/", homeTopZ),
		review(70, "Review: z and slash", homeTopZ, "z/", bottomGoal),
		practice(71, "Practice: z and slash", homeTopZ, "z/", bottomGoal, true),
		travel(72, "Travel: right pinky", homeTopZ, ";", bottomGoal),
		travel(73, "Travel: left pinky", homeTopZ, "a", bottomGoal),
		balloons(74, "Play: bottom row", homeTopZ),
		newKeys(75, "Keys b and n", "bn", allRows),
		review(76, "Review: b and n", allRows, "bn", bottomGoal),
		practice(77, "Practice: b and n", allRows, "bn", bottomGoal, true),
		{Number: 78, Title: "Take a break", Kind: curriculum.Intro, Pages: breakPages},
		handDrill(79, "Bottom row: left hand", leftHand+"zxcvb", bottomGoal),
		handDrill(80, "Bottom row: right hand", rightHand+"nm,./", bottomGoal),
		practice(81, "Practice: bottom row", allRows, bottomRow, bottomGoal, true),
		popUps(82, "Play: bottom row", bottomRow),
		play(83, "Play: every letter", allRows),
		travel(84, "Travel: right hand, all fingers", allRows, "jkl;", bottomGoal),
		travel(85, "Travel: left hand, all fingers", allRows, "asdf", bottomGoal),
	}
}

var qwertyPages = []string{
	"The QWERTY layout comes from typewriters of the 1870s. Its letters are named after the first six keys of the top row.",
	"Early typewriters jammed when neighbouring type bars swung up together, so common letter pairs were spread across the keyboard.",
	"Jams are long gone, but the layout stayed. Millions of people already knew it, and every keyboard since has followed.",
	"Your fingers do not care where the letters came from. Learn each reach once and QWERTY becomes second nature.",
}

var breakPages = []string{
	"You now know every letter key. That is a good moment to stop for a minute and let your hands rest.",
	"Shake out your hands, stretch your fingers wide, then make loose fists a few times. Tight hands type slower.",
	"Look away from the screen at something far off and blink a few times. Your eyes work hard while you learn.",
	"When you come back, place your fingers on the home row, feel the bumps on f and j, and carry on.",
}
