# <h1 align="center">Laporan Praktikum Modul 3 - Algoritma Pemrograman 2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Modul 3 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

#### soal1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutasi(a, c), kombinasi(a, c))

	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul3/output/soal1.png)
Penjelasan : 
jadi program ini digunakan untuk menghitung faktorial, permutasi, dan juga kombinasi. func faktorial digunakan untuk mengkalikan angka 1 sampai n. func permutasi digunakan untuk menghitung susunan dengan menggunakan rumus faktorial yaitu n! dibagi (n-r) jadi func ini bakalan manggil func faktorial yang tadi. func kombinasi buat ngitung pilihan rumusnya n! dibagi r! dikali (n-r)! ini juga pake rumus faktorial. user bakaln masukin 4 angka yaitu a, b, c, dan d. lalu program bakalan menghitung permutasi dan kombinasi dari a dan c lalu ditampilkan dibaris pertama. kemudian menghitung lagi permutasi dan kombinasi dari b dan d yang hasilnya akan ditampilak dibaris kedua

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))

	fmt.Println(g(h(f(b))))

	fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul2/output/soal2.png)
Penjelasan : 
Program ini digunakan memproses angka dengan menggunakan tiga operasi sederhana namun urutannya berbeda-beda untuk setiap input. program akan meminta pengguna untuk memasukkan tiga angka yaitu a, b, dan c. untuk a, akan ditambah dengan 1 lalu dikurangi 2 kemudian hasilnya dikalikan dengan dirinya sendiri. untuk b, angka akan dikalikan dulu lalu ditambah 1 kemudian dikurangi 2. untuk c, angka dikurangi 2 lalu dikalikan dengan dirinya sendiri kemuadian ditambah 1.

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	var lingkaran1, lingkaran2 bool

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	lingkaran1 = didalam(cx1, cy1, r1, x, y)
	lingkaran2 = didalam(cx2, cy2, r2, x, y)

	if lingkaran1 && lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if lingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul2/output/soal3.png)
Penjelasan :
Program ini digunakan untuk mengecek sebuah titik apakah dia berada dalam satu atau dua lingkaran. func jarak digunakan untuk menghitung jarak antara titik pusat lingkaran dan titik yang akan dicek. sedangkan func didalam menentukan apakah titik tersebut berada di dalam lingkaran berdasarkan jari-jari. program akan membaca data dua lingkaran dan satu titik, lalu mengecek posisi titik terhadap masing-masing lingkaran.