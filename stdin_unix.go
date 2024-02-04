//go:build !windows

package cursorposition

import (
	"os"
)

func stdin() (*os.File, error) {
	return os.Open("/dev/tty")
}
