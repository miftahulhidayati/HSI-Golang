# Struct dan Method di Go

Pada level ini, Anda akan mempelajari tipe data struct dan method di Go. Struct digunakan untuk mengelompokkan data terkait, sedangkan method memungkinkan Anda mendefinisikan fungsi yang terkait dengan struct tersebut. Memahami cara menggunakan struct dan method sangat penting untuk membangun aplikasi Go yang lebih kompleks.

## Apa itu Struct?

Struct adalah tipe data komposit yang mengelompokkan beberapa variabel (field) dalam satu nama. Mereka berguna untuk memodelkan entitas dunia nyata dan dapat berisi field dengan tipe yang berbeda.

### Contoh Struct

```go
package main

import "fmt"

// Define a struct
type Person struct {
    Nama string
    Umur int
}

func main() {
    // Create an instance of Person
    p := Person{Nama: "Alice", Umur: 30}
    fmt.Println(p)
}
```

## Method di Go

Method adalah fungsi yang terkait dengan tipe tertentu (biasanya struct). Di Go, kita dapat mendefinisikan method pada struct, memungkinkan kita untuk mengoperasikan data yang terkandung dalam struct tersebut.

### Contoh Method

```go
package main

import "fmt"

// Define a struct
type PersegiPanjang struct {
    Lebar  float64
    Tinggi float64
}

// Define a method for Rectangle
func (p PersegiPanjang) Luas() float64 {
    return p.Lebar * p.Tinggi
}

func main() {
    // Create an instance of Rectangle
    rect := PersegiPanjang{Lebar: 10, Tinggi: 5}
    fmt.Println("Luas persegi panjang:", rect.Luas())
}
```

## Studi Kasus

1. **Model Buku Perpustakaan**: Buat struct untuk buku, lengkap dengan method untuk menampilkan detail dan mengecek ketersediaan.
2. **Manajemen Pegawai**: Buat struct pegawai dengan method untuk menaikkan gaji dan menampilkan informasi pegawai.

## Ringkasan

Struct dan method sangat penting untuk membangun aplikasi Go yang kompleks dan terstruktur. Struct memungkinkan kita mengelompokkan data terkait, sementara method memungkinkan kita mendefinisikan perilaku yang terkait dengan data tersebut. Memahami konsep-konsep ini sangat penting untuk membangun aplikasi Go yang lebih kompleks.

## Langkah Selanjutnya

Lanjutkan ke level berikutnya untuk mempelajari interface di Go, yang akan lebih meningkatkan kemampuan kita dalam merancang kode yang fleksibel dan dapat digunakan kembali.

Selamat belajar!