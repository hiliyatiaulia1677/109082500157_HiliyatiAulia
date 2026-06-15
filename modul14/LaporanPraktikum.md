# <h1 align="center">Laporan Praktikum Modul 14 - Algoritma Pemrograman 2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Modul 14

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.

#### soal1.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var daerah int
	fmt.Scan(&daerah)

	for d := 0; d < daerah; d++ {
		var jumlahRumah int
		fmt.Scan(&jumlahRumah)

		var rumah arrInt
		for k := 0; k < jumlahRumah; k++ {
			fmt.Scan(&rumah[k])
		}

		selectionSort(&rumah, jumlahRumah)

		for k := 0; k < jumlahRumah; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[k])
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul14/output/soal1.png)
Penjelasan : 
Program ini intinya buat ngurutin nomor rumah dari yang paling kecil ke paling gede pakai metode selection sort. Jadi kita masukin dulu jumlah daerah, terus tiap daerah masukin jumlah rumah dan nomor rumahnya. Nah dii fungsi selectionSort, program bakal nyari angka paling kecil terus dituker ke depan. Misalnya ada angka 8 3 5 1, program bakal cari yang paling kecil yaitu 1, terus dituker ke posisi paling depan. Proses ini diulang terus sampai semua angka terurut. Variabel i dipakai buat nunjuk posisi sekarang, j buat ngecek angka lain, terus idx_min buat nyimpen posisi angka terkecil yang ketemu. Variabel t cuma dipakai sementara buat bantu tuker angka.


### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

#### soal2.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Scan(&m)

		var ganjil, genap arrInt
		var jumlahGanjil, jumlahGenap int

		for i := 0; i < m; i++ {

			var bil int
			fmt.Scan(&bil)

			if bil%2 == 1 {
				ganjil[jumlahGanjil] = bil
				jumlahGanjil++
			} else {
				genap[jumlahGenap] = bil
				jumlahGenap++
			}
		}

		selectionSortAsc(&ganjil, jumlahGanjil)
		selectionSortDesc(&genap, jumlahGenap)

		awal := true

		for i := 0; i < jumlahGanjil; i++ {

			if !awal {
				fmt.Print(" ")
			}

			fmt.Print(ganjil[i])
			awal = false
		}

		for i := 0; i < jumlahGenap; i++ {

			if !awal {
				fmt.Print(" ")
			}

			fmt.Print(genap[i])
			awal = false
		}

		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul14/output/soal2.png)
Penjelasan : 
Program ini dipakai buat misahin angka ganjil dan genap, terus ngurutin dengan cara yang beda. Angka ganjil diurut dari kecil ke besar, sedangkan angka genap diurut dari besar ke kecil. Fungsi selectionSortAsc dipakai buat ngurutin naik. Cara kerjanya nyari angka paling kecil terus dituker ke depan sampai semua angka urut. Nah kalau selectionSortDesc kebalikannya, dia nyari angka paling besar terus dipindah ke depan jadi hasilnya urut menurun. Di bagian func main, program bakal baca dulu jumlah data yang mau diproses. Setelah itu tiap angka dicek, kalau angkanya ganjil masuk ke array ganjil, kalau genap masuk ke array genap. Setelah dipisah, array ganjil diurut naik dan array genap diurut turun. Terakhir semua angka ditampilin, dimulai dari ganjil dulu baru genap. Jadi misalnya inputnya 1 2 3 4 5 6, hasilnya bisa jadi 1 3 5 6 4 2.


### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var i, j, temp int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	var bilangan int
	n := 0

	for {
		fmt.Scan(&bilangan)
		if bilangan == -5313 {
			break
		}
		if bilangan == 0 {
			insertionSort(&data, n)
			median := 0
			if n%2 != 0 {
				median = data[n/2]
			} else {
				median = (data[n/2-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data[n] = bilangan
			n++
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul14/output/soal3.png)
Penjelasan :
Program ini dipakai buat nyimpen angka terus nyari median dari angka-angka yang udah masuk. Data bakal diurut dulu pakai insertion sort sebelum mediannya dihitung. Di fungsi insertionSort, program ngurutinnya dengan cara ngambil satu angka terus nyisipin dia ke posisi yang pas. Jadi kalau ada angka yang lebih besar di depannya, angka itu digeser dulu sampai posisi yang cocok ketemu. Proses ini diulang terus sampai semua angka urut dari kecil ke besar. Di func main, program bakal terus baca input angka. Kalau inputnya -5313, program langsung berhenti. Kalau inputnya 0, berarti program harus ngitung median dari data yang udah masuk sebelumnya.Kalau jumlah datanya ganjil, median diambil dari angka tengah. Kalau jumlah datanya genap, median didapat dari rata-rata dua angka tengah. Setelah itu hasil mediannya ditampilin.Kalau inputnya bukan 0 atau -5313, angka bakal disimpan ke array data.

### 4. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal4.go

```go
package main

import "fmt"

func insertionSort(T *[100]int, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func checkJarak(T *[100]int, n int) {
    var jarak, i int
    jarak = T[1] - T[0]
    i = 2
    for i <= n-1 {
        if T[i]-T[i-1] != jarak {
            fmt.Printf("Data berjarak tidak tetap\n")
            return
        }
        i = i + 1
    }
    fmt.Printf("Data berjarak %d\n", jarak) 
}

func main() {
	var T [100]int
	var angka, n int

	n = 0
	fmt.Scan(&angka)
	for angka >= 0 {
		T[n] = angka
		n = n + 1
		fmt.Scan(&angka)
	}

	insertionSort(&T, n)

	var i int
	i = 0
	for i <= n-1 {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
		i = i + 1
	}
	fmt.Println()

	(checkJarak(&T, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul14/output/soal4.png)
Penjelasan :
Program ini buat baca sekumpulan angka non-negatif, ngurutin pake insertion sort, terus ngecek apakah jarak antar angka yang udah terurut itu sama atau tidak. Program dimulai dengan membuat array T untuk menyimpan data. Lalu membaca angka satu per satu dari input, jika angka yang dibaca tidak negatif maka disimpan ke array T, jika negatif maka berhenti membaca. Setelah semua data masuk, array diurutkan dengan insertionSort, lalu dicetak, lalu dicek jaraknya dengan checkJarak. 


### 5. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
const nMax : integer = 7919
type Buku = <
id, judul, penulis, penerbit : string
eksemplar, tahun, rating : integer >
type DaftarBuku = array [ 1..nMax] of Buku
Pustaka : DaftarBuku
nPustaka: integer
Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang
menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya,
masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris
terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.
Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua
adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku
yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### soal5.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul14/output/soal5.png)
Penjelasan :
Program ini digunakan untuk mengelola data buku di perpustakaan, mulai dari membaca data buku, mencari buku terfavorit, mengurutkan buku berdasarkan rating, menampilkan 5 buku terbaik, hingga mencari buku berdasarkan rating tertentu.
Pada func CetakTerfavorit, Fungsi ini mencari buku dengan rating tertinggi. Rating buku pertama dijadikan patokan awal, lalu dibandingkan satu per satu dengan buku berikutnya. Jika ada yang lebih tinggi maka patokan diperbarui, setelah selesai data buku terfavorit dicetak.
Func UrutBuku: Fungsi ini mengurutkan buku dari rating tertinggi ke terendah menggunakan insertion sort. Ambil satu buku, bandingkan dengan buku sebelumnya, jika ratingnya lebih tinggi maka buku sebelumnya digeser ke kanan, ulangi sampai ketemu posisi yang tepat.
Func Cetak5Terbaru: Fungsi ini mencetak 5 judul buku dengan rating tertinggi. Jika jumlah buku kurang dari 5 maka semua dicetak, jika lebih maka hanya 5 pertama saja.
Func CariBuku: Fungsi ini mencari buku berdasarkan rating menggunakan binary search. Cek bagian tengah array, jika rating di tengah terlalu kecil maka pencarian ke kiri, jika terlalu besar maka ke kanan, begitu terus sampai ketemu atau tidak ada.
