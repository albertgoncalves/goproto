package main

import (
    "fmt"
)

type fifo[T any] struct {
    in  []T
    out []T
}

func (self *fifo[T]) push(x T) {
    self.in = append(self.in, x)
}

func (self *fifo[T]) pop() T {
    if len(self.out) == 0 {
        n := len(self.in)
        for i := range n {
            self.out = append(self.out, self.in[n-(i+1)])
        }
        self.in = self.in[:0]
    }

    n := len(self.out)
    x := self.out[n-1]
    self.out = self.out[:n-1]
    return x
}

func main() {
    fifo := fifo[int]{}

    for i := range 10 {
        fifo.push(i)
    }
    fmt.Println(fifo)

    for range 5 {
        fmt.Println(fifo.pop())
    }
    fmt.Println(fifo)

    for i := range 5 {
        fifo.push(i)
    }
    fmt.Println(fifo)

    for range 10 {
        fmt.Println(fifo.pop())
    }
    fmt.Println(fifo)
}
