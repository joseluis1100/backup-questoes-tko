package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, p float64
	fmt.Scan(&a, &b, &c)
	p = (a + b + c) / 2
	fmt.Printf("%.2f\n", math.Sqrt(p*(p-a)*(p-b)*(p-c)))
}
