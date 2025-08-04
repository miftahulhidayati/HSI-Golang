package main

import "fmt"

// Array adalah tipe data yang digunakan untuk menyimpan sekumpulan nilai dengan tipe data yang sama.
// Array adalah struktur data yang memiliki panjang tetap yang ditentukan saat deklarasi.
func main() {
	// cara membuat array metode 1: gaperlu isi nilai diawal
	// index 0-9
	var arrayAngka [10]int
	// cara membuat array metode 2: langsung isi nilai
	arrayAngka1 := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array Angka 1:", arrayAngka1[0])

	// mengisi data array metode 1
	arrayAngka[0] = 10
	arrayAngka[1] = 20
	arrayAngka[2] = 30
	arrayAngka[3] = 40
	arrayAngka[4] = 50

	fmt.Println("Array Angka 1:", arrayAngka[4])

	// cara membuat array metode 3: panjang array otomatis ditentukan
	arrayAngka2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Panjang Array:", len(arrayAngka2))

}