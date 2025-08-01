package main

import (
    "cmp"
    "fmt"
)

type tree[K cmp.Ordered, V any] interface {
    getWeight() int

    mergeNode(node[K, V]) node[K, V]
    merge(tree[K, V]) tree[K, V]

    insert(K, V) tree[K, V]

    pop() (*V, tree[K, V])
}

type leaf[K cmp.Ordered, V any] struct{}

type node[K cmp.Ordered, V any] struct {
    weight int
    key    K
    value  V
    left   tree[K, V]
    right  tree[K, V]
}

func (self leaf[K, V]) getWeight() int {
    return 0
}

func (self node[K, V]) getWeight() int {
    return self.weight
}

func (self leaf[K, V]) mergeNode(other node[K, V]) node[K, V] {
    return other
}

func (self node[K, V]) mergeNode(other node[K, V]) node[K, V] {
    if other.key < self.key {
        return other.mergeNode(self)
    }

    right := self.right.mergeNode(other)

    if self.left.getWeight() < right.weight {
        self.weight = self.left.getWeight() + 1
        self.right = self.left
        self.left = right
    } else {
        self.weight = right.weight + 1
        self.right = right
    }

    return self
}

func (self leaf[K, V]) merge(other tree[K, V]) tree[K, V] {
    return other
}

func (self node[K, V]) merge(other tree[K, V]) tree[K, V] {
    return other.mergeNode(self)
}

func (self leaf[K, V]) insert(key K, value V) tree[K, V] {
    return node[K, V]{1, key, value, leaf[K, V]{}, leaf[K, V]{}}
}

func (self node[K, V]) insert(key K, value V) tree[K, V] {
    return self.mergeNode(node[K, V]{1, key, value, leaf[K, V]{}, leaf[K, V]{}})
}

func (self leaf[K, V]) pop() (*V, tree[K, V]) {
    return nil, leaf[K, V]{}
}

func (self node[K, V]) pop() (*V, tree[K, V]) {
    return &self.value, self.left.merge(self.right)
}

func insert[K cmp.Ordered, V any](self *tree[K, V], key K, value V) {
    *self = (*self).insert(key, value)
}

func pop[K cmp.Ordered, V any](self *tree[K, V]) *V {
    value, other := (*self).pop()
    *self = other
    return value
}

func main() {
    var tree tree[int, string] = &leaf[int, string]{}

    insert(&tree, 5, "a")
    insert(&tree, 3, "b")
    insert(&tree, 8, "c")
    insert(&tree, 1, "d")
    insert(&tree, -3, "e")
    insert(&tree, 2, "f")

    for {
        value := pop(&tree)
        if value == nil {
            break
        }
        fmt.Print(*value)
    }
    fmt.Println()
}
