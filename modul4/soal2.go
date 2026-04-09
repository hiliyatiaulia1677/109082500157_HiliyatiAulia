package main

import "fmt"

func hitungSkor(soal, skor *int) {
	*soal = 0
	*skor = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int
	maxSoal := -1
	minSkor := 999999

	for {
		fmt.Scan(&nama)
		if nama == "selesai" {
			break
		}
		hitungSkor(&soal, &skor)
	

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}