package main
import "fmt"


func cetakBintang(N int) {
	if N == 0 {
		return
	}
	cetakBintang(N - 1)
	fmt.Print("*")
}
func pola(N int, i int) {
	if i > N {
		return
	}
	cetakBintang(i)
	fmt.Println()
	pola(N, i+1)
}

func main() {
	var N int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&N)

	pola(N, 1)
}