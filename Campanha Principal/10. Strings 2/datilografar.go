package main

import (
	"fmt"
	"slices"
)

func main() {
	var remover int
	fmt.Scan(&remover)
	numero := make([]rune, 0)
	novo_numero := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		numero = append(numero, x)
	}
	for i := range numero {
		if int(numero[i])-48 != remover {
			novo_numero = append(novo_numero, numero[i])
		}
	}
	for range novo_numero {
		if novo_numero[0] == '0' {
			novo_numero = slices.Delete(novo_numero, 0, 1)
		} else {
			break
		}
	}
	if len(novo_numero) == 0 {
		fmt.Println("0")
	} else {
		fmt.Println(string(novo_numero))
	}
}
