# <h1 align="center">Laporan Praktikum Modul 1 - Review Alpro 1 </h1>

<p align="center">Barret Fairuz Azizah - 109082530034</p>

## Unguided

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

#### soal1.go

```go
package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul2/output/soal1.png)
Program di atas meminta user untuk memasukkan nilai/value dari satu, dua, dan tiga dalam bentuk string. Setelah itu pada Output awal, program memberikan keluaran/output dari nilai/value yang telah user masukkan tadi. Di sini diberi aturan dimana value dari temp itu sama dengan value satu, value dari satu sama dengan value dua, dan value dari tiga sama dengan value temp, atau di sini dilakukan pergeseran atau rotation. Jadi pada Output akhir, program akan memberikan output/keluaran berdasarkan peraturan yang sudah ada di atasnya/sebelumnya.

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.

#### soal2.go

```go
//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import "fmt"

func main(){
	var m, k, h, u string
	var i int
	var berhasil bool

	berhasil = true

	for i = 1 ; i <= 5 ; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&m, &k, &h, &u)
			if m != "merah" || k != "kuning" || h != "hijau" || u != "ungu" {
				berhasil = false
			}
		}
	fmt.Print("BERHASIL: ", berhasil)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul2/output/soal2.png)
Program di atas merupakan program untuk menampilkan suatu status “BERHASIL” apakah “true” atau “false”. Kondisi tersebut tergantung dari inputan yang user inputkan. Status “BERHASIL” akan “true” ketika user menginputkan warna dari ke 4 tabung reaksi secara berurutan yaitu ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’sebanyak 5 kali percobaan. Jika salah satu saja tidak sesuai urutan maka status “BERHASIL” akan menjadi “false”.

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

#### soal3.go

```go
//Nama	: Barret Fairuz Azizah
//NIM	: 109082530034

package main

import"fmt"

func main(){

var berat, biayaKg, tambahanBiaya, tambahanBiaya1, tambahanBiaya2, totalBiaya, a1, a2 int

fmt.Print("Berat parsel (gram): ")
fmt.Scan(&berat)

a1 = berat / 1000
a2 = berat % 1000
biayaKg = a1 * 10000


	if a1 > 10 && a2 > 500 {
		tambahanBiaya = a2 * 5
		totalBiaya = biayaKg
		fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya)
		fmt.Printf("Total biaya: %d\n", totalBiaya)
	} else if a1 > 10 && a2 < 500 {
		tambahanBiaya = a2 * 15
		totalBiaya = biayaKg
		fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya)
		fmt.Printf("Total biaya: %d\n", totalBiaya)
	} else if a2 >= 500 {
		tambahanBiaya1 = a2 * 5
		totalBiaya = biayaKg + tambahanBiaya1
			fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
			fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya1)
			fmt.Printf("Total biaya: %d\n", totalBiaya)
	} else {
		tambahanBiaya2 = a2 * 15
		totalBiaya = biayaKg + tambahanBiaya2
		fmt.Printf("Detail berat: %d kg + %d gr\n", a1, a2)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahanBiaya2)
		fmt.Printf("Total biaya: %d\n", totalBiaya)
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_3](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/modul2/output/soal3.png)
Program di atas merupakan program untuk menghitung sekaligus menampilkan detail berat parsel berdasarkan berat parsel dalam gram yang diinputkan oleh user yang kemudian program harus menghitung dan menampilkannya dalam bentuk kilogram (kg) dan sisanya dalam bentuk gram (gr). Selanjutnya, program harus menghitung dan menampilkan detail biaya dimana setiap 1 kg paket dihargai sebesar 10k (Rp. 10.000) dan sisa berat dalam gram tadi dikenakan tambahan biaya dengan ketentuan jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Lalu program akan menghitung dan menampilkan total biayanya, dan ada ketentuan tambahan, jika total berat paket > 10 kg dan sisa berat (yang kurang dari 1kg) maka tambahan biaya akan digratiskan.