// Sejumlah N Mahasiswa Tel-U akan mengadakan Tamasya ke Danau Kembar, Alahan Panjang. Buatlah program yang digunakan untuk menentukan jumlah mobil yang diperlukan dari Bandara menuju lokasi wisata, apabila kapasitas satu mobil adalah 7 orang.

// Masukan terdiri dari satu bilangan bulat positif N.
// Keluaran terdiri dari sebuah bilangan yang menyatakan jumlah mobil yang diperlukkan

// # Pseudocode

// #program tamasya
// kamus
// 	n : integer
// 	jumlah_mobil : integer
// algoritma
// 	input(n)
// 	jumlah_mobil = n div 7
// 	if n mod 7 != 0 then
// 		jumlah_mobil++
// 	end if
// 	output(jumlah_mobil)
// endprogram

package main

import "fmt"


func main() {
	var n int
	var jumlah_mobil int

	fmt.Scanln(&n)
	jumlah_mobil = n / 7
	if n % 7 != 0 {
		jumlah_mobil++
	}
	fmt.Println(jumlah_mobil)
}