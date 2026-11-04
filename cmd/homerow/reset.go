package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/helmedeiros/homerow/internal/app"
)

func confirmReset(trainer *app.Trainer, in io.Reader, out io.Writer) error {
	_, _ = fmt.Fprint(out, "Erase all progress? [y/N] ")
	answer, _ := bufio.NewReader(in).ReadString('\n')
	switch strings.ToLower(strings.TrimSpace(answer)) {
	case "y", "yes":
		if err := trainer.Reset(); err != nil {
			return err
		}
		_, _ = fmt.Fprintln(out, "progress erased")
	default:
		_, _ = fmt.Fprintln(out, "nothing changed")
	}
	return nil
}
