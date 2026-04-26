# <h1 align="center">Laporan Praktikum Modul 7 & 9 - Array </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### soal1.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul9/output/outputOne.png)
Program di atas meminta user untuk memasukkan beberapa bilangan bulat yang terbagi menjadi 3 baris, baris pertama dan kedua berisi koordinat titik pusat (cx, cy) dan radius dari lingkaran 1 dan lingkaran 2, baris ketiga berisi koordinat titik sembarang. Selanjutnya, program akan menghitung jarak dengan menggunakan rumus √(a-c)² + (b-d)². Kemudian, melalui program akan mengecek apakah titik berada di dalam suatu circle (jika jarak <= r). Berdasarkan pengecekan tersebut program akan menampilkan apakah titik berada di dalam lingkaran 1 atau lingkaran 2, atau keduanya, atau di luar kedua lingkaran.

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
### a. Menampilkan keseluruhan isi dari array.
### b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
### c. enampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
### d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
### e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
### f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
### g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

#### soal2.go

```go
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

	// c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
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

	// e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil. 
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
	fmt.Printf("Standar deviasi = %.2f\n", standarDeviasi)

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
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul9/output/outputTwo.png)
Program di atas meminta user untuk memasukkan bilangan bulat positif. Lalu kemudian program akan langsung menjalankan sesuai dengan perintah. Dimana program akan mencetak baris pola bintang dari atas hingga bawah terurut sesuai dengan inputan yang user beri (bil bulat positif tadi). Pola bintang dimulai 1 hingga n (bil bul positif).

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}

	if n % i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modu5/output/output_tiga.png)
Program di atas merupakan program untuk menemukan bilangan yang menjadi faktor dari suatu bilangan yang diinputkan oleh user. Pertama-tama user akan diminta untuk memasukkan bilangan bulat lalu program akan memprosesnya dan akan menampilkan hasilnya. 