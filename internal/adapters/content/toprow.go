package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeRU   = homeRow + "ru"
	homeRUEI = homeRU + "ei"
	topSix   = "rueiwo"
	homeTop6 = homeRow + topSix
	homeQY   = homeTop6 + "qy"

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
		newKeys(35, "Keys w and o", "wo", homeTop6),
		review(36, "Review: w and o", homeTop6, "wo", topGoal),
		practice(37, "Practice: w and o", homeTop6, "wo", topGoal, true),
		play(38, "Play: top six keys", topSix),
		newKeys(39, "Keys q and y", "qy", homeQY),
		review(40, "Review: q and y", homeQY, "qy", topGoal),
		practice(41, "Practice: q and y", homeQY, "qy", topGoal, true),
	}
}

var posturePages = []string{
	"Sit back in your chair with your feet flat on the floor. A straight back lets your arms hang loose instead of holding you up.",
	"Keep your elbows close to your sides and your forearms level with the keyboard. Your wrists should float, not rest on the desk.",
	"Put the screen at eye level, about an arm's length away, so your neck stays straight while your fingers do the work.",
	"Every half hour, stand up, roll your shoulders and look at something far away. Short breaks keep you fast.",
}
