package cursorposition

import (
	"errors"
	"io"
	"os"
	"regexp"

	"golang.org/x/term"
)

func gets(in io.Reader) ([]byte, error) {
	var line [256]byte

	n, err := in.Read(line[:])
	return line[:n], err
}

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

var rxPattern = regexp.MustCompile(`\x1B\[(\d+)\;(\d+)R`)

func btoi(b []byte) int {
	n := 0
	for len(b) > 0 {
		n = n*10 + int(b[0]-'0')
		b = b[1:]
	}
	return n
}

var ErrAnsiEscapeSequenceNotSupported = errors.New("ANSI Escape sequence not supported")

func Request() (int, int, error) {
	io.WriteString(os.Stderr, "\x1B[6n")

	if oldState, err := term.MakeRaw(int(os.Stdin.Fd())); err != nil {
		return 0, 0, err
	} else {
		defer term.Restore(int(os.Stdin.Fd()), oldState)
	}

	var err error
	for err == nil {
		var s []byte
		s, err = gets(os.Stdin)

		for len(s) > 0 && s[0] != '\x1B' {
			s = s[1:]
		}
		if len(s) <= 2 || s[1] != '[' {
			return 0, 0, ErrAnsiEscapeSequenceNotSupported
		}
		m := rxPattern.FindSubmatch(s)
		if m == nil {
			return 0, 0, ErrAnsiEscapeSequenceNotSupported
		}
		row := btoi(m[1])
		col := btoi(m[2])
		return row, col, nil
	}
	return 0, 0, err
}

func AmbiguousWidth() (int, error) {
	io.WriteString(os.Stderr, "\r\u25BD")
	_, w, err := Request()
	io.WriteString(os.Stderr, "\r\x1B[2K")
	return w - 1, err
}
