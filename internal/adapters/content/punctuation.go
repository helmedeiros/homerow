package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	marksDash     = allSymbols + "-'"
	marksBrackets = marksDash + "[]"
	marksEquals   = marksBrackets + "=\\"
	marksPlus     = marksEquals + "_+"
	marksBraces   = marksPlus + "{}"

	punctuationGoal = 8
)

func Punctuation() curriculum.Module {
	return curriculum.Module{Number: 7, Title: "Punctuation", Lessons: punctuationLessons()}
}

func punctuationLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		{Number: 188, Title: "The pinky's corner", Kind: curriculum.Intro, Pages: punctuationPages},
		newKeys(189, "Keys - and '", "-'", marksDash),
		review(190, "Review: - and '", marksDash, "-'", punctuationGoal),
		practice(191, "Practice: - and '", marksDash, "-'", punctuationGoal, true),
		newKeys(192, "Keys [ and ]", "[]", marksBrackets),
		review(193, "Review: [ and ]", marksBrackets, "[]", punctuationGoal),
		practice(194, "Practice: [ and ]", marksBrackets, "[]", punctuationGoal, true),
		newKeys(195, "Keys = and \\", "=\\", marksEquals),
		review(196, "Review: = and \\", marksEquals, "=\\", punctuationGoal),
		practice(197, "Practice: = and \\", marksEquals, "=\\", punctuationGoal, true),
		newKeys(198, "Keys _ and +", "_+", marksPlus),
		review(199, "Review: _ and +", marksPlus, "_+", punctuationGoal),
		practice(200, "Practice: _ and +", marksPlus, "_+", punctuationGoal, true),
		newKeys(201, "Keys { and }", "{}", marksBraces),
		review(202, "Review: { and }", marksBraces, "{}", punctuationGoal),
		practice(203, "Practice: { and }", marksBraces, "{}", punctuationGoal, true),
	}
}

var punctuationPages = []string{
	"Most punctuation sits to the right of p, ; and 0, where only the right pinky reaches. Two keys live elsewhere: the backtick at the far left of the number row, under the left pinky, and the backslash at the end of the top row.",
	"Reach with the pinky alone and keep the other fingers resting on their home keys. The hand may lean a little, but it comes straight back to the home row.",
	"Brackets, braces and quotes come in pairs and wrap a word, as in [list] or \"quote\". A dash, an underscore or an equals sign joins two words.",
	"_ + { } \" | < > and ~ need shift. Hold it with the pinky of the other hand, just like a capital.",
}
