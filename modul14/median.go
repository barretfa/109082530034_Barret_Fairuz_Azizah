//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

type arrData [1000000]int

func insertionSort(T *arrData, n int) {
	var i, j int
	var t int
	i = 1
	for i < n {
		t = T[i]
		j = i - 1
		for j >= 0 && T[j] > t {
			T[j+1] = T[j]
			j = j - 1
		}
		T[j+1] = t
		i = i + 1
	}
}

func cetakMedian(T *arrData, n int) {
	if n%2 != 0 {
		fmt.Println(T[n/2])
	} else {
		fmt.Println((T[n/2-1] + T[n/2]) / 2)
	}
}

func main() {
	var data arrData
	var input, n int

	n = 0
	fmt.Scan(&input)
	for input != -5313 {
		if input == 0 {
			insertionSort(&data, n)
			cetakMedian(&data, n)
		} else {
			data[n] = input
			n = n + 1
		}
		fmt.Scan(&input)
	}
}