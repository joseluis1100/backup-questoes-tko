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
		if letras[c] == 97 || letras[c] == 101 || letras[c] == 105 || letras[c] == 111 || letras[c] == 117 {
			fmt.Printf("%c", letras[c])
		}
	}
	fmt.Println()
	for c := range letras {
		if letras[c] < 97 || letras[c] > 122 || letras[c] == 97 || letras[c] == 101 || letras[c] == 105 || letras[c] == 111 || letras[c] == 117 {
			continue
		} else {
			fmt.Printf("%c", letras[c])
		}
	}
	fmt.Println()
}
