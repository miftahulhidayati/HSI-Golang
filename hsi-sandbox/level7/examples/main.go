package main

import (
    "fmt"
    "testing"
)

// Function to add two integers
func Add(a int, b int) int {
    return a + b
}

// Test for the Add function
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Function to subtract two integers
func Subtract(a int, b int) int {
    return a - b
}

// Test for the Subtract function
func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    expected := 2
    if result != expected {
        t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
    }
}

// Main function to run tests
func main() {
    fmt.Println("Running tests...")
    testing.Main(func(pat, str string) (bool, error) { return true, nil }, []testing.InternalTest{
        {"TestAdd", TestAdd},
        {"TestSubtract", TestSubtract},
    }, nil, nil)
}