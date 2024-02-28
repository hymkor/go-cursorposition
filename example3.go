//go:build run

package main

import (
	"os"

	"github.com/mattn/go-tty"

	"github.com/hymkor/go-windows1x-virtualterminal"

	"github.com/hymkor/go-cursorposition"
)

func main() {
	// On Windows, enable ANSI ESCAPE SEQUENCE.
	// On other OSes, do nothing.
	if closer, err := virtualterminal.EnableStderr(); err != nil {
		panic(err.Error())
	} else {
		defer closer()
	}

	tty1, err := tty.Open()
	if err != nil {
		panic(err.Error())
	}
	tty1.Close()

	// Measure how far the cursor moves while the `▽` is printed
	w, err := cursorposition.AmbiguousWidthGoTty(tty1, os.Stderr)
	if err != nil {
		println(err.Error())
	} else {
		println(w)
	}
}
