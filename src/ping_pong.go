package main

import (
    "fmt"
    "time"
)

func pingPong(in <-chan int, out chan<- int, done chan<- bool, message string) {
    for {
        n := <-in
        if n == 0 {
            done <- true
            return
        }
        fmt.Println(message)

        time.Sleep(100 * time.Millisecond)
        out <- n - 1
    }
}

func main() {
    ping := make(chan int, 1)
    pong := make(chan int, 1)
    done := make(chan bool, 1)

    go pingPong(ping, pong, done, "ping")
    go pingPong(pong, ping, done, "pong")

    ping <- 5
    <-done

    fmt.Println("Done!")
}
