package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, x4 float64
	fmt.Scan(&x1, &x2, &x3, &x4)
	fmt.Println(math.Max(math.Max(x1, x2), math.Max(x3, x4)))
}
