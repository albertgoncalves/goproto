package main

import (
    "fmt"
)

var DONE = make(chan string, 1)

func pingPong(in <-chan int, out chan<- int, message string) {
    for {
        n := <-in
        if n == 0 {
            DONE <- "Done!"
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
    fmt.Println(<-DONE)
}
