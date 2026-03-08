package main

import "fmt"

func main() {
	var tahun int
	var hasil bool 
	fmt.Print("masukan tahun: ")
	fmt.Scan(&tahun)
	if tahun%400==0 || (tahun%4==0 && tahun%100!=0){
		hasil = true
	}else{
		hasil = false
	}
	fmt.Print("kabisat: ",hasil)
}