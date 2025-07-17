package main

import (
	"fmt"
	"math"
)

func main() {
	var n, menor int
	fmt.Scan(&n)
	a := make([]float64, n)
	b := make([]float64, n)
	d := make([]float64, n)
	for c := range a {
		fmt.Scan(&a[c], &b[c])
		if a[c] < 10 || b[c] < 10 {
			d[c] = 101
		} else {
			d[c] = math.Abs(a[c] - b[c])
		}
		if c == 0 {
			menor = c
		} else if d[c] < d[c-1] {
			menor = c
		}
	}
	if d[menor] > 100 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(menor)
	}
}
