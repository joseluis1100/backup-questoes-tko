package main

import (
	"fmt"
	"math"
)

func main() {
	var c1, c2, q, p int
	fmt.Scan(&c1, &c2, &q)
	a := make([]string, q)
	for c := range a {
		fmt.Scan(&a[c])
		if a[c] == "v" || a[c] == "c" {
			p += 4
		} else {
			p += 2
		}
	}
	fmt.Println(p)
	if math.Abs(float64(c1)-float64(p)) == math.Abs(float64(c2)-float64(p)) {
		fmt.Println("empate")
	} else if math.Abs(float64(c1)-float64(p)) < math.Abs(float64(c2)-float64(p)) {
		fmt.Println("Chico Bento")
	} else {
		fmt.Println("Cebolinha")
	}
}
