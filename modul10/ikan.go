//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var berat [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x+y-1)/y
	var total [1000]float64
	for i := 0; i < x; i++ {
		total[i/y] += berat[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 { fmt.Print(" ") }
		fmt.Printf("%.2f", total[i])
	}
	fmt.Println()

	rataRata := 0.0
	for i := 0; i < jumlahWadah; i++ {
		rataRata += total[i]
	}
	fmt.Printf("%.2f\n", rataRata/float64(jumlahWadah))

}