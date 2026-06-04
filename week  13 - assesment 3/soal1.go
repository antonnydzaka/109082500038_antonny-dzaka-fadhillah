package main

import "fmt"

func main() {
	var angka []int
	var median []int
	var masukan int = 1
	for {
		fmt.Scan(&masukan)
		if masukan < 0 {
			break
		}
		if masukan == 0 {
			selectionSort(angka)
			m := findMedian(angka)
			median = append(median, int(m))
			fmt.Println("Median:", m)
		} else {
			angka = append(angka, masukan)
		}
	}

}

func findMedian(data []int) float64 {
	n := len(data)
	if n%2 == 1 {
		return float64(data[n/2])
	}
	return float64(data[n/2-1]+data[n/2]) / 2.0
}

func selectionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}
