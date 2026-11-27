package content

import (
	"strings"

	"github.com/helmedeiros/homerow/internal/domain/curriculum"
	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

const codeGoal = 8

func Go() curriculum.Module {
	return curriculum.Module{Number: 8, Title: "Go", Lessons: goLessons()}
}

func goLessons() []curriculum.Lesson {
	return []curriculum.Lesson{
		{Number: 219, Title: "Typing Go", Kind: curriculum.Intro, Pages: goPages},
		introducing("\n", code(220, "var, := and const",
			"Inside a function, := declares a variable and gives it a value in one go, and Go works out the type. Use var for a zero value or a type that differs from the value, and const for values fixed at compile time.",
			"const maxRetries = 3\nvar total int\nname := \"homerow\"",
			"count := 0\nratio := 0.75\nvar ready bool",
		)),
		code(221, "Practice: declarations",
			"The three forms side by side. := needs at least one new name on its left, and a const block groups related constants.",
			"limit := 10\nvar buf []byte\nconst unit = \"wpm\"",
			"x, y := 1, 2\nx, z := 3, 4\nvar err error",
			"const (\n\tlow  = 8\n\thigh = 40\n)",
		),
		code(222, "Functions",
			"A function lists its parameters with their types, then its result type. Parameters of the same type can share it, as in a, b int.",
			"func add(a, b int) int {\n\treturn a + b\n}",
			"func greet(name string) string {\n\treturn \"hi \" + name\n}",
		),
		code(223, "Multiple returns",
			"A Go function can return several values. The last one is often an error, and the caller takes both at once with :=.",
			"func divide(a, b float64) (float64, error) {\n\tif b == 0 {\n\t\treturn 0, errDivide\n\t}\n\treturn a / b, nil\n}",
			"q, err := divide(10, 4)",
		),
		code(224, "Practice: functions",
			"Small functions with clear results. Blank identifiers like _ drop the values you do not need.",
			"func minMax(xs []int) (int, int) {\n\tlo, hi := xs[0], xs[0]\n\tfor _, x := range xs {\n\t\tlo = min(lo, x)\n\t\thi = max(hi, x)\n\t}\n\treturn lo, hi\n}",
			"func swap(a, b string) (string, string) {\n\treturn b, a\n}",
		),
		code(225, "if err != nil",
			"Errors are values. Check each one right where it happens and return early, so the happy path stays at the left edge of the code.",
			"f, err := os.Open(path)\nif err != nil {\n\treturn err\n}\ndefer f.Close()",
			"n, err := strconv.Atoi(s)\nif err != nil {\n\treturn 0, err\n}",
		),
		code(226, "Wrapping errors",
			"Add context as an error travels up with fmt.Errorf and the %w verb. The message says what failed, and the caller can still reach the original error.",
			"if err != nil {\n\treturn fmt.Errorf(\"open %s: %w\", path, err)\n}",
			"return fmt.Errorf(\"parse port %q: %w\", s, err)",
		),
		code(227, "Sentinel errors",
			"Declare errors the caller may want to recognise with errors.New, and compare with errors.Is. Unlike ==, it also finds them inside wrapped errors.",
			"var ErrNotFound = errors.New(\"not found\")",
			"if errors.Is(err, ErrNotFound) {\n\treturn nil\n}",
		),
		code(228, "Practice: errors",
			"One function that checks, wraps and returns. Every error leaves with context, and success returns nil for the error.",
			"func load(path string) ([]byte, error) {\n\tdata, err := os.ReadFile(path)\n\tif err != nil {\n\t\treturn nil, fmt.Errorf(\"load: %w\", err)\n\t}\n\treturn data, nil\n}",
		),
	}
}

var goPages = []string{
	"Code is typing with a purpose. Every bracket, quote and operator you have learned now has a job. Each lesson in this module shows one Go pattern, says when to use it, and asks you to type it exactly as written.",
	"Press enter at the end of each line, shown as ↵. The indentation on the next line is typed for you, as gofmt and your editor would do, so you only type the code itself.",
	"Go is read far more often than it is written. Short names for short scopes, early returns and small functions keep it easy to read, and the snippets here follow those habits.",
}

func code(number int, title, note string, snippets ...string) curriculum.Lesson {
	return curriculum.Lesson{
		Number: number, Title: title, Kind: curriculum.Code,
		Pages: []string{note}, Snippets: snippets, Keys: codeKeys(snippets), GoalWPM: codeGoal,
	}
}

func introducing(keys string, l curriculum.Lesson) curriculum.Lesson {
	l.NewKeys = []rune(keys)
	return l
}

func codeKeys(snippets []string) []rune {
	var keys []rune
	add := func(r rune) {
		if !strings.ContainsRune(string(keys), r) {
			keys = append(keys, r)
		}
	}
	for _, r := range curriculum.Typed(strings.Join(snippets, "\n")) {
		add(r)
		if key, ok := keyboard.Lookup(r); ok && key.Shifted {
			add(key.Base)
		}
	}
	return keys
}
