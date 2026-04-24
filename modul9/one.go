//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	c Titik
	radius int
}
func dalamLingkaran(c Lingkaran, p Titik) bool {
	dx := p.x - c.c.x
	dy := p.y - c.c.y
	return dx*dx + dy*dy <= c.radius*c.radius
}

func main() {
	var Lingkaran [2]Lingkaran
	var p Titik

	fmt.Scan(&Lingkaran[0].c.x, &Lingkaran[0].c.y, &Lingkaran[0].radius)

	fmt.Scan(&Lingkaran[1].c.x, &Lingkaran[1].c.y, &Lingkaran[1].radius)

	fmt.Scan(&p.x, &p.y)

	in1 := dalamLingkaran(Lingkaran[0], p)
	in2 := dalamLingkaran(Lingkaran[1], p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}