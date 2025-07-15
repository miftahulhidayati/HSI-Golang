package main

import (
    "errors"
    "fmt"
)

// Function to divide two numbers with error handling
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Example of successful division
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result of 10 / 2 =", result)
    }

    // Example of division by zero
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result of 10 / 0 =", result)
    }
}