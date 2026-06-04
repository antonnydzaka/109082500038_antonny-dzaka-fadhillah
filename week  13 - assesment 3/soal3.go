package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type Partai [NMAX]partai

func posisi(t Partai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var x Partai
	var n int
	var suara int

	fmt.Println("Masukkan proses input suara :")
	
	fmt.Scan(&suara)
	for suara != -1 {
		idx := posisi(x, n, suara)
		if idx != -1 {
			x[idx].suara++
		} else {
			x[n].nama = suara
			x[n].suara = 1
			n++
		}
		fmt.Scan(&suara)
	}

	for i := 1; i < n; i++ {
		temp := x[i]
		j := i - 1
		for j >= 0 && x[j].suara < temp.suara {
			x[j+1] = x[j]
			j--
		}
		x[j+1] = temp
	}

	fmt.Println("\nHasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", x[i].nama, x[i].suara)
	}
	fmt.Println()
}