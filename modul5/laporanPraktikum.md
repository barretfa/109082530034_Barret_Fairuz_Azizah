# <h1 align="center">Laporan Praktikum Modul 5 - Rekursif </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### soal1.go

```go
package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
		} else if n == 1 {
			return 1
			} else {
				return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	var n, i int
	n = 10

	for i = 0; i <= n; i++ {
		fmt.Printf("%d ", fibo(i))
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_one.png)
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

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_two.png)
Program di atas meminta user untuk memasukkan bilangan bulat positif. Lalu kemudian program akan langsung menjalankan sesuai dengan perintah. Dimana program akan mencetak baris pola bintang dari atas hingga bawah terurut sesuai dengan inputan yang user beri (bil bulat positif tadi). Pola bintang dimulai 1 hingga n (bil bul positif).

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