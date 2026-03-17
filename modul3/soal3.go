// package main

// import (
// 	"fmt"
// 	"math"
// )

// // fungsi menghitung jarak dua titik
// func distance(x1, y1, x2, y2 int) float64 {
// 	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
// }

// // fungsi mengecek apakah titik di dalam lingkaran
// func inside(cx, cy, r, x, y int) bool {
// 	return distance(cx, cy, x, y) <= float64(r)
// }

// func main() {
// 	var cx1, cy1, r1 int
// 	var cx2, cy2, r2 int
// 	var x, y int

// 	// input lingkaran 1
// 	fmt.Scan(&cx1, &cy1, &r1)

// 	// input lingkaran 2
// 	fmt.Scan(&cx2, &cy2, &r2)

// 	// input titik
// 	fmt.Scan(&x, &y)

// 	in1 := inside(cx1, cy1, r1, x, y)
// 	in2 := inside(cx2, cy2, r2, x, y)

// 	if in1 && in2 {
// 		fmt.Println("Titik di dalam lingkaran 1 dan 2")
// 	} else if in1 {
// 		fmt.Println("Titik di dalam lingkaran 1")
// 	} else if in2 {
// 		fmt.Println("Titik di dalam lingkaran 2")
// 	} else {
// 		fmt.Println("Titik di luar lingkaran 1 dan 2")
// 	}
// }

package main

import (
	"fmt"
	"math"
)

// function jarak(a,b,c,d : real) -> real
// Mengembalikan jarak antara titik (a,b) dan (c,d)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// function didalam(cx,cy,r,x,y : real) -> boolean
// Mengembalikan true jika titik (x,y) berada di dalam lingkaran
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	// input lingkaran 1
	fmt.Scan(&cx1, &cy1, &r1)

	// input lingkaran 2
	fmt.Scan(&cx2, &cy2, &r2)

	// input titik
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}