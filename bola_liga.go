// Buatlah program yang digunakan untuk menentukan jumlah gol terbanyak dan jumlah gol tersedikit dari suatu grup liga sepak bola

// Masukan terdiri dari empat bilangan yang menyatakan jumlah gol yang berhasil dicetak empat tim bola dalam suatu grup
// Keluaran terdiri dari dua nilai yang menyatakan jumlah gol terbanyak dan jumlah gol tersedikit.
// Catatan: Tidak perlu menggunakan perulangan. Bandingkan antara dua bilangan, lakukan secara berulang

// # Pseudocode

// #program bola_liga
// kamus
// 	team1, team2, team3, team4 : integer
// algoritma
// 	input(team1, team2, team3, team4)
// 	if team1 >= team2 AND team1 >= team3 AND team1 >= team4 then
// 		output(team1)
// 	else if team2 >= team1 AND team2 >= team3 AND team2 >= team4 then
// 		output(team2)
// 	else if team3 >= team1 AND team3 >= team2 AND team3 >= team4 then
// 		output(team3)
// 	else if team4 >= team1 AND team4 >= team2 AND team4 >= team3 then
// 		output(team4)

// 	end if
// 	if team1 <= team2 AND team1 <= team3 AND team1 <= team4 then
// 		output(team1)
// 	else if team2 <= team1 AND team2 <= team3 AND team2 <= team4 then
// 		output(team2)
// 	else if team3 <= team1 AND team3 <= team2 AND team3 <= team4 then
// 		output(team3)
// 	else if team4 <= team1 AND team4 <= team2 AND team4 <= team3 then
// 		output(team4)
// 	end if

package main

import "fmt"

func main() {
	var team1 int
	var team2 int
	var team3 int
	var team4 int

	fmt.Scanln(&team1)
	fmt.Scanln(&team2)
	fmt.Scanln(&team3)
	fmt.Scanln(&team4)

	// Pencarian Goal terbanyak
	if team1 >= team2 && team1 >= team3 && team1 >= team4 {
		fmt.Println(team1)
	} else if team2 >= team1 && team2 >= team3 && team2 >= team4 {
		fmt.Println(team2)
	} else if team3 >= team1 && team3 >= team2 && team3 >= team4 {
		fmt.Println(team3)
	} else if team4 >= team1 && team4 >= team2 && team4 >= team3 {
		fmt.Println(team4)
	}

	// Pencarian Goal tersedikit
	if team1 <= team2 && team1 <= team3 && team1 <= team4 {
		fmt.Println(team1)
	} else if team2 <= team1 && team2 <= team3 && team2 <= team4 {
		fmt.Println(team2)
	} else if team3 <= team1 && team3 <= team2 && team3 <= team4 {
		fmt.Println(team3)
	} else if team4 <= team1 && team4 <= team2 && team4 <= team3 {
		fmt.Println(team4)
	}
}