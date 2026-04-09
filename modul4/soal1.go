package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i 
	}
}
func permutasi(n, r int, hasil *int) {
	var an, anr int
	faktorial(n, &an)
	faktorial(n-r, &anr)
	*hasil = an / anr
}
func kombinasi(n, r int, hasil *int) {
	var an, ar, anr int
	faktorial(n, &an)
	faktorial(r, &ar)
	faktorial(n-r, &anr)
	*hasil = an / (ar * anr)
}

func main() {
	var a, b, c, d, hasilP, hasilC int
	fmt.Scan(&a, &b, &c, &d)

	permutasi(a, c, &hasilP)
	kombinasi(a, c, &hasilC)
	fmt.Println(hasilP, hasilC)

	permutasi(b, d, &hasilP)
	kombinasi(b, d, &hasilC)
	fmt.Println(hasilP, hasilC)
}