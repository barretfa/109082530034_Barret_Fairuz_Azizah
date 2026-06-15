//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

func main() {
	var angka []int

	fmt.Print("Masukkan bil. bulat: ")
	for {
		var n int
		fmt.Scan(&n)
		if n < 0 {
			break
		}
		angka = append(angka, n)
	}

	for i := 1; i < len(angka); i++ {
		kunci := angka[i]
		j := i - 1

		for j >= 0 && angka[j] > kunci {
			angka[j+1] = angka[j]
			j--
		}
		angka[j+1] = kunci
	}

	for i, v := range angka {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()

	if len(angka) < 2 {
		fmt.Println("Inputkan lebih dari 1 bil.bulat")
		return
	}

	jarakPertama := angka[1] - angka[0]
	berjarak := true

	for i := 2; i < len(angka); i++ {
		if angka[i]-angka[i-1] != jarakPertama {
			berjarak = false
			break
		}
	}

	if berjarak {
		fmt.Printf("Data berjarak %d\n", jarakPertama)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}