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

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_satu.png)
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

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_dua.png)
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

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul3/output/output_tiga.png)
Program di atas merupakan program untuk menemukan bilangan yang menjadi faktor dari suatu bilangan yang diinputkan oleh user. Pertama-tama user akan diminta untuk memasukkan bilangan bulat lalu program akan memprosesnya dan akan menampilkan hasilnya. 