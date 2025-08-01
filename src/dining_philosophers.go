package main

import (
    "fmt"
    "time"
)

func dine(left <-chan bool, right <-chan bool) {
    fmt.Println("Waiting for fork [", left, "]")
    <-left

    // NOTE: Create opportunity for deadlock.
    time.Sleep(50 * time.Millisecond)

    fmt.Println("Waiting for fork [", right, "]")
    <-right

    fmt.Println("Eating")
    time.Sleep(500 * time.Millisecond)
}

func think(left chan<- bool, right chan<- bool) {
    fmt.Println("Returning fork [", left, "]")
    left <- true

    fmt.Println("Returning fork [", right, "]")
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

const (
    N = 5
    M = 1
)

func main() {
    var (
        forks [N]chan bool
        done  [N]chan bool
    )

    for i := 0; i < N; i++ {
        forks[i] = make(chan bool, 1)
        forks[i] <- true
        done[i] = make(chan bool, 1)
    }

    // NOTE: This will cause a deadlock.
    /*
       for i := 0; i < N; i++ {
           go loop(forks[i], forks[(i+1)%N], done[i])
       }
    */

    go loop(forks[1], forks[0], done[0])
    for i := 1; i < N; i++ {
        go loop(forks[i], forks[(i+1)%N], done[i])
    }

    for i := 0; i < N; i++ {
        <-done[i]
    }
    fmt.Println("Done!")
}
