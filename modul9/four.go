//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	var karakter rune
	fmt.Scanf("%c", &karakter)
	for karakter != '.' {
		if karakter != '\n' && karakter != ' ' {
			t[*n] = karakter
			*n = *n + 1
		}
		fmt.Scanf("%c", &karakter)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks		: ")
	isiArray(&tab, &m)


	balikanArray(&tab, m)

	fmt.Print("Reverse teks	: ")
	cetakArray(tab, m)

	fmt.Print("Palindrom	: ")
	fmt.Println(palindrom(tab, m))
}