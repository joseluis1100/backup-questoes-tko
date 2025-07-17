package main

import (
	"fmt"
	"math"
)

func main() {
	var c, a float64
	fmt.Scan(&c, &a)
	fmt.Println(math.Ceil(a / (c - 1)))
}
