package main

import "fmt"

func main() {
	var t1, t2, t3, t4 string
	var i int
	var hasil bool

	hasil = true
	for i = 1; i <= 5; i++ {
	fmt.Print("Percobaan ", i, ":")
	fmt.Scan(&t1, &t2, &t3, &t4)

	if t1 != "merah" || t2 != "kuning" || t3 != "hijau" || t4 != "ungu" {
	hasil = false
	}
}
fmt.Print("Berhasil: ", hasil)
}