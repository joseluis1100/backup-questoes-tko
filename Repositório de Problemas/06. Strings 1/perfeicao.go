package main

import "fmt"

func main() {
	nome := make([]rune, 0)
	var soma, x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		soma += x
		nome = append(nome, x)
	}
	nova_letra := 50 - (soma+97)%50 + 97
	if nova_letra <= 122 {
		fmt.Printf("%s%c\n", string(nome), nova_letra)
	} else {
		fmt.Println("sem sorte")
	}
}
