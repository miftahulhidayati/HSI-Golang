package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        for i := 1; i <= 5; i++ {
            fmt.Printf("Goroutine 1: %d\n", i)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 1; i <= 5; i++ {
            fmt.Printf("Goroutine 2: %d\n", i)
        }
    }()

    wg.Wait()
    fmt.Println("All goroutines finished executing.")
}