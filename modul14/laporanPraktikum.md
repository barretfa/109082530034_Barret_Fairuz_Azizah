# <h1 align="center">Laporan Praktikum Modul 10 - Pencarian Nilai Max/Min </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1.Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer 𝒏 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 𝒏 baris berikutnya selalu dimulai dengan sebuah integer 𝒎 (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-masing daerah. Keterangan: Terdapat 3 daerah dalam contoh input, dan di masing-masing daerah mempunyai 5, 6, dan 3 kerabat.

#### soal1.go

```go
//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

type arrRumah [1000000]int

func selectionSort(T *arrRumah, n int) {
	var i, j, idx_min int
	var t int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	rumahkerabat := make([]arrRumah, n)
	semuaM := make([]int, n)

	daerah := 0
	for daerah < n {
		fmt.Scan(&m)
		semuaM[daerah] = m
		k := 0
		for k < m {
			fmt.Scan(&rumahkerabat[daerah][k])
			k = k + 1
		}
		daerah = daerah + 1
	}

	daerah = 0
	for daerah < n {
		selectionSort(&rumahkerabat[daerah], semuaM[daerah])
		k := 0
		for k < semuaM[daerah] {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumahkerabat[daerah][k])
			k = k + 1
		}
		fmt.Println()
		daerah = daerah + 1
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul14/output/hercules1.png)
Program di atas meminta user untuk memasukkan bilangan bulat positif dan riil. Bilangan bulat positif untuk menyatakan ada berapa kelinci yang akan dijual lalu bilangan riil untuk menyatakan berat dari kelinci-kelinci tersebut. Lalu, program akan memunculkan dua buah bilangan riil, bilangan pertama menyatakan berat kelinci yang terkecil dan bilangan kedua menyatakan berat kelinci terbesar. Program ini pertama-tama mengambil berat (isi array) pada indeks ke-0 dan menjadikannya nilai max dan min, dan nilai tersebut dibandingkan dengan nilai lain (isi dari array dengan indeks 1-n) dan ketika nilai lain dibandingkan dengan nilai max/min tadi ternyata nilainya lebih kecil atau lebih besar nilai lain tersebut yang akan menjadi max dan min menggantikan nilai max/min yang sebelumnya (isi darri array dengan indeks 0).

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap. Petunjuk: • Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.• Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.

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

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. Keluaran adalah median yang diminta, satu data per baris. Keterangan: Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median saat itu adalah 11. Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah (11+13)/2=12. Petunjuk: Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.

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
