package main

import (
    "fmt"
    "time"
)

const (
    N = 5
    M = 1
)

var (
    FORKS [N]chan bool
    DONE  [N]chan bool
)

func dine(left <-chan bool, right <-chan bool) {
    fmt.Printf("Waiting for fork [%v]\n", left)
    <-left

    // NOTE: Create opportunity for deadlock.
    time.Sleep(50 * time.Millisecond)

    fmt.Printf("Waiting for fork [%v]\n", right)
    <-right

    fmt.Println("Eating")
    time.Sleep(500 * time.Millisecond)
}

func think(left chan<- bool, right chan<- bool) {
    fmt.Printf("Returning fork [%v]\n", left)
    left <- true

    fmt.Printf("Returning fork [%v]\n", right)
    right <- true

    fmt.Println("Thinking")
    time.Sleep(500 * time.Millisecond)
}

func loop(left chan bool, right chan bool, done chan<- bool) {
    for i := 0; i < M; i++ {
        dine(left, right)
        think(left, right)
    }
    done <- true
}

func main() {
    for i := 0; i < N; i++ {
        FORKS[i] = make(chan bool, 1)
        FORKS[i] <- true
        DONE[i] = make(chan bool, 1)
    }

    // NOTE: This will cause a deadlock.
    // for i := 0; i < N; i++ {
    //     go loop(FORKS[i], FORKS[(i+1)%N], DONE[i])
    // }

    go loop(FORKS[1], FORKS[0], DONE[0])
    for i := 1; i < N; i++ {
        go loop(FORKS[i], FORKS[(i+1)%N], DONE[i])
    }

    for i := 0; i < N; i++ {
        <-DONE[i]
    }
    fmt.Println("Done!")
}
