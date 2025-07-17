package main

import "fmt"

func main() {
	letras := make([]rune, 0)
	var x rune
	for x != '\n' {
		fmt.Scanf("%c", &x)
		letras = append(letras, x)
	}
	for c := range letras {
		if letras[c] >= 65 && letras[c] <= 92 {
			letras[c] += 32
		} else if letras[c] >= 97 && letras[c] <= 122 {
			letras[c] -= 32
		}
		fmt.Printf("%c", letras[c])
	}
}
