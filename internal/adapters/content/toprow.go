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
		{Number: 30, Title: "Sit straight", Kind: curriculum.Intro, Pages: posturePages},
		practice(31, "Review: home + ruei", homeRUEI, "ruei", topGoal, true),
		handDrill(32, "Top row: left hand", leftHand+"re", topGoal),
		handDrill(33, "Top row: right hand", rightHand+"ui", topGoal),
		play(34, "Play: ruei", "ruei"),
	}
}

var posturePages = []string{
	"Sit back in your chair with your feet flat on the floor. A straight back lets your arms hang loose instead of holding you up.",
	"Keep your elbows close to your sides and your forearms level with the keyboard. Your wrists should float, not rest on the desk.",
	"Put the screen at eye level, about an arm's length away, so your neck stays straight while your fingers do the work.",
	"Every half hour, stand up, roll your shoulders and look at something far away. Short breaks keep you fast.",
}
