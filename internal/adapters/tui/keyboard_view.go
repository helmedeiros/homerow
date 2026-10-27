package tui

import (
	"strings"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

var rowIndent = []int{0, 2, 3, 5}

const spaceBarIndent = 12

func renderKeyboard(press rune) string {
	var lines []string
	for i, row := range keyboard.Rows() {
		var b strings.Builder
		b.WriteString(strings.Repeat(" ", rowIndent[i]))
		for _, k := range row {
			b.WriteString(renderKey(string(k), k == press))
		}
		lines = append(lines, b.String())
	}
	spaceBar := strings.Repeat(" ", spaceBarIndent) + renderKey("        space        ", press == ' ')
	return strings.Join(append(lines, spaceBar), "\n")
}

func renderKey(label string, pressed bool) string {
	if pressed {
		return pressStyle.Render("‹" + label + "›")
	}
	return keyStyle.Render("[" + label + "]")
}
