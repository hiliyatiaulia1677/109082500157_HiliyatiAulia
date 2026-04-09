package main
import "fmt"

func fibonacci(a int) int {
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		return fibonacci(a-1) + fibonacci(a-2)
	}
}

func main() {
	var a int
	fmt.Print("Masukkan jumlah suku: ")
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}