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
	leftShift := renderKey("shift", keyboard.LeftPinky, shifted && shift == keyboard.LeftPinky)
	rightShift := renderKey("shift", keyboard.RightPinky, shifted && shift == keyboard.RightPinky)
	edges := []string{
		renderKey("`", keyboard.LeftPinky, known && key.Base == '`'),
		renderKey("tab", keyboard.LeftPinky, known && press == '\t'),
		"", leftShift,
	}
	var lines []string
	for i, row := range keyboard.Rows() {
		var b strings.Builder
		pad := strings.Repeat(" ", edgeWidth-lipgloss.Width(edges[i])+rowIndent[i])
		if i == 0 {
			b.WriteString(pad + edges[i])
		} else {
			b.WriteString(edges[i] + pad)
		}
		for _, k := range row {
			b.WriteString(renderKey(string(k), fingerOf(k), known && k == key.Base))
		}
		if i == len(edges)-1 {
			b.WriteString(rightShift)
		}
		lines = append(lines, b.String())
	}
	spaceBar := strings.Repeat(" ", spaceBarIndent) + renderKey("        space        ", fingerOf(' '), press == ' ')
	return strings.Join(append(lines, spaceBar), "\n")
}

func renderKey(label string, finger keyboard.Finger, pressed bool) string {
	if pressed {
		return keyCapStyle(finger, true).Render("‹" + label + "›")
	}
	return keyCapStyle(finger, false).Render("[" + label + "]")
}

func keyCapStyle(finger keyboard.Finger, pressed bool) lipgloss.Style {
	if pressed {
		return lipgloss.NewStyle().Foreground(colorOnFinger).Background(fingerColors[finger]).Bold(true)
	}
	return lipgloss.NewStyle().Foreground(fingerColors[finger]).Faint(true)
}

func fingerOf(r rune) keyboard.Finger {
	key, _ := keyboard.Lookup(r)
	return key.Finger
}
