package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func f(wg *sync.WaitGroup, j int) {
    defer wg.Done()
    t := time.Duration(rand.Int31n(2500)) * time.Millisecond
    fmt.Printf("  [ %2d ] sleeping for %v\n", j, t)
    time.Sleep(t)
    fmt.Printf("  [ %2d ] done\n", j)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go f(&wg, i)
    }
    wg.Wait()
}
