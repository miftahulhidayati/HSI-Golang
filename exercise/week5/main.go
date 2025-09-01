package main

import (
	"week5/pegawai"
)

func main() {
	// Data dummy
	p := pegawai.Pegawai{
		Nama:        "Budi Santoso",
		Posisi:      "Software Engineer",
		GajiBulanan: 8000000,
	}

	// Panggil method HitungGajiTahunan
	gajiTahunan := p.HitungGajiTahunan()
	println("Gaji Tahunan:", gajiTahunan)

	// Gunakan interface InformasiPegawai
	var info pegawai.InformasiPegawai = p
	info.TampilkanInformasi()
}
