package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" && jumlahHari > 3 {
		jumlahHari = 3
	}else if tujuan == "mancanegara" && jumlahHari > 8 {
		jumlahHari = 8
	}
	return jumlahHari
}
func biayaPerHari(jumlahMHS int) int {
	return jumlahMHS * (70000 + 300000 + 250000)
}
func perhitunganBiaya(jumlahMHS, lamaperjalanan int, tujuan string, totalBiaya *float64) {
	var lama int = tanggunganHari(lamaperjalanan, tujuan)
	*totalBiaya = float64(biayaPerHari(jumlahMHS))
	*totalBiaya = *totalBiaya * float64(lama)
	if tujuan == "mancanegara" {
		*totalBiaya = 1.5 * *totalBiaya
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Scan(&jumlah, &lama, &tujuan)
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)
	fmt.Printf("Rp %.0f",biaya)
}