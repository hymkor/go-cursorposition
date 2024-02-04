package cursorposition

import (
	"os"

	"golang.org/x/sys/windows"
)

func stdin() (*os.File, error) {
	conin := []uint16{'C', 'O', 'N', 'I', 'N', '$', 0}
	in, err := windows.CreateFile(
		&conin[0],
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		uint32(windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE),
		nil, windows.OPEN_EXISTING, windows.FILE_ATTRIBUTE_NORMAL, 0)

	if err != nil {
		return nil, err
	}

	return os.NewFile(uintptr(in), "/dev/tty"), nil
}
