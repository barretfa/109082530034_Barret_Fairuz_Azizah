//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 1; i <= *n; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	idxTerfavorit := 1
	for i := 2; i <= n; i++ {
		if pustaka[i].rating > pustaka[idxTerfavorit].rating {
			idxTerfavorit = i
		}
	}
	b := pustaka[idxTerfavorit]
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul		: %s\n", b.judul)
	fmt.Printf("Penulis		: %s\n", b.penulis)
	fmt.Printf("Penerbit	: %s\n", b.penerbit)
	fmt.Printf("Tahun		: %d\n", b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		kunci := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < kunci.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = kunci
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < 5 {
		batas = n
	}
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 1; i <= batas; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = tengah
			break
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[ketemu]
		fmt.Println("Buku Ditemukan:")
		fmt.Printf("Judul		: %s\n", b.judul)
		fmt.Printf("Penulis		: %s\n", b.penulis)
		fmt.Printf("Penerbit	: %s\n", b.penerbit)
		fmt.Printf("Tahun		: %d\n", b.tahun)
		fmt.Printf("Eksemplar	: %d\n", b.eksemplar)
		fmt.Printf("Rating		: %d\n", b.rating)
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var ratingCari int

	DaftarkanBuku(&pustaka, &nPustaka)

	fmt.Scan(&ratingCari)

	CetakTerfavorit(pustaka, nPustaka)
	fmt.Println()

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)
	fmt.Println()

	CariBuku(pustaka, nPustaka, ratingCari)
}