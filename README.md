go-cursorposition
=================

Request()
---------

```example1.go
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

```example2.go
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
```

Referneces
----------
+ [端末の文字幅問題の傾向と対策 | IIJ Engineers Blog](https://eng-blog.iij.ad.jp/archives/12576)
+ [ANSI Escape Codes](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)
