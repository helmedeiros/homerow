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

func renderHands(key keyboard.Key, show bool) string {
	active := map[keyboard.Finger]bool{}
	caption := ""
	if show {
		active[key.Finger] = true
		caption = "  " + accentStyle.Render(key.Finger.String())
		if shift, ok := key.ShiftFinger(); ok {
			active[shift] = true
			caption = "  " + accentStyle.Render(key.Finger.String()+" + "+shift.String()+" on shift")
		}
	}
	tips, bases, labels, markers := blankRow(), blankRow(), blankRow(), blankRow()
	for finger, column := range fingerColumns {
		style := fingerStyle(finger, active[finger])
		if active[finger] {
			markers[column] = accentStyle.Render("▲")
		}
		if finger == keyboard.LeftThumb || finger == keyboard.RightThumb {
			bases[column] = style.Render("▄")
			labels[column] = mutedStyle.Render("⎵")
			continue
		}
		tips[column] = style.Render("▄")
		bases[column] = style.Render("█")
		labels[column] = mutedStyle.Render(string(keyboard.HomeKey(finger)))
	}
	return strings.Join([]string{join(tips), join(bases), join(labels), join(markers), caption}, "\n")
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
