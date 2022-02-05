package main

import (
    "fmt"
    "os"
)

type lazyListInt struct {
    cached bool
    result *listInt
    thunk  func() *listInt
}

func lazyNew(f func() *listInt) *lazyListInt {
    return &lazyListInt{false, nil, f}
}

func (x *lazyListInt) force() *listInt {
    if !x.cached {
        x.result = x.thunk()
        x.cached = true
    }
    return x.result
}

type listInt struct {
    value int
    next  *lazyListInt
}

func (l *listInt) head() int {
    return l.value
}

func (l *listInt) tail() *listInt {
    return l.next.force()
}

func (l *listInt) drop(n int) *listInt {
    for ; n != 0; n-- {
        l = l.tail()
    }
    return l
}

func zipSum(a *listInt, b *listInt) *listInt {
    return &listInt{a.head() + b.head(), lazyNew(
        func() *listInt {
            return zipSum(a.tail(), b.tail())
        },
    )}
}

func f1(l *listInt) func() *listInt {
    return func() *listInt {
        return &listInt{1, lazyNew(f0(l))}
    }
}

func f0(l *listInt) func() *listInt {
    return func() *listInt {
        return zipSum(l, l.tail())
    }
}

func main() {
    var fibs listInt
    fibs = listInt{0, lazyNew(f1(&fibs))}
    x := fibs.drop(50).head()
    if x != 12586269025 {
        os.Exit(1)
    }
    fmt.Println(x)
}
