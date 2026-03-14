package main

import "fmt"

func main() {
	var beratParsel int
	var beratKg int
	var sisaGram int
	var biayaPerKg int
	var biayaPerGram int
	var totalBiaya int

	fmt.Print("Masukkan berat parsel (Gram): ")
	fmt.Scan(&beratParsel)

	beratKg = beratParsel / 1000
	sisaGram = beratParsel % 1000
	biayaPerKg = beratKg * 10000

	if beratKg > 10 {
	totalBiaya = 0
	} else if sisaGram < 500 {
	biayaPerGram = sisaGram * 15
	} else {
	biayaPerGram = sisaGram * 5
	}

	totalBiaya = biayaPerKg + biayaPerGram
	fmt.Println("Berat Parsel (Gram): ",beratParsel)
	fmt.Println("Detail Berat: ", beratKg,"kg +", sisaGram, "gr")
	fmt.Println("Biaya Kg: ", "Rp.", biayaPerKg, "+ Rp.", biayaPerGram)
	fmt.Println("Total Biaya: ", totalBiaya)
}