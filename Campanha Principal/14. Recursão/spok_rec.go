package main

import "fmt"

var num_invertido = make([]rune, 0)

func inverso(numero []rune) []rune {
	pos := len(numero) - len(num_invertido) - 1
	if len(num_invertido) < len(numero) {
        num_invertido = append(num_invertido, numero[pos])
		inverso(numero)
	}
	return num_invertido
}

func main() {
	numero := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		numero = append(numero, x)
	}
	inverso(numero)
	if string(numero) == string(num_invertido) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
