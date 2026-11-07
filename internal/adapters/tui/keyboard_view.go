package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

var rowIndent = []int{0, 2, 3, 5}

const (
	edgeWidth      = 8
	spaceBarIndent = 20
)

func renderKeyboard(press rune) string {
	key, known := keyboard.Lookup(press)
	shift, shifted := key.ShiftFinger()
	leftShift := renderKey("shift", shifted && shift == keyboard.LeftPinky)
	rightShift := renderKey("shift", shifted && shift == keyboard.RightPinky)
	edges := []string{"", renderKey("tab", known && press == '\t'), "", leftShift}
	var lines []string
	for i, row := range keyboard.Rows() {
		var b strings.Builder
		b.WriteString(edges[i] + strings.Repeat(" ", edgeWidth-lipgloss.Width(edges[i])+rowIndent[i]))
		for _, k := range row {
			b.WriteString(renderKey(string(k), known && k == key.Base))
		}
		if i == len(edges)-1 {
			b.WriteString(rightShift)
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
