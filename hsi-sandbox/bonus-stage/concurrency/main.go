package main

import "errors"

type rawData struct {
	data         []int
	channel      chan int
	channelError chan error
}

func hitungData(rd rawData) {
	total := 0
	if len(rd.data) == 0 {
		rd.channel <- 0
		rd.channelError <- errors.New("Data tidak boleh kosong")
		return
	}
	for _, value := range rd.data {
		total += value
	}
	rd.channel <- total
	rd.channelError <- nil
}

func main() {
	angkaGenap := []int{2, 4, 6, 8, 10}
	angkaGanjil := []int{}
	c1 := make(chan int)
	c2 := make(chan int)
	c1Error := make(chan error)
	c2Error := make(chan error)

	deret1 := rawData{data: angkaGenap, channel: c1, channelError: c1Error}
	deret2 := rawData{data: angkaGanjil, channel: c2, channelError: c2Error}

	go hitungData(deret1)
	go hitungData(deret2)

	totalAngkaGenap := <-c1  // tunggu dan ambil total dari channel c1
	totalAngkaGanjil := <-c2 // tunggu dan ambil total dari channel c2

	err1 := <-c1Error
	err2 := <-c2Error

	if err1 != nil {
		println("Error angka genap:", err1.Error())
	}
	if err2 != nil {
		println("Error angka ganjil:", err2.Error())
	}
	println("Total Angka Genap:", totalAngkaGenap)
	println("Total Angka Ganjil:", totalAngkaGanjil)

	println("Total Keseluruhan:", totalAngkaGenap+totalAngkaGanjil)
}