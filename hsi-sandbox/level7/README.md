# Level 7: Pengujian (Testing) di Go

Pada level ini, Anda akan mempelajari konsep pengujian di Go. Pengujian sangat penting untuk memastikan kode Anda berjalan dengan benar dan andal.

## Ringkasan Pengujian di Go

1. **Mengapa Perlu Pengujian?**
   - Menemukan bug lebih awal.
   - Memastikan kode berjalan sesuai harapan.
   - Dokumentasi perilaku kode.

2. **Menulis Pengujian**
   - Pengujian ditulis di file dengan akhiran `_test.go`.
   - Gunakan package `testing`.
   - Fungsi pengujian diawali dengan `Test` dan menerima parameter `*testing.T`.

3. **Menjalankan Pengujian**
   - Jalankan dengan perintah `go test`.
   - Semua fungsi pengujian akan dijalankan otomatis.

4. **Struktur Pengujian**
   - Fungsi pengujian
   - Fungsi pembantu (jika perlu)
   - Assertion untuk memeriksa hasil

## Studi Kasus: Pengujian Fungsi Sederhana

Misal, fungsi penjumlahan dua bilangan bulat dan pengujiannya.

```go
package main

import "testing"

// Fungsi Tambah untuk diuji
func Tambah(a int, b int) int {
    return a + b
}

// Pengujian untuk fungsi Tambah
func TestTambah(t *testing.T) {
    hasil := Tambah(2, 3)
    ekspektasi := 5
    if hasil != ekspektasi {
        t.Errorf("Tambah(2, 3) = %d; ingin %d", hasil, ekspektasi)
    }
}
```

## Latihan

1. Buat pengujian untuk fungsi pengurangan dua bilangan bulat.
2. Buat pengujian untuk fungsi pengecekan bilangan genap.
3. Buat pengujian untuk fungsi penggabungan dua string.

Setelah menyelesaikan level ini, Anda akan terbiasa menulis dan menjalankan pengujian di Go.

Selamat belajar!