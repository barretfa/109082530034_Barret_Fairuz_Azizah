//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

type arrRumah [1000000]int

func selectionSort(T *arrRumah, n int) {
	var i, j, idx_min int
	var t int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	rumahkerabat := make([]arrRumah, n)
	semuaM := make([]int, n)

	daerah := 0
	for daerah < n {
		fmt.Scan(&m)
		semuaM[daerah] = m
		k := 0
		for k < m {
			fmt.Scan(&rumahkerabat[daerah][k])
			k = k + 1
		}
		daerah = daerah + 1
	}

	daerah = 0
	for daerah < n {
		selectionSort(&rumahkerabat[daerah], semuaM[daerah])
		k := 0
		for k < semuaM[daerah] {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumahkerabat[daerah][k])
			k = k + 1
		}
		fmt.Println()
		daerah = daerah + 1
	}
}