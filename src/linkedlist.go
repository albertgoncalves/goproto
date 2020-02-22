package main

import (
    "fmt"
)

type node struct {
    value uint8
    next  *node
}

type linkedList struct {
    head *node
}

func (l *linkedList) push(v uint8) {
    currentNode := l.head
    nextNode := node{
        value: v,
        next:  currentNode,
    }
    l.head = &nextNode
}

func (l *linkedList) pop() (uint8, error) {
    if l.head == nil {
        return 0, fmt.Errorf("(%v).pop()", l)
    }
    currentNode := l.head
    value := currentNode.value
    l.head = currentNode.next
    return value, nil
}

func (l *linkedList) popAt(n uint) (uint8, error) {
    var prevNode *node
    currentNode := l.head
    for i := uint(0); i < n; i++ {
        if currentNode == nil || currentNode.next == nil {
            return 0, fmt.Errorf("(%v).popAt(%d)", l, n)
        }
        prevNode = currentNode
        currentNode = currentNode.next
    }
    if prevNode == nil {
        l.head = currentNode.next
    } else {
        prevNode.next = currentNode.next
    }
    return currentNode.value, nil
}

func (l linkedList) print() {
    var currentNode *node
    currentNode = l.head
    fmt.Print("[")
    for currentNode != nil {
        fmt.Print(" ")
        fmt.Print(currentNode.value)
        currentNode = currentNode.next
    }
    fmt.Println("]")
}

func main() {
    l := linkedList{
        head: nil,
    }
    for i := uint8(0); i < 10; i++ {
        l.push(i)
    }
    l.print()
    fmt.Println(l.popAt(3))
    fmt.Println(l.popAt(8))
    fmt.Println(l.popAt(0))
    fmt.Println(l.popAt(7))
    l.print()
    for err := error(nil); err == nil; _, err = l.pop() {
        if l.head != nil {
            fmt.Println(l.head.value)
        }
    }
    l.print()
    for i := uint8(11); i < 13; i++ {
        l.push(i)
    }
    l.print()
}
