package main

import (
    "fmt"
    "time"
)

func main() {
    // Start a goroutine
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine:", i)
            time.Sleep(500 * time.Millisecond)
        }
    }()

    // Main goroutine work
    for i := 0; i < 5; i++ {
        fmt.Println("Main:", i)
        time.Sleep(700 * time.Millisecond)
    }

    // Wait for user input to prevent the program from exiting immediately
    var input string
    fmt.Scanln(&input)
}