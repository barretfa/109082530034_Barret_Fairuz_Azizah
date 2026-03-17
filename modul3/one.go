package main

import "fmt"

func factorial(n int) int {
	var facto int = 1
	var i int
	for i = 2 ; i <= n; i++ {
		facto = facto * i
	}
	return facto
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}


func main (){
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b>= d {
	fmt.Println(permutation(a, c), (combination(a, c)))
	fmt.Println(permutation(b, d), (combination(b, d)))
	} else {
		fmt.Println("error, a >= c && b>= d ")
	}
}