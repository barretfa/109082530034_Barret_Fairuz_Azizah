package main

import "fmt"

func f(x int) int {
	return x * x
} 

func g(x int) int {
	return x - 2
} 

func h(x int) int {
	return x + 1
} 

func main (){
	var a, b, c, fogoh, gohof, hofog int
	fmt.Scan(&a, &b, &c)

	fogoh = f(g(h(a)))
	gohof = g(h(f(b)))
	hofog = h(f(g(c)))
	
	fmt.Println(fogoh)
	fmt.Println(gohof)
	fmt.Println(hofog)
}