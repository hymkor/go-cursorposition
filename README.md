go-cursorposition
=================

Request()
---------

```example1.go
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
    row, col, err := cursorposition.Request()
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
```

Referneces
----------
+ [端末の文字幅問題の傾向と対策 | IIJ Engineers Blog](https://eng-blog.iij.ad.jp/archives/12576)
+ [ANSI Escape Codes](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)
