package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

const (
	symbolsDA  = allKeys + "$&"
	symbolsHS  = symbolsDA + "#*"
	symbolsAO  = symbolsHS + "@("
	symbolsBC  = symbolsAO + "!)"
	symbols    = "$&#*@(!)%^"
	allSymbols = allKeys + symbols

	symbolGoal = 8
)

func Symbols() curriculum.Module {
	return curriculum.Module{Number: 6, Title: "Symbols", Lessons: symbolLessons()}
}

func symbolLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		newKeys(158, "Keys $ and &", "$&", symbolsDA),
		review(159, "Review: $ and &", symbolsDA, "$&", symbolGoal),
		practice(160, "Practice: $ and &", symbolsDA, "$&", symbolGoal, true),
		newKeys(161, "Keys # and *", "#*", symbolsHS),
		review(162, "Review: # and *", symbolsHS, "#*", symbolGoal),
		practice(163, "Practice: # and *", symbolsHS, "#*", symbolGoal, true),
		newKeys(164, "Keys @ and (", "@(", symbolsAO),
		review(165, "Review: @ and (", symbolsAO, "@(", symbolGoal),
		practice(166, "Practice: @ and (", symbolsAO, "@(", symbolGoal, true),
		newKeys(167, "Keys ! and )", "!)", symbolsBC),
		review(168, "Review: ! and )", symbolsBC, "!)", symbolGoal),
		practice(169, "Practice: ! and )", symbolsBC, "!)", symbolGoal, true),
		newKeys(170, "Keys % and ^", "%^", allSymbols),
		review(171, "Review: % and ^", allSymbols, "%^", symbolGoal),
		practice(172, "Practice: % and ^", allSymbols, "%^", symbolGoal, true),
		balloons(173, "Play: every symbol", allSymbols),
		travel(174, "Travel: $", "f4$", "f", symbolGoal),
		travel(175, "Travel: #", "d3#", "d", symbolGoal),
		travel(176, "Travel: @", "s2@", "s", symbolGoal),
		travel(177, "Travel: !", "a1!", "a", symbolGoal),
		travel(178, "Travel: %", "f5%", "f", symbolGoal),
		travel(179, "Travel: &", "j7&", "j", symbolGoal),
		travel(180, "Travel: *", "k8*", "k", symbolGoal),
		travel(181, "Travel: (", "l9(", "l", symbolGoal),
		travel(182, "Travel: )", ";0)", ";", symbolGoal),
		travel(183, "Travel: ^", "j6^", "j", symbolGoal),
		travel(184, "Travel: left hand", "asdf12345!@#$%", "asdf", symbolGoal),
		travel(185, "Travel: right hand", "jkl;67890^&*()", "jkl;", symbolGoal),
		review(186, "Review: every symbol", allSymbols, symbols, symbolGoal),
		practice(187, "Practice: every symbol", allSymbols, symbols, symbolGoal, true),
	}
}
