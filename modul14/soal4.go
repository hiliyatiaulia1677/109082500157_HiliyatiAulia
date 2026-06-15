package main

import "fmt"

func insertionSort(T *[100]int, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func checkJarak(T *[100]int, n int) {
    var jarak, i int
    jarak = T[1] - T[0]
    i = 2
    for i <= n-1 {
        if T[i]-T[i-1] != jarak {
            fmt.Printf("Data berjarak tidak tetap\n")
            return
        }
        i = i + 1
    }
    fmt.Printf("Data berjarak %d\n", jarak) 
}

func main() {
	var T [100]int
	var angka, n int

	n = 0
	fmt.Scan(&angka)
	for angka >= 0 {
		T[n] = angka
		n = n + 1
		fmt.Scan(&angka)
	}

	insertionSort(&T, n)

	var i int
	i = 0
	for i <= n-1 {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
		i = i + 1
	}
	fmt.Println()

	(checkJarak(&T, n))
}