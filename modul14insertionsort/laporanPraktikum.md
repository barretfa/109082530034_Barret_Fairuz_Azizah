# <h1 align="center">Laporan Praktikum Modul 14 - Sorting </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1.Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal1.go

```go
//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

func main() {
	var angka []int

	fmt.Print("Masukkan bil. bulat: ")
	for {
		var n int
		fmt.Scan(&n)
		if n < 0 {
			break
		}
		angka = append(angka, n)
	}

	for i := 1; i < len(angka); i++ {
		kunci := angka[i]
		j := i - 1

		for j >= 0 && angka[j] > kunci {
			angka[j+1] = angka[j]
			j--
		}
		angka[j+1] = kunci
	}

	for i, v := range angka {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()

	if len(angka) < 2 {
		fmt.Println("Inputkan lebih dari 1 bil.bulat")
		return
	}

	jarakPertama := angka[1] - angka[0]
	berjarak := true

	for i := 2; i < len(angka); i++ {
		if angka[i]-angka[i-1] != jarakPertama {
			berjarak = false
			break
		}
	}

	if berjarak {
		fmt.Printf("Data berjarak %d\n", jarakPertama)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul14insertionsort/output/dataBerjarak.png)
Program di atas adalah program untuk mengurutkan bilangan bulat yang diinputkan oleh user menggunakan metode insertion sort, lalu mengecek apakah jarak antar bilangan setelah diurutkan selalu sama atau tidak. Pertama, program meminta user memasukkan bilangan satu per satu. Selama bilangan yang dimasukkan bernilai 0 atau lebih, bilangan tersebut akan disimpan ke dalam array. Jika user memasukkan bilangan negatif, program berhenti membaca input dan bilangan negatif itu tidak ikut disimpan. Setelah input selesai, program mengurutkan bilangan-bilangan tersebut dari kecil ke besar menggunakan insertion sort. Setelah terurut, program menampilkan semua bilangan dalam satu baris. Kemudian program mengecek jarak antar bilangan yang berdekatan, misalnya jarak antara bilangan pertama dan kedua dijadikan patokan. Jika semua jarak antar bilangan yang berdekatan sama dengan patokan tersebut, maka program menampilkan "Data berjarak x" dengan x adalah nilai jaraknya. Jika ada salah satu jarak yang berbeda, program menampilkan "Data berjarak tidak tetap".

### 2. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini: Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari. Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir. Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan.

#### soal2.go

```go
//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 1; i <= *n; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	idxTerfavorit := 1
	for i := 2; i <= n; i++ {
		if pustaka[i].rating > pustaka[idxTerfavorit].rating {
			idxTerfavorit = i
		}
	}
	b := pustaka[idxTerfavorit]
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul		: %s\n", b.judul)
	fmt.Printf("Penulis		: %s\n", b.penulis)
	fmt.Printf("Penerbit	: %s\n", b.penerbit)
	fmt.Printf("Tahun		: %d\n", b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		kunci := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < kunci.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = kunci
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < 5 {
		batas = n
	}
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 1; i <= batas; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = tengah
			break
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[ketemu]
		fmt.Println("Buku Ditemukan:")
		fmt.Printf("Judul		: %s\n", b.judul)
		fmt.Printf("Penulis		: %s\n", b.penulis)
		fmt.Printf("Penerbit	: %s\n", b.penerbit)
		fmt.Printf("Tahun		: %d\n", b.tahun)
		fmt.Printf("Eksemplar	: %d\n", b.eksemplar)
		fmt.Printf("Rating		: %d\n", b.rating)
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var ratingCari int

	DaftarkanBuku(&pustaka, &nPustaka)

	fmt.Scan(&ratingCari)

	CetakTerfavorit(pustaka, nPustaka)
	fmt.Println()

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)
	fmt.Println()

	CariBuku(pustaka, nPustaka, ratingCari)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul14insertionsort/output/buku.png)
Program di atas adalah program untuk mengelola data buku di sebuah perpustakaan. Pertama, user memasukkan jumlah buku beserta data tiap buku seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating, lalu di baris terakhir user memasukkan angka rating yang ingin dicari. Setelah semua data tersimpan, program langsung mencari dan menampilkan buku dengan rating tertinggi sebagai buku terfavorit. Kemudian program mengurutkan semua buku dari rating tertinggi ke terendah menggunakan insertion sort, lalu menampilkan 5 buku teratas hasil pengurutan tersebut. Terakhir, program mencari buku berdasarkan rating yang diinputkan user menggunakan pencarian biner, jika ditemukan maka seluruh data buku ditampilkan, jika tidak ditemukan program menampilkan pesan "Tidak ada buku dengan rating seperti itu". 