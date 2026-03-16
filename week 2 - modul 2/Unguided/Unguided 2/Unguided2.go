package main

import "fmt"

func main() {
	var merah, kuning, hijau, ungu string
	var hasil bool = true
	for i := 1; i <= 5; i++ {
		fmt.Printf("percobaan %v: ",i)
		fmt.Scan(&merah,&kuning,&hijau,&ungu)
		fmt.Println()
		if merah!="merah" && kuning!="kuning" && hijau!="hijau" && ungu!="ungu"{
			hasil = false
		}
	}
	fmt.Print("hasil: ",hasil)
}