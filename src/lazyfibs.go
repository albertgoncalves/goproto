package main

import (
    "fmt"
    "os"
)

type lazy[T any] struct {
    thunk func() *T
}

func lazyNew[T any](f func() *T) *lazy[T] {
    var l lazy[T]
    l.thunk = func() *T {
        x := f()
        l.thunk = func() *T {
            return x
        }
        return x
    }
    return &l
}

func (l *lazy[T]) force() *T {
    return l.thunk()
}

type list[T any] struct {
    value T
    next  *lazy[list[T]]
}

func (l *list[T]) head() T {
    return l.value
}

func (l *list[T]) tail() *list[T] {
    return l.next.force()
}

func (l *list[T]) drop(n int) *list[T] {
    for ; n != 0; n-- {
        l = l.tail()
    }
    return l
}

func zipWith[T any](f func(T, T) T, a *list[T], b *list[T]) *list[T] {
    return &list[T]{f(a.head(), b.head()), lazyNew(
        func() *list[T] {
            return zipWith(f, a.tail(), b.tail())
        },
    )}
}

func main() {
    var fibs list[int]
    fibs = list[int]{0, lazyNew(
        func() *list[int] {
            return &list[int]{1, lazyNew(
                func() *list[int] {
                    return zipWith(
                        func(a, b int) int {
                            return a + b
                        },
                        &fibs,
                        (&fibs).tail(),
                    )
                },
            )}
        },
    )}
    x := fibs.drop(50).head()
    if x != 12586269025 {
        os.Exit(1)
    }
    fmt.Println(x)
}
