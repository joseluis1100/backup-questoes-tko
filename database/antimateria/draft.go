package main

import (
	"fmt"
	"math"
)

func main() {
	var x rune
	var contador int
	palavra1 := make([]rune, 0)
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		palavra1 = append(palavra1, x)
	}
	palavra2 := make([]rune, 0)
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		palavra2 = append(palavra2, x)
	}
	for i := 0; i < int(math.Min(float64(len(palavra1)), float64(len(palavra2)))); i++ {
		if palavra2[i] == palavra1[len(palavra1)-i-1] {
			contador++
		} else {
			break
		}
	}
	fmt.Printf("%v%v\n", string(palavra1[:len(palavra1)-contador]), string(palavra2[contador:]))
}
