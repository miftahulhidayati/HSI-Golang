package main

import (
	"fmt"
	"time"
)

func cetakPesan(pesan string, channelGoRoutine chan bool) {
	for i := 0; i < 5; i++ {
		println("Pesan ke-", i+1, ":", pesan)
		time.Sleep(1 * time.Millisecond)
	}
	channelGoRoutine <- true
}

func main1() {
	fmt.Println("Memulai program...")
	channel := make(chan bool, 2)
	// cetakPesan("Makan gapake nasi")
	// goRoutine utama
	go cetakPesan("Nasi Goreng", channel)
	// goRoutine kedua
	go cetakPesan("Mie Goreng", channel)
	// go routine utama akan menunggu goroutine lainnya selesai atau sempat berjalan
	// dengan cara menunggu selama 5 milidetik
	// agar goroutine lainnya sempat berjalan
	// time.Sleep(20 * time.Millisecond)
	// tidak akan mencetak pesan "Selesai" sebelum channel menerima data dan bernilai true
	<-channel // menunggu goroutine selesai
	fmt.Print("Selesai")
}

