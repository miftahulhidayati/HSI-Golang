package main

import "fmt"

func main() {
    // Exercise 1: Declare a variable of type string and assign it a value.
    var greeting string = "Hello, Golang!"
    fmt.Println(greeting)

    // Exercise 2: Create a constant for the value of Pi.
    const Pi = 3.14
    fmt.Println("Value of Pi:", Pi)

    // Exercise 3: Use a basic if statement to check if a number is positive.
    number := 10
    if number > 0 {
        fmt.Println("The number is positive.")
    } else {
        fmt.Println("The number is not positive.")
    }

    // Exercise 4: Create a simple for loop that prints numbers from 1 to 5.
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}