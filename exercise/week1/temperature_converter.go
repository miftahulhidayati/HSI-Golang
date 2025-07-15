package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Display header
	fmt.Println("--- Konverter Suhu ---")
	fmt.Print("Masukkan suhu dalam Celsius: ")

	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Read input until Enter is pressed
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

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