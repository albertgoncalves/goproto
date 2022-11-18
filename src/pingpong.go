package main

import (
    "fmt"
)

var DONE = make(chan bool, 1)

func pingPong(in <-chan int, out chan<- int, message string) {
    for {
        n := <-in
        if n == 0 {
            DONE <- true
            return
        }
        fmt.Println(message)
        out <- n - 1
    }
}

func main() {
    ping := make(chan int, 1)
    pong := make(chan int, 1)
    go pingPong(ping, pong, "ping")
    go pingPong(pong, ping, "pong")
    ping <- 5
    <-DONE
    fmt.Println("Done!")
}
