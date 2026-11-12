package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	capsFJ = allRows + "FJ"
	capsDK = capsFJ + "DK"

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
	}
}

var shiftPages = []string{
	"Capitals need the shift key. There is one at each end of the bottom row, right under your pinkies.",
	"Hold shift with the pinky of the other hand, press the letter, then let go of both. A capital F takes the right shift, a capital J the left one.",
	"The typing hand stays on its home keys, so every reach is the same as for the small letter. Only the pinky of the other hand moves.",
	"Release shift as soon as the capital is typed. Holding it too long turns the next letter into a capital as well.",
}
