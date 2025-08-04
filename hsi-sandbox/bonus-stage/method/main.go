package main

import "fmt"

type FormPendaftaran struct {
	NamaLengkap string
	Email       string
	JenisKelamin string
	Usia        int
}
// FormPendaftaran adalah struktur data yang digunakan untuk menyimpan informasi pendaftaran
// TampilkanData adalah metode untuk menampilkan informasi pendaftaran

// method yang bersifat pass by value
// artinya ketika kita memanggil method ini, data yang dikirimkan adalah salinan
func (f FormPendaftaran) TampilkanData() {
	fmt.Printf("Nama Lengkap: %s\n", f.NamaLengkap)
	fmt.Printf("Email: %s\n", f.Email)
	fmt.Printf("Jenis Kelamin: %s\n", f.JenisKelamin)
	fmt.Printf("Usia: %d\n", f.Usia)
}
// method yang bersifat pass by reference
func (f *FormPendaftaran) MerubahUsia(usiaBaru int) {
	f.Usia = usiaBaru
	fmt.Printf("Usia telah diubah menjadi: %d\n", f.Usia)
}
// MencetakNamaLengkapDenganFunction adalah fungsi untuk mencetak nama lengkap dari FormPendaftaran
func MencetakNamaLengkapDenganFunction(form FormPendaftaran){
	fmt.Println("Usia:", form.Usia)
}

// bedanya function dengan method adalah
// method adalah fungsi yang terkait dengan tipe data tertentu (struct),
// sedangkan function adalah fungsi umum yang tidak terkait dengan tipe data tertentu.
func main() {
	pendaftaran1 := FormPendaftaran{
		NamaLengkap: "Budi Santoso",
		Email: "email@email.com",
		JenisKelamin: "Laki-laki",
		Usia: 30,
	}
	pendaftaran1.TampilkanData()
	pendaftaran1.MerubahUsia(35)
	MencetakNamaLengkapDenganFunction(pendaftaran1)
}