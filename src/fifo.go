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

type fifoQueue struct {
    first *node
    last  *node
}

func (queue *fifoQueue) push(value t) {
    nextNode := &node{
        value: value,
        next:  nil,
    }
    if (queue.first != nil) && (queue.last == nil) {
        queue.first.next = nextNode
        queue.last = nextNode
    } else if queue.last != nil {
        queue.last.next = nextNode
        queue.last = nextNode
    } else {
        queue.first = nextNode
    }
}

func (queue *fifoQueue) pop() (t, error) {
    if queue.first == nil {
        return DEFAULT, fmt.Errorf("(%v).pop()", queue)
    }
    value := queue.first.value
    queue.first = queue.first.next
    if queue.first == queue.last {
        queue.last = nil
    }
    return value, nil
}

func (queue fifoQueue) print() {
    var currentNode *node
    currentNode = queue.first
    fmt.Print("[")
    for currentNode != nil {
        fmt.Print(" ", currentNode.value)
        currentNode = currentNode.next
    }
    fmt.Println("]")
}

func main() {
    fmt.Println("Hello!")
    queue := fifoQueue{
        first: nil,
        last:  nil,
    }
    for i := t(0); i < 10; i++ {
        queue.push(i)
    }
    queue.print()
    for {
        value, err := queue.pop()
        if err != nil {
            break
        }
        fmt.Println(value)
    }
    queue.print()
    queue.push(10)
    queue.push(11)
    queue.push(12)
    queue.print()
}
