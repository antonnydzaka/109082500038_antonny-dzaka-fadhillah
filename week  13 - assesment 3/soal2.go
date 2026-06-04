package main

import "fmt"

type pemainBola struct {
	name   string
	goal   int
	assist int
}

func main() {
	var data []pemainBola
	var n int
	fmt.Scan(&n)
	input(&data, n)
	selection(data)
	fmt.Println("hasil :")
	for _, vel := range data {
		fmt.Println(vel.name, vel.goal, vel.assist)
	}
}

func input(data *[]pemainBola, n int) {
	for i := 0; i < n; i++ {
		var p pemainBola
		fmt.Scan(&p.name, &p.goal, &p.assist)
		*data = append(*data, p)
	}
}

func selection(data []pemainBola) {
	for i := 0; i < len(data)-1; i++ {
		max := i
		for j := i + 1; j < len(data); j++ {
			if data[j].goal > data[max].goal {
				max = j
			} else if data[j].goal == data[max].goal {
				if data[j].assist > data[max].assist {
					max = j
				}
			}
		}
		data[i], data[max] = data[max], data[i]
	}
}
