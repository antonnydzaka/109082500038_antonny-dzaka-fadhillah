package main

import "fmt"

type mhs struct {
	nim   int
	nama  string
	nilai int
}

const max = 51

type arrayMhs [max]mhs

func main() {
	var n int
	var mahasiswa arrayMhs
	var nimMhs int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)
	inputarrayMhs(&mahasiswa, n)
	fmt.Print("NIM mahasiswa yang nilai pertama dan terbesarnya: ")
	fmt.Scan(&nimMhs)
	var min , max int = nilaiMinMax(&mahasiswa, nimMhs)
	fmt.Printf("Nilai minimum: %d\n", min)
	fmt.Printf("Nilai maksimum: %d\n", max)

}

func inputarrayMhs(T *arrayMhs, n int) {
	for i := 0; i < n; i++ {
		var nim int
		var nama string
		var nilai int
		fmt.Print("Masukkan data mahasiswa: ")
		fmt.Scan(&nim)
		fmt.Scan(&nama)
		fmt.Scan(&nilai)
		T[i] = mhs{nim, nama, nilai}
	}
}

func nilaiMinMax(T *arrayMhs, nim int)(int, int) {
	var min,max int =T[0].nilai,T[0].nilai
	for i := 0; i < len(T) ; i++ {
		if T[i].nim == nim {
			if T[i].nilai < min {
				min = T[i].nilai
			}
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max , min
}