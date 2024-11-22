// Buatlah program yang akan menentukan apakah sebuah bilangan adalah kelipatan 3 atau kelipatan 5 atau kelipatan 3 dan juga kelipatan 5

// Masukan adalah sebuah bilangan bulat positif.
// Keluaran berupa string yang menyatakan jenis bilangan tersebut, "Kelipatan 5", "Kelipatan 3" dan jika bukan kelipatan 3 atau 5, maka tidak menampilkan apapun

// # Pseudocode

// #program tigadanlima
// kamus
// 	bilangan : integer
// algoritma
// 	input(bilangan)
// 	if bilangan mod 3 == 0 AND bilangan mod 5 == 0 then
// 		output("Kelipatan 3")
// 		output("Kelipatan 5")
// 	else if bilangan mod 3 == 0 then
// 		output("Kelipatan 3")
// 	else if bilangan mod 5 == 0 then
// 		output("Kelipatan 5")
// 	end if

package main

import "fmt"

func main() {
	var bilangan int

	fmt.Scanln(&bilangan)

	if bilangan % 3 == 0 && bilangan % 5 == 0 {
		fmt.Println("Kelipatan 3")
		fmt.Println("Kelipatan 5")
	} else if bilangan % 3 == 0 {
		fmt.Println("Kelipatan 3")
	} else if bilangan % 5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}