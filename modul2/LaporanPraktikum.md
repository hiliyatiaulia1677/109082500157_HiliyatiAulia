# <h1 align="center">Laporan Praktikum Modul 2 - Review alpro2 </h1>
<p align="center">Hiliyati Aulia - 109082500157</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/output/soal1.png)(https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/modul2/output/soal1.png)
jadi program tersebut melakukan pertukaran urutan dari tiga string yang sudah dimasukkan oleh pengguna. dengan memanfaatkan variabel temp, string pertama berpindah ke posisi terakhir, sedangkan string kedua dan ketiga bergeser ke depan.

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
#### soal2.go

```go
package main

import "fmt"

func main() {
	var t1, t2, t3, t4 string
	var i int
	var hasil bool

	hasil = true
	for i = 1; i <= 5; i++ {
	fmt.Print("Percobaan ", i, ":")
	fmt.Scan(&t1, &t2, &t3, &t4)

	if t1 != "merah" || t2 != "kuning" || t3 != "hijau" || t4 != "ungu" {
	hasil = false
	}
}
fmt.Print("Berhasil: ", hasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/output/soal2.png)
Program diatas digunakan untuk mengecek apakah sebuah rangkaian percobaan praktikum kimia berhasil atau tidak dengan melakukan pengecekan sebanyak 5 kali percobaan dengan urutan warna yang benar yaitu merah kuning hijau dan ungu. Disini saya membuat 4 variabel yang bertipe data string yaitu t1, t2, t3, t4 untuk menyimpan warna pada gelas dari 1 sampai 4. Satu variabel yang bertipe data int yaitu i untuk perulangannya. Dan satu variabel yang bertipe data Boolean yaitu hasil untuk menyimpan apakah seluruh percobaan berhasil atau tidak.
Diawal saya membuat variabel hasil bernilai true untuk mengasumsikan semua percobaan berhasil sampai terbukti. Lalu masuk pada bagian perulangan, perulangan dilakukan sebanyak 5 kali. Pengguna akan memasukkan 4 urutan warna. Lalu pada bagian if program akan mengecek urutan yang telah dimasukkan oleh pengguna jika urutannya salah maka outpunya Adalah false dan sebaliknya.

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal3.go

```go
package main

import "fmt"

func main() {
	var beratParsel int
	var beratKg int
	var sisaGram int
	var biayaPerKg int
	var biayaPerGram int
	var totalBiaya int

	fmt.Print("Masukkan berat parsel (Gram): ")
	fmt.Scan(&beratParsel)

	beratKg = beratParsel / 1000
	sisaGram = beratParsel % 1000
	biayaPerKg = beratKg * 10000

	if beratKg > 10 {
	totalBiaya = 0
	} else if sisaGram < 500 {
	biayaPerGram = sisaGram * 15
	} else {
	biayaPerGram = sisaGram * 5
	}
	
	totalBiaya = biayaPerKg + biayaPerGram
	fmt.Println("Berat Parsel (Gram): ",beratParsel)
	fmt.Println("Detail Berat: ", beratKg,"kg +", sisaGram, "gr")
	fmt.Println("Biaya Kg: ", "Rp.", biayaPerKg, "+ Rp.", biayaPerGram)
	fmt.Println("Total Biaya: ", totalBiaya)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/hiliyatiaulia1677/109082500157_HiliyatiAulia/blob/main/output/soal3.png)
Program digunakan untuk menghitung biaya kirim berdasarkan berat parsel dalam satuan gram. Dalam func main terdapat beberapa variabel yang dideklarasikan dengan tipe data integer, yaitu beratparsel, beratkg, sisagram, biayaperkg,biayapergram, dan totalbiaya. Untuk mendapatkan berat kg maka berat dari parsel harus dibagi dengan 1000 sedangkan sisa gram yang tidak mencapai kilogram akan di modulus dengan 1000. Untuk menentukan biaya perkilogram yaitu dengan mengalikan berat kilogram dengan 10000. Pada if jika berat kilogram lebih dari 10 maka total biaya akan nol. Namun, jika berat tidak lebih dari 10 maka akan diperiksa sisa gramnya. Bila sisa gram kurang dari 500, akan dikalikan dengan 15 per gramnya. Namun, jika lebih dari 500 akan dikalikan dengan 5 per gramnya. Lalu untuk total biaya dihitung dari total biaya kilogram + biaya gram.
