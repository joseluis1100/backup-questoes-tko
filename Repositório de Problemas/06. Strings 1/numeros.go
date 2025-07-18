package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	numeros := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var soma, linha int
	frase := make([][]rune, 1)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		if x == ' ' {
			linha++
			frase = append(frase, []rune{})
			continue
		}
		frase[linha] = append(frase[linha], x)
	}
	for i := range frase {
		for j := range frase[i] {
			if slices.Contains(numeros, frase[i][j]) {
				soma += (int(frase[i][j]) - 48) * int(math.Pow(10, float64(len(frase[i])-j-1)))
			}
		}
	}
	fmt.Println(soma)
}
