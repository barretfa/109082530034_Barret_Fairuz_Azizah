package main

import "fmt"

func factorial(n int, facto *int) {
	*facto = 1
	for i := 2 ; i <= n; i++ {
		*facto = *facto * i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		var hasil int

		permutation(a, c, &hasil)
		fmt.Print(hasil)
		combination(a, c, &hasil)
		fmt.Println(" ", hasil)

		permutation(b, d, &hasil)
		fmt.Print(hasil)
		combination(b, d, &hasil)
		fmt.Print(" ", hasil)

	} else {
		fmt.Println(" ")
	}
}