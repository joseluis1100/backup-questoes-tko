package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		frase = append(frase, x)
	}
	senha := make([]rune, 0)
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		senha = append(senha, x)
	}
	frase_min := []rune(strings.ToLower(string(frase)))
	nova_frase := make([]rune, len(frase))
	pos := 0
	for i := range nova_frase {
		if frase_min[i] >= 'a' && frase_min[i] <= 'z' {
			nova_frase[i] = senha[pos%len(senha)]
			pos++
		} else {
			nova_frase[i] = frase[i]
		}
	}
	fmt.Scanf("%c", &x)
	if x == '+' {
		for i := range frase {
			if frase_min[i] >= 'a' && frase_min[i] <= 'z' {
				frase[i] = (frase[i] + nova_frase[i] - 'a' * 2) % 26 + 'a'
			}
		}
	} else {
		for i := range frase {
			if frase_min[i] >= 'a' && frase_min[i] <= 'z' {
				frase[i] = (frase[i] - nova_frase[i] + 26) % 26 + 'a'
			}
		}
	}
	fmt.Println(string(frase))
}
