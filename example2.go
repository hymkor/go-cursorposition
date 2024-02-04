//go:build run

package main

import (
	"os"

	"golang.org/x/term"

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

	// Switch terminal to raw-mode.
	if oldState, err := term.MakeRaw(int(os.Stdin.Fd())); err != nil {
		panic(err.Error())
	} else {
		defer term.Restore(int(os.Stdin.Fd()), oldState)
	}

	// Measure how far the cursor moves while the `▽` is printed
	w, err := cursorposition.AmbiguousWidth(os.Stderr)
	if err != nil {
		println(err.Error())
	} else {
		println(w)
	}
}
