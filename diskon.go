// Kantin Tel-U berencana memberi semangat kepada mahasiswa yang sedang melaksanakan asesmen CLO dengan memberikan diskon 35% Saat pembelian makan siang dikantin . Buatlah program untuk menghitung besarnya belanja mahasiswa setelah potong diskon

// Masukan terdiri dari bilangan rill yang menyatakan total belanja (uang) awal mahasiswa dan boolean yang menyatakan mahasiswa sedang mengikuti assesmen CLO atau tidak
// Keluaran terdiri dari total belanja akhir yang harus dibayarkan mahasiswa dikasir.

// # Pseudocode

// #program diskon
// kamus
// 	totalBelanja : integer
// 	mengikutiAssesmen : boolean
// algoritma
// 	input(totalBelanja, mengikutiAssesmen)
// 	if mengikutiAssesmen then
// 		totalBelanja = totalBelanja - (totalBelanja * 35 / 100)
// 	end if
// 	output(totalBelanja)
// endprogram

package main

import "fmt"

func main() {
	var totalBelanja int
	var mengikutiAssesmen bool

	fmt.Scanln(&totalBelanja)
	fmt.Scanln(&mengikutiAssesmen)

	if mengikutiAssesmen {
		fmt.Println(totalBelanja - (totalBelanja * 35 / 100))
	} else {
		fmt.Println(totalBelanja)
	}
}