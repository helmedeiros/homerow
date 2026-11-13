package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	capsFJ = allRows + "FJ"
	capsDK = capsFJ + "DK"
	capsSL = capsDK + "SL"
	capsA  = capsSL + "A:"
	capsGH = capsA + "GH"
	capsTY = capsGH + "TY"
	capsRU = capsTY + "RU"
	capsEI = capsRU + "EI"

	shiftGoal = 12
)

func ShiftKey() curriculum.Module {
	return curriculum.Module{Number: 4, Title: "Shift Key", Lessons: shiftKeyLessons()}
}

func shiftKeyLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		{Number: 89, Title: "The shift key", Kind: curriculum.Intro, Pages: shiftPages},
		newKeys(90, "Capitals F and J", "FJ", capsFJ),
		review(91, "Review: F and J", capsFJ, "FJ", shiftGoal),
		practice(92, "Practice: F and J", capsFJ, "FJ", shiftGoal, true),
		newKeys(93, "Capitals D and K", "DK", capsDK),
		review(94, "Review: D and K", capsDK, "DK", shiftGoal),
		practice(95, "Practice: D and K", capsDK, "DK", shiftGoal, true),
		play(96, "Play: F J D K", "fjdkFJDK"),
		newKeys(97, "Capitals S and L", "SL", capsSL),
		review(98, "Review: S and L", capsSL, "SL", shiftGoal),
		practice(99, "Practice: S and L", capsSL, "SL", shiftGoal, true),
		newKeys(100, "Capital A and colon", "A:", capsA),
		review(101, "Review: A and colon", capsA, "A:", shiftGoal),
		practice(102, "Practice: A and colon", capsA, "A:", shiftGoal, true),
		popUps(103, "Play: home row capitals", firstHome+"FJDKSLA:"),
		newKeys(104, "Capitals G and H", "GH", capsGH),
		review(105, "Review: G and H", capsGH, "GH", shiftGoal),
		practice(106, "Practice: G and H", capsGH, "GH", shiftGoal, true),
		newKeys(107, "Capitals T and Y", "TY", capsTY),
		review(108, "Review: T and Y", capsTY, "TY", shiftGoal),
		practice(109, "Practice: T and Y", capsTY, "TY", shiftGoal, true),
		balloons(110, "Play: G H T Y", homeTop+"GHTY"),
		newKeys(111, "Capitals R and U", "RU", capsRU),
		review(112, "Review: R and U", capsRU, "RU", shiftGoal),
		practice(113, "Practice: R and U", capsRU, "RU", shiftGoal, true),
		newKeys(114, "Capitals E and I", "EI", capsEI),
		review(115, "Review: E and I", capsEI, "EI", shiftGoal),
		practice(116, "Practice: E and I", capsEI, "EI", shiftGoal, true),
		popUps(117, "Play: R U E I", "rueiRUEI"),
	}
}

var shiftPages = []string{
	"Capitals need the shift key. There is one at each end of the bottom row, right under your pinkies.",
	"Hold shift with the pinky of the other hand, press the letter, then let go of both. A capital F takes the right shift, a capital J the left one.",
	"The typing hand stays on its home keys, so every reach is the same as for the small letter. Only the pinky of the other hand moves.",
	"Release shift as soon as the capital is typed. Holding it too long turns the next letter into a capital as well.",
}
