package main

import (
    "fmt"
)

type t = uint8

const DEFAULT t = 0

type node struct {
    value t
    next  *node
}

type linkedList struct {
    head *node
}

func (list *linkedList) push(value t) {
    currentNode := list.head
    nextNode := node{
        value: value,
        next:  currentNode,
    }
    list.head = &nextNode
}

func (list *linkedList) pop() (t, error) {
    if list.head == nil {
        return DEFAULT, fmt.Errorf("(%v).pop()", list)
    }
    currentNode := list.head
    value := currentNode.value
    list.head = currentNode.next
    return value, nil
}

func (list *linkedList) popAt(n uint) (t, error) {
    var prevNode *node
    currentNode := list.head
    for i := uint(0); i < n; i++ {
        if currentNode == nil || currentNode.next == nil {
            return DEFAULT, fmt.Errorf("(%v).popAt(%d)", list, n)
        }
        prevNode = currentNode
        currentNode = currentNode.next
    }
    if prevNode == nil {
        list.head = currentNode.next
    } else {
        prevNode.next = currentNode.next
    }
    return currentNode.value, nil
}

func (list linkedList) print() {
    var currentNode *node
    currentNode = list.head
    fmt.Print("[")
    for currentNode != nil {
        fmt.Print(" ", currentNode.value)
        currentNode = currentNode.next
    }
    fmt.Println("]")
}

func main() {
    list := linkedList{}
    for i := t(0); i < 10; i++ {
        list.push(i)
    }
    list.print()
    fmt.Println(list.popAt(3))
    fmt.Println(list.popAt(8))
    fmt.Println(list.popAt(0))
    fmt.Println(list.popAt(7))
    list.print()
    for {
        value, err := list.pop()
        if err != nil {
            break
        }
        fmt.Println(value)
    }
    list.print()
    for i := t(11); i < 13; i++ {
        list.push(i)
    }
    list.print()
}
