package main

import (
	"fmt"
	"math"
)

func main() {

	var arr [100]int
	var n int

	fmt.Print("Masukkan jumlah elemen (max 100 elemen): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array.
	for i := 0; i < n; i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	// b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("arr[%d] = %d\n", i, arr[i])
		}
	}

	// c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indeks ke-0 adalah genap).
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("arr[%d] = %d\n", i, arr[i])
		}
	}

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
	var x int
	fmt.Print("\nMasukkan nilai x untuk kelipatan: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("arr[%d] = %d\n", i, arr[i])
		}
	}

	// e. Menghapus elemen array pada indeks tertentu. Menampilkan data yang tidak dihapus saja.  
	var idx int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n = n - 1
	fmt.Println("Array setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
	total := 0
	for i := 0; i < n; i++ {
		total = total + arr[i]
	}
	rataRata := float64(total) / float64(n)
	fmt.Print(rataRata)

	// g.Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
	mean := float64(total) / float64(n)
	jumlahKuadrat := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - mean
		jumlahKuadrat = jumlahKuadrat + (selisih * selisih)
	}
	standarDeviasi := math.Sqrt(jumlahKuadrat / float64(n))
	fmt.Printf("\nStandar deviasi = %.2f\n", standarDeviasi)

	// h.Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.
	var cari int
	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	frekuensi := 0
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frekuensi = frekuensi + 1
		}
	}
	fmt.Printf("Bilangan %d muncul sebanyak %d kali\n", cari, frekuensi)
}