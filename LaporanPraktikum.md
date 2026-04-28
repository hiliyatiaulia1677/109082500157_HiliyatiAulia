# <h1 align="center">Laporan Praktikum Modul 10 - Algoritma Pemrograman 2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Modul 10

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### soal1.go

```go
package main

import "fmt"

func main() {
	var berat [1000]float64
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Printf("min :%.2f\nmax :%.2f\n", min, max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul10/output/soal1.png)
Penjelasan : 
Program ini digunakan untuk mendata berat anak kelinci yang akan dijual. Pengguna akan memasukkan jumlah anak kelinci dan berat anak kelinci dari sejumlah anak kelinci yang telah dimasukkan tadi, lalu berat masing-masing disimpan dalam array. Berat pertama dijadikan nilai minimum dan maksimum sementara. Selanjutnya, setiap data dibandingkan satu per satu untuk mencari berat terkecil dan terbesar. Setelah semua data diperiksa, program menampilkan berat minimum dan maksimum dengan dua angka di belakang koma.


### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
package main

import "fmt"

func main() {
	var berat [1000]float64
	var x, y int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	totalPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		wadah := i / y
		totalPerWadah[wadah] += berat[i]
	}
	
	totalSeluruh := 0.0
	for i := 0; i < jumlahWadah; i++ {
		totalSeluruh += totalPerWadah[i]
	}
	rataRata := totalSeluruh / float64(jumlahWadah)

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", totalPerWadah[i])
	}
	fmt.Println()
	fmt.Printf("%.2f\n", rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul10/output/soal2.png)
Penjelasan : 
Program ini digunakan untuk menghitung total berat ikan pada setiap wadah dan rata-rata berat seluruh wadah. Pengguna memasukkan jumlah ikan(x) dan kapasitas setiap wadah(y), kemudian berat masing-masing ikan disimpan ke dalam array.Selanjutnya, program membagi ikan ke dalam beberapa wadah sesuai urutan input. Berat ikan dalam setiap wadah dijumlahkan dan disimpan. Setelah semua wadah terisi, program menghitung rata-rata berat seluruh wadah dengan membagi total keseluruhan berat dengan jumlah wadah.Terakhir, program menampilkan total berat pada setiap wadah di baris pertama, lalu rata-rata berat per wadah di baris kedua.


### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

#### soal3.go

```go
package main

import "fmt"

type BeratBalita [100]float64

func hitungMinMax(arrBerat BeratBalita, n int, BeratMin, BeratMax *float64) {
	*BeratMin = arrBerat[0]
	*BeratMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *BeratMin {
			*BeratMin = arrBerat[i]
		}
		if arrBerat[i] > *BeratMax {
			*BeratMax = arrBerat[i]
		}
	}
}

func rata_rata(arrBerat BeratBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data BeratBalita
	var n int

	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d (kg): ", i+1)
		fmt.Scan(&data[i])
	}

	var BeratMin, BeratMax float64
	hitungMinMax(data, n, &BeratMin, &BeratMax)
	average := rata_rata(data, n)

	fmt.Printf("\nBerat minimum balita : %.2f kg\n", BeratMin)
	fmt.Printf("Berat maksimum balita: %.2f kg\n", BeratMax)
	fmt.Printf("rata-rata berat balita  : %.2f kg\n", average)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul10/output/soal3.png)
Penjelasan :
Program ini digunakan oleh petugas posyandu untuk mencatat berat balita ke dalam sebuah array. Pertama, petugas memasukkan jumlah balita, lalu berat masing-masing balita diinput dan disimpan.Setelah semua data masuk, program mencari berat balita terkecil dan terbesar dengan membandingkan setiap nilai. Selain itu, program juga menghitung rata-rata berat dengan menjumlahkan seluruh data, kemudian membaginya dengan jumlah balita.Terakhir, program menampilkan berat minimum, berat maksimum, dan rata-rata berat seluruh balita.
