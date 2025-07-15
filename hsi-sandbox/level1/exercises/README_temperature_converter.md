# Temperature Converter / Konverter Suhu

Aplikasi console untuk mengkonversi suhu dari Celsius ke Fahrenheit dan Reamur.

## Fitur

- Konversi suhu dari Celsius ke Fahrenheit
- Konversi suhu dari Celsius ke Reamur
- Validasi input (hanya menerima angka)
- Interface console yang user-friendly

## Cara Menggunakan

### Menjalankan Aplikasi

```bash
go run temperature_converter.go
```

### Contoh Penggunaan

#### Input Valid (Angka)
```
--- Konverter Suhu ---
Masukkan suhu dalam Celsius: 25
> Suhu dalam Reamur = 20
> Suhu dalam Fahrenheit = 77
```

#### Input Tidak Valid (String)
```
--- Konverter Suhu ---
Masukkan suhu dalam Celsius: dua puluh lima
> Input Tidak Valid, hanya menerima angka
```

### Menjalankan Test
Untuk melihat berbagai skenario test:

```bash
go run test_temperature_converter.go
```

## Rumus Konversi

### Celsius ke Reamur
```
R = C × 4/5
```

### Celsius ke Fahrenheit
```
F = (C × 9/5) + 32
```

## Teknologi yang Digunakan

- Go (Golang)
- Package `bufio` untuk input console
- Package `strconv` untuk konversi string ke angka
- Package `strings` untuk manipulasi string

## Struktur Kode

- `temperature_converter.go` - Aplikasi utama
- `test_temperature_converter.go` - Script untuk testing berbagai skenario
- `README_temperature_converter.md` - Dokumentasi aplikasi

## Catatan Implementasi

1. Menggunakan `bufio.Reader` untuk membaca input dari console
2. Menggunakan `strconv.ParseFloat` untuk validasi input numerik
3. Menggunakan `strings.TrimSpace` untuk membersihkan input
4. Format output menggunakan `%.0f` untuk menampilkan angka bulat