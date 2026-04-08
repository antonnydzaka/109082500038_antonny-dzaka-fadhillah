package main

import "fmt"

func main() {
	var x, y int
	var hasil int = 1
	fmt.Scan(&x,&y)
	pangkat(x,y,hasil)
}

func pangkat(x, y, hasil int){
	if y > 0{
		hasil = x * hasil
		pangkat(x,y-1,hasil)
	}else{
		fmt.Print(hasil)
	}
}
