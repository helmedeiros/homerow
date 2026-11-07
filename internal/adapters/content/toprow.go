package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	homeRU   = homeRow + "ru"
	homeRUEI = homeRU + "ei"
	topSix   = "rueiwo"
	homeTop6 = homeRow + topSix
	homeQY   = homeTop6 + "qy"
	homeTop  = homeQY + "pt"
	topRow   = "qwertyuiop"

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
		newKeys(42, "Keys p and t", "pt", homeTop),
		review(43, "Review: p and t", homeTop, "pt", topGoal),
		practice(44, "Practice: p and t", homeTop, "pt", topGoal, true),
		play(45, "Play: top row", topRow),
		{Number: 46, Title: "Think in words", Kind: curriculum.Intro, Pages: thinkInWordsPages},
		practice(47, "Review: home + top", homeTop, "", topGoal, true),
		handDrill(48, "Top row: left hand", leftHand+"qwert", topGoal),
		handDrill(49, "Top row: right hand", rightHand+"yuiop", topGoal),
		play(50, "Play: top row", topRow),
		play(51, "Play: home and top rows", homeRow+topRow),
	}
}

var thinkInWordsPages = []string{
	"Beginners spell every word one letter at a time. Fast typists see a word and let their fingers play it as one movement.",
	"Common pieces like the, ing and er soon turn into single gestures. Type them often and they stop needing thought.",
	"Read a word or two ahead of where you are typing. Your fingers stay busy while your eyes find the next word.",
	"If you stumble, keep a steady rhythm instead of rushing. Even beats make words flow, and speed follows.",
}

var posturePages = []string{
	"Sit back in your chair with your feet flat on the floor. A straight back lets your arms hang loose instead of holding you up.",
	"Keep your elbows close to your sides and your forearms level with the keyboard. Your wrists should float, not rest on the desk.",
	"Put the screen at eye level, about an arm's length away, so your neck stays straight while your fingers do the work.",
	"Every half hour, stand up, roll your shoulders and look at something far away. Short breaks keep you fast.",
}
