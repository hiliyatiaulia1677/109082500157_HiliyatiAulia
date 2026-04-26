package main

import (
	"fmt"
	"math"
)

type IsiArrayInt [100]int

func cetakArray(data IsiArrayInt, n int) {
	fmt.Print("Semua elemen : ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func cetakGanjil(data IsiArrayInt, n int) {
	fmt.Print("Indeks ganjil : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func cetakGenap(data IsiArrayInt, n int) {
	fmt.Print("Indeks genap : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func tampilKelipatanX(data IsiArrayInt, n int, x int) {
	fmt.Print("Indeks kelipatan ", x, " : ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndex(data *IsiArrayInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n = *n - 1
}

func hitungRata(data IsiArrayInt, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total = total + data[i]
	}
	rata := float64(total) / float64(n)
	return rata
}

func hitungStdDev(data IsiArrayInt, n int) float64 {
	rata := hitungRata(data, n)
	jumlah := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - rata
		jumlah = jumlah + (selisih * selisih)
	}
	hasil := math.Sqrt(jumlah / float64(n))
	return hasil
}

func hitungFrekuensi(data IsiArrayInt, n int, angka int) int {
	count := 0
	for i := 0; i < n; i++ {
		if data[i] == angka {
			count++
		}
	}
	return count
}

func main() {
	var data IsiArrayInt
	var n int

	fmt.Print("Masukkan jumlah elemen : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Elemen ke-", i, " : ")
		fmt.Scan(&data[i])
	}

	fmt.Println()

	// a
	cetakArray(data, n)

	// b
	cetakGanjil(data, n)

	// c
	cetakGenap(data, n)

	// d
	var x int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	tampilKelipatanX(data, n, x)

	// e
	var idx int
	fmt.Print("Masukkan indeks yang dihapus : ")
	fmt.Scan(&idx)
	hapusIndex(&data, &n, idx)
	fmt.Print("Array setelah dihapus : ")
	cetakArray(data, n)

	// f
	fmt.Printf("Rata-rata : %.2f\n", hitungRata(data, n))

	// g
	fmt.Printf("Standar deviasi : %.2f\n", hitungStdDev(data, n))

	// h
	var cari int
	fmt.Print("Cari frekuensi angka : ")
	fmt.Scan(&cari)
	fmt.Println("Frekuensi", cari, ":", hitungFrekuensi(data, n, cari), "kali")
}