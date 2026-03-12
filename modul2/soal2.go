//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

func main(){
	var m, k, h, u string
	var i int
	var berhasil bool

	berhasil = true

	for i = 1 ; i <= 5 ; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&m, &k, &h, &u)
			if m != "merah" || k != "kuning" || h != "hijau" || u != "ungu" {
				berhasil = false
			}
		}
	fmt.Print("BERHASIL: ", berhasil)
}