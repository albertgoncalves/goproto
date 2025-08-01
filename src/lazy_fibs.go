package main

import (
    "fmt"
    "os"
)

/* */

type lazy[T any] func() T

func newLazy[T any](f func() T) *lazy[T] {
    var self lazy[T]
    self = func() T {
        x := f()
        self = func() T {
            return x
        }
        return x
    }
    return &self
}

func (self lazy[T]) eval() T {
    return self()
}

/* */

type list[T any] struct {
    head T
    tail *lazy[list[T]]
}

func (self list[T]) evalTail() list[T] {
    return self.tail.eval()
}

func (self list[T]) drop(n int) list[T] {
    for range n {
        self = self.evalTail()
    }
    return self
}

func (self list[T]) zipWith(f func(T, T) T, other list[T]) list[T] {
    return list[T]{f(self.head, other.head), newLazy(
        func() list[T] {
            return self.evalTail().zipWith(f, other.evalTail())
        },
    )}
}

/* */

func main() {
    var fibs list[int]

    fibs = list[int]{0, newLazy(
        func() list[int] {
            return list[int]{1, newLazy(
                func() list[int] {
                    return fibs.zipWith(
                        func(a, b int) int {
                            return a + b
                        },
                        fibs.evalTail(),
                    )
                },
            )}
        },
    )}

    x := fibs.drop(50).head
    if x != 12586269025 {
        os.Exit(1)
    }
    fmt.Println(x)
}
