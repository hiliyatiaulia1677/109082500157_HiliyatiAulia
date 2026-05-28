package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
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

func main() {
	var daerah int
	fmt.Scan(&daerah)

	for d := 0; d < daerah; d++ {
		var jumlahRumah int
		fmt.Scan(&jumlahRumah)

		var rumah arrInt
		for k := 0; k < jumlahRumah; k++ {
			fmt.Scan(&rumah[k])
		}

		selectionSort(&rumah, jumlahRumah)

		for k := 0; k < jumlahRumah; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[k])
		}
		fmt.Println()
	}
}