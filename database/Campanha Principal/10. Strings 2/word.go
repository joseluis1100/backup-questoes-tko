package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	frase := make([]rune, 0)
	var x, formatacao rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		frase = append(frase, x)
	}
	fmt.Scanf("%c", &formatacao)
	if formatacao == 'M' {
		fmt.Println(strings.ToUpper(string(frase)))
	} else if formatacao == 'm' {
		fmt.Println(strings.ToLower(string(frase)))
	} else if formatacao == 'p' {
		palavras := strings.Split(string(frase), " ")
		for i := range palavras {
			if len(palavras[i]) == 1 {
				palavras[i] = strings.ToLower(palavras[i])
			} else {
				palavras[i] = strings.Title(palavras[i])
			}
			fmt.Print(palavras[i])
			if i == len(palavras) - 1 {
				fmt.Println()
				return
			}
			fmt.Print(" ")
		}
	} else if formatacao == 'i' {
		for i := range frase {
			if unicode.IsUpper(frase[i]) {
				fmt.Print(strings.ToLower(string(frase[i])))
			} else {
				fmt.Print(strings.ToUpper(string(frase[i])))
			}
		}
		fmt.Println()
	}
}
