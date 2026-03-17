package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(((a-c)*(a-c)) + ((b-d)* (b-d)))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main(){
	var cx1, cx2, cy1, cy2, r1, r2, x, y float64
	var dalamcircle1, dalamcircle2 bool
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalamcircle1 = didalam(cx1, cy1, r1, x, y)
	dalamcircle2 = didalam(cx2, cy2, r2, x, y)

	if dalamcircle1 && dalamcircle2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if dalamcircle1 {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if dalamcircle2 {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}