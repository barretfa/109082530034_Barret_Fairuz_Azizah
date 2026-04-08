# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n!/(n−r)!, sedangkan C(n, r) = n!/r!(n−r)!

#### soal1.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul4/output/one_output.png)
Program di atas meminta user untuk memasukkan empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b≥ d. Kemudian program mulai menghitung permutasi dan kombinasi dengan pasangan bilangan a dan c, b dan d, lalu setelah menghitungnya program akan menampilkan hasilnya menjadi dua baris dengan baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)

#### soal2.go

```go
package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var winner string
	var jmlSoal, jmlSkor int

	for {
		var nama string
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > jmlSoal || (soal == jmlSoal && (jmlSkor == 0 || skor < jmlSkor)) {
			winner = nama
			jmlSoal = soal
			jmlSkor = skor
		}
	}

	fmt.Println(winner, jmlSoal, jmlSkor)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul4/output/two_output.png)
Program di atas meminta user untuk memasukkan tiga buah bilangan bulat a, b, c yang dipisahkan oleh spasi kemudian program akan mulai menghitungnya menggunakan fungsi matematika berikut: f(x) = x ^ 2 , g(x) = x - 2, h(x) =x+1 lalu setelah menghitung, program akan memberikan tiga baris keluaran yang mana baris pertama adalah (fogoh) (a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c).
