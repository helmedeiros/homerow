package main

import (
	"fmt"
	"io"
	"os"
)

var version = "dev"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}

func run(args []string, out io.Writer) int {
	if len(args) == 1 && args[0] == "version" {
		_, _ = fmt.Fprintf(out, "homerow %s\n", version)
		return 0
	}
	_, _ = fmt.Fprintln(out, "usage: homerow version")
	return 2
}
