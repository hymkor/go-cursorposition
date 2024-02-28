go-cursorposition
=================

Request()
---------

Query and display the cursor position with ESC[6n

```example1.go
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

    // Query and display the cursor position with ESC[6n
    row, col, err := cursorposition.Request(os.Stderr)
    if err != nil {
        println(err.Error())
    } else {
        println(row, col)
    }
}
```

AmbiguousWidth()
----------------

Measure how far the cursor moves while the `▽` is printed

```example2.go
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
```

AmbiguousWidthGoTty
-------------------

`AmbiguousWidthGoTty` works same as `AmbiguousWidth`, but it switches the terminal raw-mode with the instance of "[github.com/mattn/go-tty][go-tty]".TTY

[go-tty]: https://github.com/mattn/go-tty

```example3.go
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
```

Referneces
----------
+ [端末の文字幅問題の傾向と対策 | IIJ Engineers Blog](https://eng-blog.iij.ad.jp/archives/12576)
+ [ANSI Escape Codes](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)
