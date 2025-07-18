package main

import (
	"fmt"
	"math"
)

func main() {
	var m, a, b float64
	fmt.Scan(&m, &a, &b)
	fmt.Println(math.Max(math.Max(m-(a+b), a), b))
}
