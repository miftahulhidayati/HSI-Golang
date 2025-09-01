package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Suhu struct {
	Lokasi      string
	Celsius     float64
	Reamur      float64
	Fahrenheit  float64
	Klasifikasi string
}

func inputLokasi() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan lokasi pengukuran suhu : ")
	lokasi, _ := reader.ReadString('\n')
	lokasi = strings.TrimSpace(lokasi)
	// Validasi lokasi hanya string (tidak boleh angka)
	if _, err := strconv.Atoi(lokasi); err == nil {
		return ""
	}
	return lokasi
}

func inputSuhu() (float64, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan suhu dalam Celsius: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, false
	}
	return val, true
}

func konversiSuhu(s *Suhu) {
	s.Reamur = s.Celsius * 4 / 5
	s.Fahrenheit = s.Celsius*9/5 + 32
}

func klasifikasiSuhu(s *Suhu) {
	if s.Celsius < 18 {
		s.Klasifikasi = "dingin"
	} else if s.Celsius <= 25 {
		s.Klasifikasi = "hangat"
	} else {
		s.Klasifikasi = "panas"
	}
}

func main() {
	fmt.Println("--- Konverter Suhu ---")
	lokasi := inputLokasi()
	if lokasi == "" {
		fmt.Println("> Input Tidak Valid, hanya menerima string lokasi")
		return
	}

	suhuC, valid := inputSuhu()
	if !valid {
		fmt.Println("> Input Tidak Valid, hanya menerima angka")
		return
	}

	s := Suhu{
		Lokasi:  lokasi,
		Celsius: suhuC,
	}
	konversiSuhu(&s)
	klasifikasiSuhu(&s)

	fmt.Printf("> Suhu di %s adalah %s\n", s.Lokasi, s.Klasifikasi)
	fmt.Printf("> Suhu di %s dalam Reamur = %.2f\n", s.Lokasi, s.Reamur)
	fmt.Printf("> Suhu di %s dalam Fahrenheit = %.2f\n", s.Lokasi, s.Fahrenheit)
}
