//go:build run

package main

import (
	"os"

	"golang.org/x/term"

	"github.com/hymkor/go-windows1x-virtualterminal"

	"github.com/hymkor/go-cursorposition"
)

func main() {
	if closer, err := virtualterminal.EnableStderr(); err != nil {
		panic(err.Error())
	} else {
		defer closer()
	}
	if oldState, err := term.MakeRaw(int(os.Stdin.Fd())); err != nil {
		panic(err.Error())
	} else {
		defer term.Restore(int(os.Stdin.Fd()), oldState)
	}
	w, err := cursorposition.AmbiguousWidth(os.Stderr)
	if err != nil {
		println(err.Error())
	} else {
		println(w)
	}
}
