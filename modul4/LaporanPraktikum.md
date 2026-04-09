# <h1 align="center">Laporan Praktikum Modul 4 - Algoritma Pemrograman 2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Modul 4 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Selesaikan program tersebut dengan memanfaatkan prosedure yang diberikan

#### soal1.go

```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i 
	}
}
func permutasi(n, r int, hasil *int) {
	var an, anr int
	faktorial(n, &an)
	faktorial(n-r, &anr)
	*hasil = an / anr
}
func kombinasi(n, r int, hasil *int) {
	var an, ar, anr int
	faktorial(n, &an)
	faktorial(r, &ar)
	faktorial(n-r, &anr)
	*hasil = an / (ar * anr)
}

func main() {
	var a, b, c, d, hasilP, hasilC int
	fmt.Scan(&a, &b, &c, &d)

	permutasi(a, c, &hasilP)
	kombinasi(a, c, &hasilC)
	fmt.Println(hasilP, hasilC)

	permutasi(b, d, &hasilP)
	kombinasi(b, d, &hasilC)
	fmt.Println(hasilP, hasilC)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul4/output/soal1.png)
Penjelasan : 
jadi program ini digunakan untuk menghitung nilai faktorial, permutasi, dan juga kombinasi dengan memanfaatkan konsep prosedur dan juga pointer. func faktorial digunakan untuk ngitung nilai faktorial dari suatu angka. hasil yang sudah didapat tidak dikembalikan pakai return tapi langsung dimasukin ke variabel pakai pointer makanya ada *hasil. terus func permutasi buat ngitung nilai P(n, r). nilai n! dan (n-r)! bakalan disimpan sementara ke dalam variabel (an dan anr) setelah itu baru dihitung hasil permutasinya dengan cara membagi n! dengan (n-r)!. lalu ada func kombinasi mirip dengan func permutasi namun yang membedakan adalah rumusnya. namun func kombinasi membutuhkan tiga nilai yaitu n!, r! dan (n-r)!. jadi program bakalan manggil fungsi faktorial tiga kali.terus hasilnya nnati dipakai buat ngitung kombinasi pakai rumus n!/ (r! x (n-r)!).
program akan membaca 4 angka yang telah di input setelah itu bakal langsung di hitung permutasian dan kombinasinya

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### soal2.go

```go
package main

import "fmt"

func hitungSkor(soal, skor *int) {
	*soal = 0
	*skor = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int
	maxSoal := -1
	minSkor := 999999

	for {
		fmt.Scan(&nama)
		if nama == "selesai" {
			break
		}
		hitungSkor(&soal, &skor)
	

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul4/output/soal2.png)
Penjelasan : 
jadi program ini tuh digunakan untuk menentukan pemenang lomba dari jumlah soal yang berhasil dikerjakan dan total waktu pengerjaannya.
program akan ngebaca 8 waktu pengerjaan soal. kalo waktu ≤ 300, soal dianggap selesai, jadi soal bertambah dan waktunya dijumlah ke skor. Nilai ini langsung mengubah variabel di main karena menggunakan pointer. Di main, program membaca nama peserta satu per satu sampai input "selesai". Setiap peserta dihitung jumlah soal dan total waktunya, lalu dibandingkan. Pemenang adalah yang soalnya paling banyak, dan kalau sama, yang waktunya paling kecil.
