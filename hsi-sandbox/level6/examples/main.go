package main

import (
    "fmt"
    "math/rand"
    "time"
)

// GenerateRandomNumber generates a random number between min and max
func GenerateRandomNumber(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min+1) + min
}

// main function to demonstrate package usage
func main() {
    min := 1
    max := 100
    randomNumber := GenerateRandomNumber(min, max)
    fmt.Printf("Generated random number between %d and %d: %d\n", min, max, randomNumber)
}