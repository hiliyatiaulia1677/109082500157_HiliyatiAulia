package main

import "fmt"

func main() {
	var berat [1000]float64
	var x, y int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	totalPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		wadah := i / y
		totalPerWadah[wadah] += berat[i]
	}
	
	totalSeluruh := 0.0
	for i := 0; i < jumlahWadah; i++ {
		totalSeluruh += totalPerWadah[i]
	}
	rataRata := totalSeluruh / float64(jumlahWadah)

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", totalPerWadah[i])
	}
	fmt.Println()
	fmt.Printf("%.2f\n", rataRata)
}