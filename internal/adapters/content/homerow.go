package content

import (
	"time"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
)

const (
	fj        = "fj"
	fjdk      = "fjdk"
	fjdksl    = "fjdksl"
	firstHome = "fjdksla;"
	leftHand  = "asdfg"
	rightHand = "hjkl;"
	homeRow   = "asdfghjkl;"

	starterGoal = 8
	middleGoal  = 10
	homeGoal    = 12
	playGoal    = 8

	practiceTime = time.Minute
)

func HomeRow() curriculum.Module {
	return curriculum.Module{Number: 1, Title: "Home Row", Lessons: homeRowLessons()}
}

func homeRowLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		{Number: 1, Title: "Welcome to touch typing", Kind: curriculum.Intro, Pages: welcomePages},
		newKeys(2, "Keys f and j", fj, fj),
		newKeys(3, "Space bar", " ", fj+" "),
		review(4, "Review: f and j", fj, fj, starterGoal),
		newKeys(5, "Keys d and k", "dk", fjdk),
		review(6, "Review: d and k", fjdk, "dk", starterGoal),
		practice(7, "Practice: d and k", fjdk, "dk", starterGoal, false),
		play(8, "Play: f j d k", fjdk),
		newKeys(9, "Keys s and l", "sl", fjdksl),
		review(10, "Review: s and l", fjdksl, "sl", middleGoal),
		practice(11, "Practice: s and l", fjdksl, "sl", middleGoal, false),
		newKeys(12, "Keys a and ;", "a;", firstHome),
		review(13, "Review: a and ;", firstHome, "a;", middleGoal),
		practice(14, "First eight keys", firstHome, "", middleGoal, true),
		play(15, "Play: first eight keys", firstHome),
		{Number: 16, Title: "Finding home", Kind: curriculum.Intro, Pages: findingHomePages},
		handDrill(17, "Home row: left hand", leftHand, homeGoal),
		handDrill(18, "Home row: right hand", rightHand, homeGoal),
		newKeys(19, "Keys g and h", "gh", homeRow),
		review(20, "Review: g and h", homeRow, "gh", homeGoal),
		practice(21, "Practice: g and h", homeRow, "gh", homeGoal, false),
		practice(22, "Home row review", homeRow, "", homeGoal, true),
		play(23, "Play: home row", homeRow),
	}
}

var welcomePages = []string{
	"Touch typing means typing without looking at the keyboard. Each finger owns a few keys, so your hands barely move and your eyes stay on the screen.",
	"Rest your fingers on the home row: the left hand on a s d f, the right hand on j k l ;. Your thumbs hover over the space bar.",
	"Feel for the small bumps on f and j. They let your index fingers find home without looking down.",
	"Accuracy first, speed later. Every exercise earns up to five stars, and you pass it with 80% accuracy or more.",
}

var findingHomePages = []string{
	"Whenever you lose your place, slide your index fingers back to the bumps on f and j. The other fingers fall into place beside them.",
	"Keep your wrists relaxed and slightly raised, with your fingers curved as if holding a ball.",
	"g and h are one step inward from f and j. Reach with the index finger, press, and come straight back home.",
}

func handDrill(number int, title, keys string, goal float64) curriculum.Lesson {
	return curriculum.Lesson{Number: number, Title: title, Kind: curriculum.HandDrill, Keys: []rune(keys), GoalWPM: goal}
}

func newKeys(number int, title, keys, allowed string) curriculum.Lesson {
	return curriculum.Lesson{Number: number, Title: title, Kind: curriculum.NewKeys, NewKeys: []rune(keys), Keys: []rune(allowed)}
}

func review(number int, title, keys, focus string, goal float64) curriculum.Lesson {
	return curriculum.Lesson{Number: number, Title: title, Kind: curriculum.Review, Keys: []rune(keys), Focus: []rune(focus), GoalWPM: goal}
}

func practice(number int, title, keys, focus string, goal float64, words bool) curriculum.Lesson {
	return curriculum.Lesson{
		Number: number, Title: title, Kind: curriculum.Practice,
		Keys: []rune(keys), Focus: []rune(focus), Words: words,
		GoalWPM: goal, TimeLimit: practiceTime,
	}
}

func play(number int, title, keys string) curriculum.Lesson {
	return curriculum.Lesson{Number: number, Title: title, Kind: curriculum.Play, Keys: []rune(keys), GoalWPM: playGoal}
}

func travel(number int, title, keys, homes string, goal float64) curriculum.Lesson {
	return curriculum.Lesson{Number: number, Title: title, Kind: curriculum.Travel, Keys: []rune(keys), Focus: []rune(homes), GoalWPM: goal}
}

func popUps(number int, title, keys string) curriculum.Lesson {
	l := play(number, title, keys)
	l.Game = curriculum.PopUpWarmUp
	return l
}

func balloons(number int, title, keys string) curriculum.Lesson {
	l := play(number, title, keys)
	l.Game = curriculum.BalloonWords
	return l
}
