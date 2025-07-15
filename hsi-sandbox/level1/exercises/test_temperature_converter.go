package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to simulate temperature conversion
func convertTemperature(input string) {
	fmt.Println("--- Konverter Suhu ---")
	fmt.Printf("Masukkan suhu dalam Celsius: %s\n", input)

	// Remove newline character and trim spaces
	input = strings.TrimSpace(input)

	// Try to convert input to float64
	celsius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		// If conversion fails, show error message
		fmt.Println("> Input Tidak Valid, hanya menerima angka")
		return
	}

	// Convert Celsius to Reamur: R = C × 4/5
	reamur := celsius * 4 / 5

	// Convert Celsius to Fahrenheit: F = (C × 9/5) + 32
	fahrenheit := (celsius * 9 / 5) + 32

	// Display results
	fmt.Printf("> Suhu dalam Reamur = %.0f\n", reamur)
	fmt.Printf("> Suhu dalam Fahrenheit = %.0f\n", fahrenheit)
}

func main() {
	fmt.Println("=== Test Aplikasi Konverter Suhu ===\n")

	// Test 1: Valid input (25)
	fmt.Println("Test 1: Input angka valid")
	convertTemperature("25")

	fmt.Println("\n" + strings.Repeat("-", 40) + "\n")

	// Test 2: Invalid input (string)
	fmt.Println("Test 2: Input tidak valid (string)")
	convertTemperature("dua puluh lima")

	fmt.Println("\n" + strings.Repeat("-", 40) + "\n")

	// Test 3: Valid input with decimal
	fmt.Println("Test 3: Input angka desimal")
	convertTemperature("32.5")

	fmt.Println("\n" + strings.Repeat("-", 40) + "\n")

	// Test 4: Invalid input (mixed)
	fmt.Println("Test 4: Input tidak valid (campuran)")
	convertTemperature("25abc")
}