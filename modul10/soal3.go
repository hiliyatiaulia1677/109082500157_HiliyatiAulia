package main

import "fmt"

type BeratBalita [100]float64

func hitungMinMax(arrBerat BeratBalita, n int, BeratMin, BeratMax *float64) {
	*BeratMin = arrBerat[0]
	*BeratMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *BeratMin {
			*BeratMin = arrBerat[i]
		}
		if arrBerat[i] > *BeratMax {
			*BeratMax = arrBerat[i]
		}
	}
}

func rata_rata(arrBerat BeratBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data BeratBalita
	var n int

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d (kg): ", i+1)
		fmt.Scan(&data[i])
	}

	var BeratMin, BeratMax float64
	hitungMinMax(data, n, &BeratMin, &BeratMax)
	average := rata_rata(data, n)

	fmt.Printf("\nBerat minimum balita : %.2f kg\n", BeratMin)
	fmt.Printf("Berat maksimum balita: %.2f kg\n", BeratMax)
	fmt.Printf("rata-rata berat balita  : %.2f kg\n", average)
}