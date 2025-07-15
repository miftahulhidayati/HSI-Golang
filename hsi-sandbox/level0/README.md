# Persiapan Lingkungan Pengembangan Golang

Selamat datang di Level 0 **HSI - Golang Sandbox**! Pada level ini, Anda akan mempelajari cara menyiapkan lingkungan pengembangan untuk pemrograman Go, mulai dari instalasi Go di berbagai sistem operasi hingga konfigurasi IDE.

## Daftar Isi

1. [Instalasi Go](#instalasi-go)
   - [Windows](#instalasi-go-di-windows)
   - [Linux](#instalasi-go-di-linux)
2. [Pengaturan IDE](#pengaturan-ide)
   - [Visual Studio Code](#pengaturan-visual-studio-code)
   - [Sublime Text](#pengaturan-sublime-text)
3. [Verifikasi Instalasi](#verifikasi-instalasi)

## Instalasi Go

### Instalasi Go di Windows

1. Unduh installer Go dari [situs resmi Go](https://golang.org/dl/).
2. Jalankan installer dan ikuti petunjuknya.
3. Setelah selesai, buka Command Prompt dan jalankan `go version` untuk memastikan instalasi berhasil.

### Instalasi Go di Linux

1. Buka terminal.
2. Jalankan perintah berikut untuk mengunduh dan menginstal Go:
   ```bash
   wget https://golang.org/dl/go1.XX.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.XX.linux-amd64.tar.gz
   ```
3. Tambahkan Go ke PATH dengan menambahkan baris berikut ke `.bashrc` atau `.bash_profile`:
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
4. Jalankan `source ~/.bashrc` dan verifikasi dengan `go version`.

## Pengaturan IDE

### Pengaturan Visual Studio Code

1. Instal Visual Studio Code dari [situs resmi](https://code.visualstudio.com/).
2. Buka VS Code dan instal ekstensi Go.
3. Ikuti petunjuk untuk konfigurasi workspace Go.



## Verifikasi Instalasi

Buat file `main.go` dengan isi berikut:

```go
package main

import "fmt"

func main() {
    fmt.Println("Halo, Dunia!")
}
```

Jalankan program menggunakan perintah `go run main.go`. Jika muncul "Halo, Dunia!", maka instalasi Anda sudah benar!

Selamat belajar!
