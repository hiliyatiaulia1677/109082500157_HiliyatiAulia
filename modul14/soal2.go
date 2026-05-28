package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Scan(&m)

		var ganjil, genap arrInt
		var jumlahGanjil, jumlahGenap int

		for i := 0; i < m; i++ {

			var bil int
			fmt.Scan(&bil)

			if bil%2 == 1 {
				ganjil[jumlahGanjil] = bil
				jumlahGanjil++
			} else {
				genap[jumlahGenap] = bil
				jumlahGenap++
			}
		}

		selectionSortAsc(&ganjil, jumlahGanjil)
		selectionSortDesc(&genap, jumlahGenap)

		awal := true

		for i := 0; i < jumlahGanjil; i++ {

			if !awal {
				fmt.Print(" ")
			}

			fmt.Print(ganjil[i])
			awal = false
		}

		for i := 0; i < jumlahGenap; i++ {

			if !awal {
				fmt.Print(" ")
			}

			fmt.Print(genap[i])
			awal = false
		}

		fmt.Println()
	}
}