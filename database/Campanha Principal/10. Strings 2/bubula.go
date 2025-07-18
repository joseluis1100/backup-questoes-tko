package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	frase := make([]rune, 0)
	var x rune
	vogais := []rune{'a', 'e', 'i', 'o', 'u', 225}
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		frase = append(frase, x)
	}
	palavras := strings.Split(string(frase), " ")
	for i := range palavras {
		var contador int
		for k := range palavras[i] {
			if k > 0 && slices.Contains(vogais, rune(palavras[i][k-1])) && !(slices.Contains(vogais, rune(palavras[i][k]))) {
				contador++
			}
		}
		for j := range palavras[i] {
			if contador > 0 {
				if j > 0 && slices.Contains(vogais, rune(palavras[i][j-1])) && !(slices.Contains(vogais, rune(palavras[i][j]))) {
					fmt.Print(palavras[i][:j], palavras[i][:j], palavras[i][:j], palavras[i][j:])
					break
				}
			} else {
				fmt.Print(palavras[i])
				break
			}
		}
		if i == len(palavras)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
