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
	fmt.Println(soma % 50)
	if soma%50 == 0 {
		fmt.Println(string(nome))
		fmt.Println(0)
		return
	}
	nova_letra := 97
	for i := range 27 {
		if (soma+rune(i)+97)%50 < (soma+rune(nova_letra))%50 {
			nova_letra = 97 + i
		}
	}
	fmt.Printf("%s%c\n", string(nome), nova_letra)
	fmt.Println((soma + rune(nova_letra)) % 50)
}
