package main

import "fmt"

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64
	fmt.Print("masukan jumlah mahasiswa:")
	fmt.Scan(&jumlah)
	fmt.Print("masukan lama hari:")
	fmt.Scan(&lama)
	fmt.Print("masukan tujuan domestik/macanegara:")
	fmt.Scan(&tujuan)
	perhitungan(jumlah,lama,tujuan,&biaya)
	fmt.Print(biaya)
}

func tanggunganhari(lama int, tujuan string)int{
	if tujuan == "domestik"{
		if lama < 3{
			return lama
		}else{
			return 3
		}
	}else{
		if lama < 8{
			return lama 
		}else{
			return 8
		}
	}
}

func biayaperhari(jumlah int)int{
	totalbiaya := 70+250+300
	return totalbiaya* jumlah
}

func perhitungan(jumlah,lama int, tujuan string,totalbiaya *float64){
	var lamahari int 
	if tujuan == "domestik"{
		lamahari = tanggunganhari(lama,tujuan)
		*totalbiaya = float64(lamahari) * float64(biayaperhari(jumlah))
	}else{
		lamahari = tanggunganhari(lama,tujuan)
		*totalbiaya = float64(lamahari) * (float64(biayaperhari(jumlah)) * 1.5)
	}
}