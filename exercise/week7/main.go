package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Model Pegawai
type Pegawai struct {
	ID           uint    `gorm:"primaryKey"`
	Nama         string
	Posisi       string
	GajiPerBulan float64
}

// Fungsi untuk menghitung gaji setahun
func (p Pegawai) GajiSetahun() float64 {
	return p.GajiPerBulan * 12
}

func main() {
	// Koneksi ke database SQLite
	db, err := gorm.Open(sqlite.Open("pegawai.db"), &gorm.Config{})
	if err != nil {
		panic("gagal koneksi database")
	}

	// Migrasi schema
	db.AutoMigrate(&Pegawai{})

	// 1. Create: Tambah 5 data pegawai (hardcoded)
	pegawaiList := []Pegawai{
		{Nama: "Gita", Posisi: "Manager", GajiPerBulan: 10000000},
		{Nama: "Ikke", Posisi: "Staff", GajiPerBulan: 5000000},
		{Nama: "Aldifa", Posisi: "Supervisor", GajiPerBulan: 7000000},
		{Nama: "Ely", Posisi: "Staff", GajiPerBulan: 5500000},
		{Nama: "Miftah", Posisi: "Admin", GajiPerBulan: 4000000},
	}
	for _, p := range pegawaiList {
		db.Create(&p)
	}

	fmt.Println("== List Pegawai Setelah Create ==")
	tampilPegawai(db)

	// 2. Update: Update gaji per bulan pegawai dengan ID 2
	var pegawai Pegawai
	db.First(&pegawai, 2)
	pegawai.GajiPerBulan = 6000000
	db.Save(&pegawai)
	fmt.Println("\n== Setelah Update Gaji Pegawai ID 2 ==")
	tampilPegawai(db)

	// 3. Delete: Hapus pegawai dengan ID 3
	db.Delete(&Pegawai{}, 3)
	fmt.Println("\n== Setelah Delete Pegawai ID 3 ==")
	tampilPegawai(db)
}

// Fungsi untuk menampilkan list pegawai beserta gaji setahun
func tampilPegawai(db *gorm.DB) {
	var pegawai []Pegawai
	db.Find(&pegawai)
	for _, p := range pegawai {
		fmt.Printf("ID: %d, Nama: %s, Posisi: %s, Gaji/Bulan: %.0f, Gaji/Setahun: %.0f\n", p.ID, p.Nama, p.Posisi, p.GajiPerBulan, p.GajiSetahun())
	}
}
