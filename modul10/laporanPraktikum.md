# <h1 align="center">Laporan Praktikum Modul 10 - Pencarian Nilai Max/Min </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### soal1.go

```go
//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

func main() {

	var berat[1000]float64
	var n int
	
	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan berat kelinci ke-", i+1, ": ")
		fmt.Scan(&berat[i])
	}

	min, max := berat[0], berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}

	}
	fmt.Printf("Kelinci terkecil: %.2f", min)
	fmt.Printf("\nKelinci terbesar: %.2f", max)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul10/output/outputKelinci.png)
Program di atas meminta user untuk memasukkan bilangan bulat positif dan riil. Bilangan bulat positif untuk menyatakan ada berapa kelinci yang akan dijual lalu bilangan riil untuk menyatakan berat dari kelinci-kelinci tersebut. Lalu, program akan memunculkan dua buah bilangan riil, bilangan pertama menyatakan berat kelinci yang terkecil dan bilangan kedua menyatakan berat kelinci terbesar. Program ini pertama-tama mengambil berat (isi array) pada indeks ke-0 dan menjadikannya nilai max dan min, dan nilai tersebut dibandingkan dengan nilai lain (isi dari array dengan indeks 1-n) dan ketika nilai lain dibandingkan dengan nilai max/min tadi ternyata nilainya lebih kecil atau lebih besar nilai lain tersebut yang akan menjadi max dan min menggantikan nilai max/min yang sebelumnya (isi darri array dengan indeks 0).

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dany. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var berat [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x+y-1)/y
	var total [1000]float64
	for i := 0; i < x; i++ {
		total[i/y] += berat[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 { fmt.Print(" ") }
		fmt.Printf("%.2f", total[i])
	}
	fmt.Println()

	rataRata := 0.0
	for i := 0; i < jumlahWadah; i++ {
		rataRata += total[i]
	}
	fmt.Printf("%.2f\n", rataRata/float64(jumlahWadah))

}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul10/output/outputIkan.png)
Program di atas merupakan program untuk menentukan total berat ikan di setiap wadah serta rata-rata berat per-wadah berdasarkan data yang user inputkan. Pertama-tama program akan meminta user untuk menginputkan dua bilangan bulat positif, yaitu x (jumlah ikan yang akan dijual) dan y (kapasitas ikan per-wadah). Selanjutnya program akan meminta user menginputkan sebanyak x bilangan riil yang menyatakan berat masing-masing ikan. Setelah semua data terkumpul, program akan memasukkan ikan ke dalam wadah secara berurutan, ikan ke-1 sampai ke-y masuk wadah pertama, ikan ke-(y+1) sampai ke-2y masuk wadah kedua, dan seterusnya. Program juga menjumlahkan berat total ikan di setiap wadah. Jumlah wadah yang dibutuhkan dihitung otomatis berdasarkan nilai x dan y. Setelah proses selesai, program akan menampilkan dua baris output, yaitu baris pertama berisi total berat ikan di setiap wadah dan baris kedua berisi rata-rata total berat per-wadah.

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

#### soal3.go

```go
//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			 *bMin = arr[i] }
		if arr[i] > *bMax {
			 *bMax = arr[i] }
	}
}

func rerata(arr arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	var min, max float64
	
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	hitungMinMax(arr, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg", min)
	fmt.Printf("\nBerat balita maksimum: %.2f kg", max)
	fmt.Printf("\nRerata berat balita: %.2f kg", rerata(arr, n))
} 

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul10/output/outputPosyandu.png)
Program di atas merupakan program untuk menentukan berat balita minimum dan maksimum serta rata-rata berat balita berdasarkan data yang user inputkan. Pertama-tama program akan meminta user untuk menginputkan bilangan bulat positif sebagai jumlah dari data atau berat balita yang akan user inputkan (program akan berulang selama n kali sesuai bil.bulat yang user inputkan) kemudian dengan menggunakan percabangan, program akan membandingkan data berat yang user inputkan sebelumnya dan ketika sudah mendapatkan nilai minimum dan maksimum program akan menampilkannya dan kemudian kembali menghitung rata-rata dari berat balita kemudian menampilkannya pada baris terakhir.
