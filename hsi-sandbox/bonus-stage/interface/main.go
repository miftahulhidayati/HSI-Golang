package main

import "fmt"

// mendefinisikan interface/kontrak/method yang harus dimiliki oleh struct
// interface ini digunakan untuk validasi usia pada form pendaftaran
// konsep interface itu polimorfisme, satu nama method mempunyai banyak implementasi
// namun memiliki method yang sama, sehingga bisa digunakan secara umum

// kalau lebih dari satu method, semua method harus diimplementasikan
type FormPendaftaranInterface interface {
	ValidasiUsia(usia int) bool
	ValidasiGender(jenisKelamin string) bool

}

type FormPendaftaran struct {
	NamaLengkap  string
	Email        string
	JenisKelamin string
	Usia         int
}

type FormPendaftareanUsiaSenja struct {
	FormPendaftaran
	Penyakit string
}

func (f FormPendaftaran) ValidasiUsia(usia int) bool {
	if usia < 15 || usia > 75 {
		return false
	}
	return true
}
func (f FormPendaftareanUsiaSenja) ValidasiUsia(usia int) bool {
	if usia > 75 {
		return true
	}
	return false
}
func (f FormPendaftaran) ValidasiGender(jenisKelamin string) bool {
	if jenisKelamin == "Laki-laki" || jenisKelamin == "Perempuan" {
		return true
	}
	return false
}
func (f FormPendaftareanUsiaSenja) ValidasiGender(jenisKelamin string) bool {
	if jenisKelamin == "Laki-laki" || jenisKelamin == "Perempuan" {
		return true
	}
	return false
}
func ValidasiUsiaForm(formInt FormPendaftaranInterface, usia int) bool {
	return formInt.ValidasiUsia(usia)
}
func main() {
	pendaftaran1 := FormPendaftaran{
		NamaLengkap:  "Budi Santoso",
		Email:"email@email.com",
		Usia:709,
		JenisKelamin: "Laki-laki",
	}
	pendaftaran2 := FormPendaftareanUsiaSenja{
	FormPendaftaran: FormPendaftaran{
			NamaLengkap:  "Siti Aminah",
			Email: "siti@gmail.com",
			Usia: 80,
			JenisKelamin: "Perempuan",
		},
		Penyakit: "Diabetes",
	}
	apakahUsiaValid := ValidasiUsiaForm(pendaftaran1, pendaftaran1.Usia)
	fmt.Println("Apakah usia valid?", apakahUsiaValid)
	apakahJenisKelaminValid := pendaftaran1.ValidasiGender(pendaftaran1.JenisKelamin)
	fmt.Println("Apakah jenis kelamin valid?", apakahJenisKelaminValid)

	apakahUsiaValid2 := ValidasiUsiaForm(pendaftaran2, pendaftaran2.Usia)
	fmt.Println("Apakah usia senja valid?", apakahUsiaValid2)
	apakahJenisKelaminValid2 := pendaftaran2.ValidasiGender(pendaftaran2.JenisKelamin)
	fmt.Println("Apakah jenis kelamin senja valid?", apakahJenisKelaminValid2)
}