# <h1 align="center">Laporan Praktikum Modul 9 - Algoritma Pemrograman 2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Modul 9 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### soal1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul9/output/soal1.png)
Penjelasan : 
Program ini digunakan untuk mengecek posisi sebuah titik terhadap dua lingkaran. Pertama, pengguna memasukkan titik pusat dan jari-jari untuk masing-masing lingkaran, lalu memasukkan koordinat titik yang ingin diperiksa.Program menggunakan dua tipe data, yaitu Point untuk menyimpan koordinat titik, dan Circle untuk menyimpan pusat serta jari-jari lingkaran. Fungsi distance dipakai untuk menghitung jarak antara titik dengan pusat lingkaran. Setelah itu, fungsi isInside menentukan apakah titik berada di dalam lingkaran dengan membandingkan jarak tersebut dengan jari-jari.Terakhir, program menampilkan hasilnya, apakah titik berada di dalam kedua lingkaran, hanya di lingkaran pertama, hanya di lingkaran kedua, atau di luar keduanya.

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi
tersebut.

#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

type IsiArrayInt [100]int

func cetakArray(data IsiArrayInt, n int) {
	fmt.Print("Semua elemen : ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func cetakGanjil(data IsiArrayInt, n int) {
	fmt.Print("Indeks ganjil : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func cetakGenap(data IsiArrayInt, n int) {
	fmt.Print("Indeks genap : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func tampilKelipatanX(data IsiArrayInt, n int, x int) {
	fmt.Print("Indeks kelipatan ", x, " : ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndex(data *IsiArrayInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n = *n - 1
}

func hitungRata(data IsiArrayInt, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total = total + data[i]
	}
	rata := float64(total) / float64(n)
	return rata
}

func hitungStdDev(data IsiArrayInt, n int) float64 {
	rata := hitungRata(data, n)
	jumlah := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - rata
		jumlah = jumlah + (selisih * selisih)
	}
	hasil := math.Sqrt(jumlah / float64(n))
	return hasil
}

func hitungFrekuensi(data IsiArrayInt, n int, angka int) int {
	count := 0
	for i := 0; i < n; i++ {
		if data[i] == angka {
			count++
		}
	}
	return count
}

func main() {
	var data IsiArrayInt
	var n int

	fmt.Print("Masukkan jumlah elemen : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Elemen ke-", i, " : ")
		fmt.Scan(&data[i])
	}

	fmt.Println()

	// a
	cetakArray(data, n)

	// b
	cetakGanjil(data, n)

	// c
	cetakGenap(data, n)

	// d
	var x int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	tampilKelipatanX(data, n, x)

	// e
	var idx int
	fmt.Print("Masukkan indeks yang dihapus : ")
	fmt.Scan(&idx)
	hapusIndex(&data, &n, idx)
	fmt.Print("Array setelah dihapus : ")
	cetakArray(data, n)

	// f
	fmt.Printf("Rata-rata : %.2f\n", hitungRata(data, n))

	// g
	fmt.Printf("Standar deviasi : %.2f\n", hitungStdDev(data, n))

	// h
	var cari int
	fmt.Print("Cari frekuensi angka : ")
	fmt.Scan(&cari)
	fmt.Println("Frekuensi", cari, ":", hitungFrekuensi(data, n, cari), "kali")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul9/output/soal2.png)
Penjelasan : 
Program ini digunakan untuk menyimpan sejumlah bilangan bulat ke dalam array, lalu melakukan berbagai proses terhadap data tersebut. Pertama, pengguna memasukkan banyaknya elemen yang akan disimpan, kemudian mengisi setiap elemen satu per satu.
Setelah array terisi, program dapat menampilkan seluruh isi array dengan menelusuri setiap elemen dari indeks pertama hingga terakhir. Program juga dapat menampilkan elemen yang berada pada indeks ganjil maupun genap dengan memeriksa hasil pembagian indeks dengan 2.Selain itu, pengguna dapat memasukkan sebuah nilai x untuk menampilkan elemen-elemen yang berada pada indeks kelipatan x. Program cukup mengecek apakah indeks habis dibagi dengan nilai tersebut.Untuk menghapus elemen, pengguna memasukkan indeks yang diinginkan. Program kemudian menggeser seluruh elemen setelah indeks tersebut satu langkah ke kiri, lalu jumlah elemen dikurangi satu agar data yang dihapus tidak lagi dianggap bagian dari array. Program juga mampu menghitung rata-rata dengan menjumlahkan seluruh elemen lalu membaginya dengan banyak data. Standar deviasi dihitung dengan mencari selisih setiap elemen terhadap rata-rata, menguadratkannya, menjumlahkan semuanya, kemudian membaginya dengan jumlah elemen dan mengambil akar kuadratnya. Terakhir, program dapat menghitung frekuensi suatu angka, yaitu dengan menelusuri seluruh array dan menghitung berapa kali angka tersebut muncul.

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

#### soal3.go

```go
package main

import "fmt"

const NMAX int = 100

type tabelHasil [NMAX]string

func main() {
	var klubA, klubB string
	var hasil tabelHasil
	var n int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	nomor := 1
	for {
		var skorA int
		var skorB int
		fmt.Print("Pertandingan ", nomor, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}
		n++
		nomor++
	}

	for i := 0; i < n; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul9/output/soal3.png)
Penjelasan : 
Program ini digunakan untuk mencatat hasil pertandingan antara dua klub sepak bola. Pertama, pengguna memasukkan nama kedua klub yang akan bertanding. Setelah itu, program akan meminta skor dari masing-masing klub untuk setiap pertandingan secara berulang.
Setiap kali skor dimasukkan, program langsung menentukan hasil pertandingan. Jika skor klub pertama lebih besar, maka nama klub pertama disimpan ke dalam array. Jika skor klub kedua lebih besar, maka nama klub kedua yang disimpan. Sedangkan jika kedua skor sama, program menyimpan tulisan "Draw". Proses input akan terus berlangsung sampai pengguna memasukkan nilai negatif pada salah satu skor. Nilai negatif tersebut berfungsi sebagai tanda bahwa tidak ada pertandingan lagi yang akan dimasukkan. Setelah proses selesai, seluruh hasil pertandingan yang telah tersimpan di dalam array akan ditampilkan satu per satu sesuai urutan pertandingan. Terakhir, program menampilkan pesan bahwa seluruh pertandingan telah selesai.

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)
		if ch == '\n' || ch == ' ' {
			continue
		}
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks  : ")
	isiArray(&tab, &m)

	for i := 0; i < m; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()

	// palindrom
	if palindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}

	// reverse
	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul9/output/soal4.png)
Penjelasan : 
Program ini menyimpan karakter yang dimasukkan pengguna ke dalam array bertipe rune. Input dibaca satu per satu dan akan berhenti saat pengguna mengetik tanda titik (.). Spasi dan enter tidak disimpan. Fungsi palindrom digunakan untuk mengecek apakah teks sama jika dibaca dari depan maupun belakang. Jika sama, maka hasilnya adalah palindrom. Fungsi balikanArray digunakan untuk membalik urutan karakter dalam array dengan menukar karakter dari ujung kiri dan kanan secara bertahap hingga bertemu di tengah.Lalu, program akan menampilkan apakah teks tersebut palindrom serta menampilkan versi terbaliknya.