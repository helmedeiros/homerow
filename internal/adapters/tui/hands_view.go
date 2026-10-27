package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

const handsWidth = 28

var fingerColumns = map[keyboard.Finger]int{
	keyboard.LeftPinky: 2, keyboard.LeftRing: 4, keyboard.LeftMiddle: 6, keyboard.LeftIndex: 8,
	keyboard.LeftThumb: 11, keyboard.RightThumb: 16,
	keyboard.RightIndex: 19, keyboard.RightMiddle: 21, keyboard.RightRing: 23, keyboard.RightPinky: 25,
}

var fingerColors = map[keyboard.Finger]lipgloss.Color{
	keyboard.LeftPinky: "#3B82F6", keyboard.LeftRing: "#84CC16", keyboard.LeftMiddle: "#EAB308", keyboard.LeftIndex: "#EF4444",
	keyboard.LeftThumb: "#94A3B8", keyboard.RightThumb: "#94A3B8",
	keyboard.RightIndex: "#EF4444", keyboard.RightMiddle: "#EAB308", keyboard.RightRing: "#84CC16", keyboard.RightPinky: "#3B82F6",
}

func renderHands(active keyboard.Finger, show bool) string {
	tips, bases, labels := blankRow(), blankRow(), blankRow()
	for finger, column := range fingerColumns {
		thumb := finger == keyboard.LeftThumb || finger == keyboard.RightThumb
		style := fingerStyle(finger, show && finger == active)
		if thumb {
			bases[column] = style.Render("▄")
			labels[column] = mutedStyle.Render("⎵")
			continue
		}
		tips[column] = style.Render("▄")
		bases[column] = style.Render("█")
		labels[column] = mutedStyle.Render(string(keyboard.HomeKey(finger)))
	}
	marker, caption := "", ""
	if show {
		marker = strings.Repeat(" ", fingerColumns[active]) + accentStyle.Render("▲")
		caption = "  " + accentStyle.Render(active.String())
	}
	return strings.Join([]string{join(tips), join(bases), join(labels), marker, caption}, "\n")
}

func fingerStyle(f keyboard.Finger, active bool) lipgloss.Style {
	if active {
		return lipgloss.NewStyle().Foreground(fingerColors[f]).Bold(true)
	}
	return mutedStyle
}

func blankRow() []string {
	row := make([]string, handsWidth)
	for i := range row {
		row[i] = " "
	}
	return row
}

func join(cells []string) string {
	return strings.TrimRight(strings.Join(cells, ""), " ")
}
