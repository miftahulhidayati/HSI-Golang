# Level 4: Interface di Go

Pada level ini, Anda akan mempelajari konsep interface di Go. Interface memungkinkan Anda mendefinisikan kontrak perilaku yang dapat diimplementasikan oleh berbagai tipe.

## Apa itu Interface?

Interface adalah tipe yang mendefinisikan sekumpulan method. Tipe apa pun yang mengimplementasikan method tersebut dianggap memenuhi interface.

### Poin Penting:
- Interface didefinisikan dengan kata kunci `interface`.
- Tipe yang mengimplementasikan semua method pada interface otomatis memenuhi interface tersebut.
- Interface memungkinkan polimorfisme.

## Studi Kasus

- **Sistem Logging**: Implementasi interface untuk berbagai strategi logging (misal: ke konsol, ke file).
- **Pembayaran**: Sistem pembayaran dengan berbagai metode yang mengimplementasikan interface yang sama.

## Latihan

1. Buat interface `Hewan` dengan method `Bersuara()` dan `Bergerak()`. Implementasikan pada struct `Anjing` dan `Kucing`.
2. Buat fungsi yang menerima interface `Hewan` dan memanggil method `Bersuara()`.
3. Buat interface `Bentuk` dengan method `Luas()` dan `Keliling()`. Implementasikan pada struct `PersegiPanjang` dan `Lingkaran`.

Setelah menyelesaikan level ini, Anda akan memahami cara mendefinisikan dan menggunakan interface di Go.

Selamat belajar!