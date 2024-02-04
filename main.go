package cursorposition

import (
	"errors"
	"io"
	"os"
	"regexp"
)

func gets(in io.Reader) ([]byte, error) {
	var line [256]byte

	n, err := in.Read(line[:])
	return line[:n], err
}

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

func Request(w io.Writer) (int, int, error) {
	io.WriteString(w, "\x1B[6n")
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

func AmbiguousWidth(w io.Writer) (int, error) {
	io.WriteString(w, "\r\u25BD")
	_, col, err := Request(w)
	io.WriteString(w, "\r\x1B[2K")
	return col - 1, err
}
