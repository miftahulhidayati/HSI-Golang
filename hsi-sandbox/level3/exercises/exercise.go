package main

import "fmt"

// Exercise 1: Define a struct for a Book
type Book struct {
	Title  string
	Author string
	Year   int
}

// Exercise 2: Create a method to display book details
func (b Book) Display() {
	fmt.Printf("Title: %s, Author: %s, Year: %d\n", b.Title, b.Author, b.Year)
}

// Exercise 3: Create a slice of books and display their details
func main() {
	books := []Book{
		{"The Go Programming Language", "Alan A. A. Donovan", 2015},
		{"Learning Go", "Jon Bodner", 2020},
		{"Go in Action", "William Kennedy", 2015},
	}

	for _, book := range books {
		book.Display()
	}
}