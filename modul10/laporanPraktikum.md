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

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer). Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

#### soal2.go

```go
package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var winner string
	var jmlSoal, jmlSkor int

	for {
		var nama string
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > jmlSoal || (soal == jmlSoal && (jmlSkor == 0 || skor < jmlSkor)) {
			winner = nama
			jmlSoal = soal
			jmlSkor = skor
		}
	}

	fmt.Println(winner, jmlSoal, jmlSkor)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul4/output/two_output.png)
Program di atas merupakan program untuk menentukan suatu pemenang dari perlombaan berdasarkan total waktunya, yang menjadi pemenang adalah peserta dengan jumlah total waktu yang paling sedikit dibanding dengan peserta lainnya. User diminta untuk menginputkan nama dan ada 8 jumlah waktu lalu user bisa memasukkan nama dan jumlah waktu lagi di baris bagian  bawah (catatan: peserta yang menginputkan 301 pada waktunya berarti tidak mengirimkan jawaban) dan jika sudah maka user harus menginputkan Selesai agar program tidak akan berulang dan program akan langsung menghitung atau memproses lalu program akan menampilkan Nama dari pemenang lalu jumlah soal yang dikerjakan dan juga jumlah total waktu pengerjaannya. 
