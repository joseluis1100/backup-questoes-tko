package main

import (
	"fmt"
	"math"
)

func primo(x int, i int) bool {
	if i > int(math.Sqrt(float64(x))) {
		return true
	}
	if x%i == 0 {
		return false
	}
	return primo(x, i+1)
}

func main() {
	var x int
	fmt.Scan(&x)
    if primo(x, 2) {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}
