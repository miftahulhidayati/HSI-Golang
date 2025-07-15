package main

import (
    "fmt"
)

// Exercise 1: Create a package named "mathutils" that contains a function to calculate the square of a number.
func main() {
    result := Square(5)
    fmt.Printf("The square of 5 is: %d\n", result)
}

// Square calculates the square of a given integer.
func Square(n int) int {
    return n * n
}

// Exercise 2: Create a package named "stringutils" that contains a function to reverse a string.
func main() {
    reversed := Reverse("golang")
    fmt.Printf("The reverse of 'golang' is: %s\n", reversed)
}

// Reverse returns the reversed version of the input string.
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// Exercise 3: Create a package named "converter" that contains a function to convert Celsius to Fahrenheit.
func main() {
    fahrenheit := CelsiusToFahrenheit(25)
    fmt.Printf("25 degrees Celsius is %.2f degrees Fahrenheit\n", fahrenheit)
}

// CelsiusToFahrenheit converts Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
    return (c * 9 / 5) + 32
}