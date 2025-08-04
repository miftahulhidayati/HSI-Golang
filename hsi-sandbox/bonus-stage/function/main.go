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

func kalkulasiPembagian(pembilang, pembagi int) (int, int, error) {
	if pembagi == 0 {
		return 0, 0, errors.New("Pembagi tidak boleh nol")
	}
	hasil := pembilang / pembagi
	sisa := pembilang % pembagi
	return hasil, sisa, nil
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
		fmt.Println("Pesan Error:", err)
	} else {
		fmt.Println("Result of 10 / 0 =", result)
	}

	hasil, sisa, err := kalkulasiPembagian(10, 6)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil pembagian 10 / 6 = %d, sisa = %d\n", hasil, sisa)
	}
}
