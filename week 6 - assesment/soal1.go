package main

import (
	"fmt"
)

const pi float64 = 3.14

func main() {
	var r float64
	var tkiri, tkanan float64
	var mjkiri, mjkanan float64
	fmt.Print("masukan jari jari:")
	fmt.Scan(&r)
	fmt.Print("masukan tinggi kiri:")
	fmt.Scan(&tkiri)
	fmt.Print("masukan masa jenis kiri:")
	fmt.Scan(&mjkiri)
	fmt.Print("masukan tinggi kanan:")
	fmt.Scan(&tkanan)
	fmt.Print("masukan masa jenis kanan:")
	fmt.Scan(&mjkanan)
	var massakiri, massakanan float64 = massa(r,tkiri,mjkiri), massa(r,tkanan,mjkanan)
	display(massakiri,massakanan)
}

func volume(r ,t float64)float64{
	volume := r*t*pi
	return volume
}

func massa (r,t,mj float64)float64{
	massa:= volume(r,t)* mj
	return massa
}

func display(m1,m2 float64){
	if m1==m2{
		fmt.Print("BALANCE")
	}else{
		hasil := m1-m2
		if hasil > 0{
			fmt.Print("selisih kiri dan kanan: ",hasil)
		}else{
			fmt.Print("selisih kiri dan kanan: ",hasil*-1)
		}
	}

}