//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import"fmt"

func main(){

var berat, biayaKg, tambahanBiaya, tambahanBiaya1, tambahanBiaya2, totalBiaya, a1, a2 int

fmt.Print("Berat parsel (gram): ")
fmt.Scan(&berat)

a1 = berat / 1000
a2 = berat % 1000
biayaKg = a1 * 10000


	if a1 > 10 && a2 > 500 {
		tambahanBiaya = a2 * 5
		totalBiaya = biayaKg 
		fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya)
		fmt.Printf("Total biaya: %d\n", totalBiaya)
	} else if a1 > 10 && a2 < 500 {
		tambahanBiaya = a2 * 15
		totalBiaya = biayaKg 
		fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya)
		fmt.Printf("Total biaya: %d\n", totalBiaya)
	} else if a2 >= 500 {
		tambahanBiaya1 = a2 * 5
		totalBiaya = biayaKg + tambahanBiaya1
			fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
			fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya1)
			fmt.Printf("Total biaya: %d\n", totalBiaya)
	} else {
		tambahanBiaya2 = a2 * 15
		totalBiaya = biayaKg + tambahanBiaya2
		fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya2)
		fmt.Printf("Total biaya: %d\n", totalBiaya)
	}
}