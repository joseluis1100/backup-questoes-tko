package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c)
	d = math.Pow(b, 2) - 4*a*c
	if d > 0 {
		fmt.Printf("%.2f\n%.2f\n", (-b+math.Sqrt(d))/(2*a), (-b-math.Sqrt(d))/(2*a))
	} else if d == 0 {
		fmt.Printf("%.2f\n", (-b+math.Sqrt(d))/(2*a))
	} else {
		fmt.Println("nao ha raiz real")
	}
}
