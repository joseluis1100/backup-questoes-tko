package main

import (
	"fmt"
	"math"
)

func main() {
	var c string
	var x float64
	fmt.Scan(&c, &x)
	switch c {
	case "r":
		fmt.Println(math.Round(x))
	case "f":
		fmt.Println(math.Floor(x))
	case "c":
		fmt.Println(math.Ceil(x))
	}
}
