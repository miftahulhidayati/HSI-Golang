# Concurrency di Go

Concurrency adalah fitur kuat di Go yang memungkinkan Anda menjalankan beberapa tugas secara bersamaan. Pada level ini, Anda akan mempelajari goroutine dan channel.

## Goroutine

Goroutine adalah thread ringan yang dikelola oleh Go runtime. Gunakan kata kunci `go` untuk menjalankan fungsi secara konkuren.

### Contoh

```go
go fungsiSaya()
```

## Channel

Channel digunakan untuk komunikasi antar goroutine.

### Contoh

```go
c := make(chan string)
go func() { c <- "Pesan dari goroutine" }()
pesan := <-c
```

## Studi Kasus

1. **Web Scraper**: Mengambil data dari banyak URL secara konkuren.
2. **Pemrosesan File**: Memproses banyak file secara bersamaan.

## Latihan

1. Buat program yang menjalankan beberapa goroutine untuk mencetak angka 1-10 secara konkuren.
2. Implementasikan fungsi yang menghitung jumlah elemen slice menggunakan goroutine dan channel.

Setelah menyelesaikan level ini, Anda akan memahami dasar concurrency di Go.

Selamat belajar!