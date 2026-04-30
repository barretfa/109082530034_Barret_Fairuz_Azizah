//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			 *bMin = arr[i] }
		if arr[i] > *bMax {
			 *bMax = arr[i] }
	}
}

func rerata(arr arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	var min, max float64
	
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	hitungMinMax(arr, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg", min)
	fmt.Printf("\nBerat balita maksimum: %.2f kg", max)
	fmt.Printf("\nRerata berat balita: %.2f kg", rerata(arr, n))
} 