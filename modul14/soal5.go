package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	var i int
	i = 0
	for i <= *n-1 {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
		i = i + 1
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var maks int
	var idxMaks int
	maks = pustaka[0].rating
	idxMaks = 0
	var i int
	i = 1
	for i <= n-1 {
		if pustaka[i].rating > maks {
			maks = pustaka[i].rating
			idxMaks = i
		}
		i = i + 1
	}
	fmt.Println(pustaka[idxMaks].judul, pustaka[idxMaks].penulis, pustaka[idxMaks].penerbit, pustaka[idxMaks].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var batas int
	if n < 5 {
		batas = n
	} else {
		batas = 5
	}
	var i int
	i = 0
	for i <= batas-1 {
		fmt.Println(pustaka[i].judul)
		i = i + 1
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kiri, kanan, tengah int
	var ketemu bool
	kiri = 0
	kanan = n - 1
	ketemu = false
	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if ketemu {
		fmt.Println(pustaka[tengah].judul, pustaka[tengah].penulis, pustaka[tengah].penerbit, pustaka[tengah].tahun, pustaka[tengah].eksemplar, pustaka[tengah].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var r int

	DaftarkanBuku(&pustaka, &nPustaka)

	fmt.Scan(&r)

	CetakTerfavorit(pustaka, nPustaka)

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)

	CariBuku(pustaka, nPustaka, r)
}