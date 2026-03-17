# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n!/(n−r)!, sedangkan C(n, r) = n!/r!(n−r)!

#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	var facto int = 1
	var i int
	for i = 2 ; i <= n; i++ {
		facto = facto * i
	}
	return facto
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}


func main (){
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b>= d {
	fmt.Println(permutation(a, c), (combination(a, c)))
	fmt.Println(permutation(b, d), (combination(b, d)))
	} else {
		fmt.Println("error, a >= c && b>= d ")
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_one.png)
Program di atas meminta user untuk memasukkan empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b≥ d. Kemudian program mulai menghitung permutasi dan kombinasi dengan pasangan bilangan a dan c, b dan d, lalu setelah menghitungnya program akan menampilkan hasilnya menjadi dua baris dengan baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

#### soal2.go

```go
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

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_two.png)
Program di atas meminta user untuk memasukkan tiga buah bilangan bulat a, b, c yang dipisahkan oleh spasi kemudian program akan mulai menghitungnya menggunakan fungsi matematika berikut: f(x) = x ^ 2 , g(x) = x - 2, h(x) =x+1 lalu setelah menghitung, program akan memberikan tiga baris keluaran yang mana baris pertama adalah (fogoh) (a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c).

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### soal3.go

```go
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

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_three.png)
Program di atas meminta user untuk memasukkan beberapa bilangan bulat yang terbagi menjadi 3 baris, baris pertama dan kedua berisi koordinat titik pusat (cx, cy) dan radius dari lingkaran 1 dan lingkaran 2, baris ketiga berisi koordinat titik sembarang. Selanjutnya, program akan menghitung jarak dengan menggunakan rumus √(a-c)² + (b-d)². Kemudian, melalui function  didalam program akan mengecek apakah titik berada di dalam suatu circle (jika jarak <= r). Berdasarkan pengecekan tersebut program akan menampilkan apakah titik berada di dalam lingkaran 1 atau lingkaran 2, atau keduanya, atau di luar kedua lingkaran.