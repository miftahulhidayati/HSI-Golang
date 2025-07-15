package exercises

import (
	"fmt"
	"testing"
)

// Sample function to demonstrate testing
func Add(a, b int) int {
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

// Exercise: Write a function that subtracts two integers and test it
func Subtract(a, b int) int {
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

// Exercise: Create a function that multiplies two integers and write a test for it
func Multiply(a, b int) int {
	return a * b
}

// Test for the Multiply function
func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20
	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; want %d", result, expected)
	}
}

// Exercise: Create a function that divides two integers and write a test for it
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Test for the Divide function
func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	expected := 5
	if err != nil || result != expected {
		t.Errorf("Divide(10, 2) = %d, %v; want %d, nil", result, err, expected)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Errorf("Divide(10, 0) should return an error")
	}
}