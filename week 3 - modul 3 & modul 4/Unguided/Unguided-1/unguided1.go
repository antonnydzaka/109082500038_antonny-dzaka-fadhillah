package main

import "fmt"

func main() {
	var a,b,c,d int 
	fmt.Scan(&a,&b,&c,&d)
	if a>=c && b<=d{
		p1 :=permutasi(a,c)
		c1 :=combinasi(a,c)
		p2 :=permutasi(b,d)
		c2 :=combinasi(b,d)
		fmt.Print(p1,c1,p2,c2)
	}else{
		fmt.Print("syarat tidak valid")
	}
}

func faktorial(number int)int{
	var hasil int = 1
	for i := 1; i <=number; i++ {
		hasil *= i
	}
	return hasil 
}

func permutasi(n,r int)int {
	return faktorial(n)/faktorial(n-r)
}

func combinasi(n,r int)int{
	return faktorial(n)/(faktorial(r)*faktorial(n-r))
}

