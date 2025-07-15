package main

import (
    "fmt"
    "errors"
)

// Exercise 1: Create a function that takes two integers and returns their sum.
func add(a int, b int) int {
    return a + b
}

// Exercise 2: Create a function that divides two integers and returns an error if the divisor is zero.
func divide(a int, b int) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return float64(a) / float64(b), nil
}

// Exercise 3: Create a function that takes a string and returns its length.
func stringLength(s string) int {
    return len(s)
}

func main() {
    // Test the add function
    sum := add(5, 3)
    fmt.Println("Sum:", sum)

    // Test the divide function
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division Result:", result)
    }

    // Test the stringLength function
    length := stringLength("Hello, Golang!")
    fmt.Println("String Length:", length)
}