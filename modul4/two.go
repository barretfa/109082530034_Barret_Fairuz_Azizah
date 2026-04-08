package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		var waktu int
		fmt.Scan(&waktu)
		
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var winner string
	var jmlSoal, jmlSkor int

	for {
		var nama string
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > jmlSoal || (soal == jmlSoal && (jmlSkor == 0 || skor < jmlSkor)) {
			winner = nama
			jmlSoal = soal
			jmlSkor = skor
		}
	}

	fmt.Println(winner, jmlSoal, jmlSkor)
}