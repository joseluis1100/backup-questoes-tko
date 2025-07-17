package main

import "fmt"

func main() {
	var n, q, a, b, agua int
	fmt.Scan(&n, &q)
	casas := make([]int, n)
	for range q {
		fmt.Scan(&a, &b, &agua)
		for i := a; i <= b; i++ {
			casas[i] += agua
		}
	}
	for i := range casas {
		fmt.Print(casas[i])
		if i == len(casas)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
