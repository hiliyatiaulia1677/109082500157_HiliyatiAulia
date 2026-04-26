package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	center titik
	r     int
}

func jarak(p, q titik) float64 {
	hasil := math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
	return hasil
}

func didalam(c lingkaran, p titik) bool {
	d := jarak(c.center, p)
	if d <= float64(c.r) {
		return true
	}
	return false
}

func main() {
	var l1 lingkaran
	var l2 lingkaran
	var p titik

	fmt.Scan(&l1.center.x, &l1.center.y, &l1.r)
	fmt.Scan(&l2.center.x, &l2.center.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	cek1 := didalam(l1, p)
	cek2 := didalam(l2, p)

	if cek1 == true && cek2 == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cek1 == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cek2 == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
