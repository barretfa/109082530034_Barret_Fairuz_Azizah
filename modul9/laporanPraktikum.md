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
//Nama: Barret Fairuz Azizah
//NIM: 109082530034

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

	// h.Menampilkan frekuensi dari suatu bilangan tertentu di dalam array tersebut. 
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
a. Menggunakan perulangan untuk menampilkan keseluruhan array. Perulangan berhenti ketika nilai i (0) kurang dari n (masukkan dari user). Example: user memasukkan nilai n sebesar 4, maka perulangan akan mulai dari 0 hingga 3. 
b. Menggunakan perulangan dan percabangan untuk menampilkan elemen array dengan indeks ganjil. Jika n-nya 4 maka yang akan muncul adalah elemen dari indeks 1 dan 3. 
c. Menggunakan perulangan dan percabangan untuk menampilkan elemen array dengan indeks genap. Jika n-nya 4 maka yang akan muncul adalah elemen dari indeks 0 dan 2.
d. Menggunakan perulangan dan percabangan untuk menampilkan elemen dari indeks kelipatan x yang mana x merupakan inputan dari user. Di sini user diminta memasukkan nilai x. Program lalu mencetak elemen yang indeksnya merupakan kelipatan dari x.
e. Menggunakan perulangan untuk menampilkan elemen pada indeks tertentu saja karena di sini user diminta untuk memasukkan indeks yang ingin dihapus, jadi nanti yang akan ditampilkan hanya elemen dari indeks-indeks yang tidak dihapus. 
f. Menggunakan perulangan untuk menampilkan rata-rata dari elemen/bilangan di dalam array tersebut setelah poin d (setelah ada elemen dari suatu indeks dihapus). Di sini program menjumlahkan semua elemen array ke variabel total, lalu dibagi dengan n. 
g. Menggunakan perulangan untuk menampilkan standar deviasi. Standar deviasi mengukur seberapa jauh nilai-nilai dalam array menyebar dari rata-ratanya. Langkah-langkahnya adalah yang pertama hitung rata-rata (mean) lalu kurangi setiap elemen dengan mean yang telah dihitung tadi kemudian jumlahkan semua hasil kuadrat lalu bagi dengan n. 
h. Menggunakan perulangan dan percabangan untuk mencari frekuensi dari suatu bilangan di dalam array tersebut. User diminta untuk memasukkan bilangan yang ingin dicari lalu program akan mengecek satu per satu isi dari array, setiap kali menemukan bilangan yang sama, variabel frekuensi ditambah 1.

### 3.Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
#### soal3.go

```go
//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

func main() {

	var hasil [100]string
	var jumlahHasil int = 0
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1

	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlahHasil] = klubA
		} else if skorB > skorA {
			hasil[jumlahHasil] = klubB
		} else {
			hasil[jumlahHasil] = "Draw"
		}

		jumlahHasil = jumlahHasil + 1
		pertandingan = pertandingan + 1
	}

	for i := 0; i < jumlahHasil; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul9/output/outputThree.png)
Program akan meminta user untuk memasukkan dua nama club sepak bola yang bertanding, inputan ini dimasukkan ke dalam variable klubA dan klubB. Lalu, program menggunakan perulangan untuk meminta user untuk kembali memasukkan bilangan sebagai skorA dan skorB, program berulang hingga salah satu dari skor tidak valid (bernilai negatif). Lalu, setiap pertandingan akan dibandingkan atau dicek siapa yang skornya lebih tinggi, jika skornya sama maka yang muncul akan "Draw". Setelah program berhenti berulang maka akan muncul tulisan "Pertandingan selesai."