package main

import (
    "fmt"
)

type tree[T any] interface {
    getHeight() int
    insert(T) tree[T]
    print()
}

type leaf[T any] struct{}

type node[T any] struct {
    height int
    left   tree[T]
    value  T
    right  tree[T]
}

func (t leaf[T]) getHeight() int {
    return -1
}

func (t node[T]) getHeight() int {
    return t.height
}

func (t leaf[T]) insert(value T) tree[T] {
    return node[T]{0, leaf[T]{}, value, leaf[T]{}}
}

func (t node[T]) insert(value T) tree[T] {
    lh := t.left.getHeight()
    rh := t.right.getHeight()

    // TODO: Is this right? Seems a little weird.
    if lh < rh {
        t.left = t.left.insert(value)
    } else if lh > rh {
        t.right = t.right.insert(value)
    } else {
        t.left = t.left.insert(value)
        t.height = t.left.getHeight() + 1
    }

    return t
}

func (t leaf[T]) print() {
}

func (t node[T]) print() {
    t.left.print()
    fmt.Print(" ", t.value)
    t.right.print()
}

func main() {
    var t tree[int] = leaf[int]{}
    for i := range 15 {
        t = t.insert(i)
    }
    t.print()
    fmt.Println()
}
