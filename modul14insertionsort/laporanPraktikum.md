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

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul14/output/hercules1.png)
Program di atas merupakan program untuk mengurutkan nomor rumah kerabat Hercules di setiap daerah secara ascending (dari kecil ke besar) menggunakan algoritma selection sort. Pertama-tama program meminta user menginputkan sebuah bilangan bulat positif (n) yang menyatakan jumlah daerah. Selanjutnya, untuk setiap daerah, program meminta user menginputkan bilangan bulat (m) yang menyatakan jumlah kerabat di daerah tersebut, diikuti dengan m bilangan bulat yang menyatakan nomor rumah masing-masing kerabat. Semua data dari seluruh daerah dikumpulkan terlebih dahulu sebelum diproses. Setelah semua input selesai, program akan mengurutkan nomor rumah di tiap daerah menggunakan fungsi selection sort yang bekerja dengan cara mencari nilai terkecil lalu menukarnya ke posisi paling depan, diulang hingga seluruh elemen terurut dari kecil ke besar. Setelah proses pengurutan selesai, program menampilkan n baris output, di mana setiap baris berisi nomor-nomor rumah kerabat di daerah tersebut yang sudah terurut dari kecil ke besar yang kemudian dipisahkan oleh spasi.

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap. Petunjuk: • Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.• Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.

#### soal2.go

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

		pertama := true
		k := 0
		for k < semuaM[daerah] {
			if rumahkerabat[daerah][k]%2 != 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(rumahkerabat[daerah][k])
				pertama = false
			}
			k = k + 1
		}

		k = semuaM[daerah] - 1
		for k >= 0 {
			if rumahkerabat[daerah][k]%2 == 0 {
				if !pertama {
					fmt.Print(" ")
				}
				fmt.Print(rumahkerabat[daerah][k])
				pertama = false
			}
			k = k - 1
		}

		fmt.Println()
		daerah = daerah + 1
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul14/output/hercules2.png)
Program di atas merupakan program untuk mengurutkan nomor rumah kerabat Hercules di setiap daerah secara ascending (dari kecil ke besar) menggunakan algoritma selection sort, ini hampir mirip dengan nomor satu tetapi bedanya nomor dua ini selain menampilkan secara ascending program ini juga mengurutkan dari yang ganjil dulu baru ke genap. Pertama-tama program meminta user menginputkan sebuah bilangan bulat positif (n) yang menyatakan jumlah daerah. Selanjutnya, untuk setiap daerah, program meminta user menginputkan bilangan bulat (m) yang menyatakan jumlah kerabat di daerah tersebut, diikuti dengan m bilangan bulat yang menyatakan nomor rumah masing-masing kerabat. Semua data dari seluruh daerah dikumpulkan terlebih dahulu sebelum diproses. Setelah semua input selesai, program akan mengurutkan nomor rumah di tiap daerah menggunakan fungsi selection sort yang bekerja dengan cara mencari nilai terkecil lalu menukarnya ke posisi paling depan, diulang hingga seluruh elemen terurut dari kecil ke besar. Setelah proses pengurutan selesai, program menampilkan n baris output dengan urutan khusus, yaitu pertama mencetak semua nomor rumah ganjil dahulu, kemudian dilanjutkan mencetak semua nomor rumah genap, antar nomor rumah dipisahkan oleh spasi. 

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. Keluaran adalah median yang diminta, satu data per baris. Keterangan: Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median saat itu adalah 11. Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah (11+13)/2=12. Petunjuk: Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.

#### soal3.go

```go
//Nama: Barret Fairuz Azizah
//NIM: 109082530034

package main

import "fmt"

type arrData [1000000]int

func insertionSort(T *arrData, n int) {
	var i, j int
	var t int
	i = 1
	for i < n {
		t = T[i]
		j = i - 1
		for j >= 0 && T[j] > t {
			T[j+1] = T[j]
			j = j - 1
		}
		T[j+1] = t
		i = i + 1
	}
}

func cetakMedian(T *arrData, n int) {
	if n%2 != 0 {
		fmt.Println(T[n/2])
	} else {
		fmt.Println((T[n/2-1] + T[n/2]) / 2)
	}
}

func main() {
	var data arrData
	var input, n int

	n = 0
	fmt.Scan(&input)
	for input != -5313 {
		if input == 0 {
			insertionSort(&data, n)
			cetakMedian(&data, n)
		} else {
			data[n] = input
			n = n + 1
		}
		fmt.Scan(&input)
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul14/output/median.png)
Program di atas merupakan program untuk menghitung dan mencetak nilai median dari sekumpulan data bil. bul positif yang diinputkan oleh user. Pertama-tama program membaca input satu per-satu secara terus-menerus hingga menemukan angka -5313 sebagai tanda program berakhir. Setiap angka yang dibaca dan bukan 0 ataupun -5313 akan disimpan ke dalam array data dan counter n bertambah satu. Ketika program menemukan angka 0, program akan mengurutkan seluruh data yang sudah tersimpan menggunakan fungsi insertion sort yang bekerja dengan cara mengambil satu elemen lalu menyisipkannya ke posisi yang tepat di antara elemen-elemen sebelumnya yang sudah terurut, diulang hingga seluruh elemen terurut dari kecil ke besar. Setelah data terurut, fungsi cetakMedian dipanggil untuk mencetak nilai median, di mana jika jumlah data ganjil maka median adalah elemen tepat di tengah array, sedangkan jika jumlah data genap maka median adalah hasil pembagian integer dari penjumlahan dua elemen tengah yang secara otomatis dibulatkan ke bawah. Proses membaca input, menyimpan data, dan mencetak median ini terus berulang hingga program membaca angka -5313 sebagai 'tanda' dari akhir program atau user selesai menginputkan. 