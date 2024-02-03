//go:build run

package main

import (
	"github.com/hymkor/go-cursorposition"
	"github.com/hymkor/go-windows1x-virtualterminal"
)

func main() {
	if closer, err := virtualterminal.EnableStderr(); err != nil {
		panic(err.Error())
	} else {
		defer closer()
	}
	w, err := cursorposition.AmbiguousWidth()
	if err != nil {
		println(err.Error())
	} else {
		println(w)
	}
}