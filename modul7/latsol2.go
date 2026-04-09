//package main

//import "fmt"

//type Waktu struct {
	//jam, menit, detik int
//}

//type DataRental struct {
	//nama         string
	//jenisNaga    string
	//waktuPinjam  Waktu
	//waktuKembali Waktu
//}

//func main() {
	//var durasi Waktu
	//var rental DataRental
	//var Wpinjam, Wkembali, Lpinjam, tarif int

	
//}
package main

import "fmt"

type waktu struct {
	jam, menit, detik int
}

type datarental struct {
	nama, jenis_naga string
	waktu_pinjam, waktu_kembali waktu
}

func main() {
	var durasi waktu
	var rental datarental
	var dPinjam, dKembali, lPinjam, tarif int

	fmt.Println("Masukkan nama, jenis naga, waktu pinjam dan waktu kembali: ")

	fmt.Scanln(&rental.nama)
	fmt.Scanln(&rental.jenis_naga)
	fmt.Scanln(&rental.waktu_pinjam.jam, &rental.waktu_pinjam.menit, &rental.waktu_pinjam.detik)
	fmt.Scanln(&rental.waktu_kembali.jam, &rental.waktu_kembali.menit, &rental.waktu_kembali.detik)

	dPinjam = rental.waktu_pinjam.detik + rental.waktu_pinjam.menit*60 + rental.waktu_pinjam.jam*3600
	dKembali = rental.waktu_kembali.detik + rental.waktu_kembali.menit*60 + rental.waktu_kembali.jam*3600

	lPinjam = dKembali - dPinjam

	if rental.jenis_naga == "Naga_Gembul" {
		tarif = lPinjam * 500
	} else if rental.jenis_naga == "Naga_Cungkring" {
		tarif = lPinjam * 200
	}

	if lPinjam > 3600 {
		tarif = tarif + 15000
	}

	durasi.jam = lPinjam / 3600
	durasi.menit = lPinjam % 3600 / 60
	durasi.detik = lPinjam % 3600 % 60

	fmt.Printf("Yth. %s, anda menyewa %s selama %d jam %d menit %d detik.\n", rental.nama, rental.jenis_naga, durasi.jam, durasi.menit, durasi.detik)
	fmt.Print("Total tagihan + Uang Seblak (jika ada): Rp ", tarif)
}

