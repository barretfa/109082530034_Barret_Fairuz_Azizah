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
![Screenshot Output Unguided 1_1](https://github.com/barretfa/109082530034_Barret_Fairuz_Azizah/blob/main/ALPRO2/modul2/output/soal1.png)
Program di atas meminta user untuk memasukkan nilai/value dari satu, dua, dan tiga dalam bentuk string. Setelah itu pada Output awal, program memberikan keluaran/output dari nilai/value yang telah user masukkan tadi. Lalu, pada baris ke-17 sampai 19 diberi aturan dimana value dari temp itu sama dengan value satu, value dari satu sama dengan value dua, dan value dari tiga sama dengan value temp, atau di sini dilakukan pergeseran atau rotation. Jadi pada Output akhir, program akan memberikan output/keluaran berdasarkan peraturan yang sudah ada di atasnya/sebelumnya. 
