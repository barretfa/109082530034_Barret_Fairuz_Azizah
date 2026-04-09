package main

import "fmt"

func star(n, i int) {
	if i > n {
		return
	}
	var j int
	for j = 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	star(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	star(n, 1)
}