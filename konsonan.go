// Sebuah program digunakan untuk menentukan karakter yang diberikan adalah huruf konsonan atau bukan.

// Masukan terdiri dari satu karakter huruf
// Keluaran terdiri dari teks "konsonan" apabila karakter adalah huruf konsonan atau "bukan konsonan" untuk kemungkinan huruf yang lain

// # Pseudocode

// #program konsonan
// kamus
// 	huruf : string
// algoritma
// 	input(huruf)
// 	if "a" >= huruf AND huruf <= "z" or "A" <= huruf AND huruf <= "Z" then
// 		output("konsonan")
// 	else
// 		output("bukan konsonan")
// 	end if

package main

import "fmt"

func main() {
	var huruf string

	fmt.Scanln(&huruf)

	if "a" >= huruf && huruf <= "z" || "A" <= huruf && huruf <= "Z" {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan konsonan")
	}
}
