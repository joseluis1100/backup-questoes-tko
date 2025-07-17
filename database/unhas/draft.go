package main

import (
	"fmt"
	"math"
)

func main() {
	var n, valor float64
	fmt.Scan(&n)
    c2 := n
	vetor := make([]float64, int(n))
	for c := range vetor {
		fmt.Scan(&vetor[c])
		c2--
        valor += vetor[c]*math.Pow(10, c2)
	}
	fmt.Println(valor)
}
