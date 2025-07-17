package main

import (
	"fmt"
	"math"
)

func main() {
	var c, b, g, m float64
	fmt.Scan(&c, &b, &g, &m)
	fmt.Println(math.Ceil((b + g + m) / c))
}
