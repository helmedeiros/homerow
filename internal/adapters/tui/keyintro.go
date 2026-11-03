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
	fmt.Fprintf(&b, "Type the %s using your %s finger.\n\n", keyName(k.key), info.Finger)
	b.WriteString(renderKeyboard(k.key) + "\n\n")
	b.WriteString(renderHands(info.Finger, true) + "\n\n")
	if k.wrong != "" {
		b.WriteString(wrongTextStyle.Render(fmt.Sprintf("That was %s. Reach with your %s finger.", k.wrong, info.Finger)) + "\n")
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
