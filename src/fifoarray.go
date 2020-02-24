package main

import (
    "fmt"
)

type t = uint8

const DEFAULT t = 0
const CAPACITY uint = 10

var QUEUE = fifoQueue{
    memory:       [CAPACITY]t{},
    index:        0,
    remainingCap: CAPACITY,
}

type fifoQueue struct {
    memory       [CAPACITY]t
    index        uint
    remainingCap uint
}

func (queue *fifoQueue) push(value t) error {
    if queue.remainingCap == 0 {
        return fmt.Errorf("(%v).push()", queue)
    }
    queue.memory[queue.index] = value
    queue.index = (queue.index + 1) % CAPACITY
    queue.remainingCap--
    return nil
}

func (queue *fifoQueue) pop() (t, error) {
    if CAPACITY <= queue.remainingCap {
        return DEFAULT, fmt.Errorf("(%v).pop()", queue)
    }
    i := (queue.index + queue.remainingCap) % CAPACITY
    value := queue.memory[i]
    queue.memory[i] = DEFAULT
    queue.remainingCap++
    return value, nil
}

func (queue fifoQueue) print() {
    fmt.Print("[")
    for i := uint(0); i < CAPACITY; i++ {
        fmt.Print(" ", queue.memory[i])
    }
    fmt.Println("]")
}

func main() {
    var i t
    f := func() {
        for {
            if err := QUEUE.push(i); err != nil {
                break
            }
            i++
        }
    }
    f()
    QUEUE.print()
    for i := uint(0); i < 3; i++ {
        value, err := QUEUE.pop()
        if err != nil {
            break
        }
        fmt.Println(value)
    }
    QUEUE.print()
    f()
    QUEUE.print()
}
