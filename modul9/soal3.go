package main

import "fmt"

const NMAX int = 100

type tabelHasil [NMAX]string

func main() {
	var klubA, klubB string
	var hasil tabelHasil
	var n int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	nomor := 1
	for {
		var skorA int
		var skorB int
		fmt.Print("Pertandingan ", nomor, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}
		n++
		nomor++
	}

	for i := 0; i < n; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}