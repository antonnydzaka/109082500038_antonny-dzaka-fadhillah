package main

import "fmt"

func main() {
	var kg, g int
	var Hkg , Hg int 
	var berat int 
	fmt.Print("masukan berat(gram): ")
	fmt.Scan(&berat)
	kg = berat/1000	
	g = berat%1000
	if kg > 10{
		Hkg = kg * 10000
		Hg = g * 0 
	}else{
		Hkg = kg * 10000
		if g>=500{
			Hg = g * 5 
		}else{
			Hg = g * 15
		}
	}
	fmt.Print("===DETAIL PERHITUNGAN===")
	fmt.Printf("detail berat: %vkg + %vg \n",kg,g)
	fmt.Printf("detail pembayaran %v + %v \n",Hkg,Hg)
	fmt.Println("total pembayaran: ",Hkg+Hg)
}