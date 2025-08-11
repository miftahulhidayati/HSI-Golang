package main

import "fmt"

var hitung int

func tambah() {
	hitung++
}

func main() {
	for i := 0; i < 100; i++ {
		go tambah() // Menjalankan fungsi tambah secara goroutine
		// 1 : 2 go routine berjalan
		// 2 : 3 go routine berjalan
		// 3 : 2 go routine berjalan
		// membuat nilai hitungnya tidak konsisten
	}

	fmt.Println("Hasil hitung:", hitung) // Hasil mungkin tidak konsisten karena data race
}