// main.go
// Entry poinrt aplikasi
// spesial package dari bahasa go
package main

import (
	"fmt"
	"praktikum/calculator"

	"praktikum/stringgenerator"
)

func main() {
	// memanggil funct TambahAngka dari package calculator
	hasil := calculator.TambahAngka(5, 3)
	fmt.Println(hasil)

	hasilKurang := calculator.KurangAngka(5, 3)
	fmt.Println(hasilKurang)

	pesanBaru := stringgenerator.MencetakSalam("miftah")
	fmt.Println(pesanBaru)

}