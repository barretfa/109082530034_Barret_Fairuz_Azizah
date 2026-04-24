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
Program di atas adalah program untuk menampilkan deret angka Fibonacci. Program di atas menggunakan fungsi rekursif. Dalam fungsi Fibonacci ada dua kondisi yaitu ketika n = 0 maka akan mengembalikan nilai 0 begitu pula ketika n = 1. Jika n sudah lebih dari 1 maka program akan menghitung fibo(n-1) + fibo(n-2) atau mengembalikan nilai. 

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

#### soal2.go

```go
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

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul5/output/output_dua.png)
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