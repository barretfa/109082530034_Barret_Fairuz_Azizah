//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

func main() {

	var hasil [100]string
	var jumlahHasil int = 0
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1

	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlahHasil] = klubA
		} else if skorB > skorA {
			hasil[jumlahHasil] = klubB
		} else {
			hasil[jumlahHasil] = "Draw"
		}

		jumlahHasil = jumlahHasil + 1
		pertandingan = pertandingan + 1
	}

	for i := 0; i < jumlahHasil; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}