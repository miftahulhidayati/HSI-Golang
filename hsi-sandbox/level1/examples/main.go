package main

import "fmt"

func main() {
    // Basic syntax and data types in Go

    // Variables
    var name string = "John Doe"
    var age int = 30
    var height float64 = 5.9
    var isStudent bool = false

    // Print variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)

    // Constants
    const pi = 3.14
    fmt.Println("Value of Pi:", pi)

    // Basic control structure: if-else
    if age < 18 {
        fmt.Println(name, "is a minor.")
    } else {
        fmt.Println(name, "is an adult.")
    }

    // Basic control structure: for loop
    fmt.Println("Counting to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}