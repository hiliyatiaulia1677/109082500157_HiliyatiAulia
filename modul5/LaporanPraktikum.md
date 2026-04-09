# <h1 align="center">Laporan Praktikum Modul 5 - Algoritma Pemrograman 2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Modul 5

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul5/output/soal1.png)
Penjelasan : 
jadi Program ini digunakan untuk menampilkan deret Fibonacci sesuai jumlah yang dimasukkan. fungsi fibonacci(a) bekerja secara rekursif, yang dimana fungsi tersebut memanggil dirinya sendiri. Jika nilai a adalah 0 maka hasilnya 0, jika 1 maka hasilnya 1. Selain itu, nilai Fibonacci dihitung dari penjumlahan dua nilai sebelumnya, yaitu fibonacci(a-1) + fibonacci(a-2).program meminta user memasukkan jumlah suku yang ingin ditampilkan. Setelah itu, program melakukan perulangan dari 0 sampai nilai yang dimasukkan, lalu setiap angka akan dihitung menggunakan fungsi fibonacci dan langsung ditampilkan. Jadi hasil akhirnya adalah deret Fibonacci seperti 0 1 1 2 3 5 dan seterusnya sesuai jumlah yang diminta.

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul5/output/soal2.png)
Penjelasan : 
jadi program ini digunakan untuk menampilkan pola bintang berbentuk segitiga menggunakan rekursif. Func cetakBintang(N) bertugas mencetak bintang dalam satu baris sebanyak N dengan memanggil dirinya sendiri sampai N menjadi 0, lalu saat kembali dari rekursi, bintang dicetak satu per satu sehingga jumlahnya sesuai. Func pola(N, i) mengatur jumlah baris yang akan ditampilkan. Fungsi ini akan mencetak baris pertama dengan 1 bintang, lalu memanggil dirinya lagi untuk mencetak baris berikutnya dengan jumlah bintang yang bertambah sampai mencapai N. program meminta input N dari user, lalu memanggil fungsi pola mulai dari baris pertama, sehingga hasil akhirnya adalah pola bintang yang bertambah dari 1 sampai N.

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

#### soal3.go

```go
package main
import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul5/output/soal3.png)
Penjelasan :
Program ini dibuat buat nampilin semua faktor dari sebuah bilangan N pakai cara rekursif. Jadi, func faktor(n, i) bakal ngecek angka satu per satu mulai dari 1 sampai N. Kalau angka i bisa ngebagi N tanpa sisa, berarti itu faktor, jadi langsung ditampilin. Setiap kali selesai ngecek, fungsi bakal manggil dirinya sendiri lagi dengan nilai i yang ditambah 1. Proses ini terus jalan sampai i lebih besar dari N, baru berhenti. Di bagian main, user diminta masukin nilai N, terus program langsung manggil func faktor mulai dari 1, jadi nanti keluar semua faktor dari N secara urut.