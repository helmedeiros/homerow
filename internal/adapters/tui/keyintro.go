package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/helmedeiros/homerow/internal/domain/keyboard"
)

type keyIntroScreen struct {
	key   rune
	wrong string
}

func newKeyIntroScreen(key rune) keyIntroScreen {
	return keyIntroScreen{key: key}
}

func (k keyIntroScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return k, nil
	}
	switch key.String() {
	case "esc":
		return k, send(leaveLessonMsg{})
	case "ctrl+c":
		return k, tea.Quit
	}
	pressed := string(key.Runes)
	if key.Type == tea.KeySpace {
		pressed = " "
	}
	if strings.ContainsRune(pressed, k.key) {
		return k, send(stepDoneMsg{})
	}
	k.wrong = key.String()
	return k, nil
}

func (k keyIntroScreen) View() string {
	info, _ := keyboard.Lookup(k.key)
	var b strings.Builder
	b.WriteString(mutedStyle.Render("NEW KEY") + "\n\n")
	fmt.Fprintf(&b, "Type the %s using your %s finger.\n", keyName(k.key), info.Finger)
	hint, shifted := shiftHint(info)
	if shifted {
		b.WriteString(hint + "\n")
	}
	b.WriteString("\n")
	b.WriteString(renderKeyboard(k.key) + "\n\n")
	b.WriteString(renderHands(info, true) + "\n\n")
	if k.wrong != "" {
		b.WriteString(wrongTextStyle.Render(k.wrongHint(info, hint, shifted)) + "\n")
	}
	b.WriteString(mutedStyle.Render("esc lessons"))
	return b.String()
}

func keyName(k rune) string {
	if k == ' ' {
		return pressStyle.Render(" space bar ")
	}
	return pressStyle.Render(" "+string(k)+" ") + " key"
}

func (k keyIntroScreen) wrongHint(info keyboard.Key, hint string, shifted bool) string {
	if shifted && k.wrong == string(info.Base) {
		return fmt.Sprintf("That was %s. %s", k.wrong, hint)
	}
	return fmt.Sprintf("That was %s. Reach with your %s finger.", k.wrong, info.Finger)
}

func shiftHint(info keyboard.Key) (string, bool) {
	shift, ok := info.ShiftFinger()
	side := "left"
	if shift.Hand() == keyboard.Right {
		side = "right"
	}
	return fmt.Sprintf("Hold the %s shift key with your %s.", side, shift), ok
}
