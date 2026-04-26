package main

import "fmt"

func misteri(arr *[6]int, n int) int {
	hasil := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			hasil += arr[i]
		} else {
			hasil -= arr[i]
		}
	}
	return hasil
}

func geser(arr *[6]int, idx int, n int) {
	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n-1] = 0
}

func main() {
	data := [6]int{9, 3, 6, 2, 8, 4}
	n := 6

	geser(&data, 2, n)
	n--

	hasil := misteri(&data, n)

	fmt.Println(hasil)
}