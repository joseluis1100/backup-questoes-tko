package main

import "fmt"

func main() {
	codigo := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		codigo = append(codigo, x)
	}
	frase := make([][]rune, 0)
	frase = append(frase, []rune{})
	var linha int
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		if x == ' ' {
			frase = append(frase, []rune{})
			linha++
		} else {
			frase[linha] = append(frase[linha], x)
		}
	}
	for i := range frase {
		var contador int
		for j := range frase[i] {
			for k := range codigo {
				if frase[i][j] == codigo[k] {
					contador++
					break
				}
			}
		}
		tamanho := len(frase[i])
		if contador == tamanho {
			fmt.Print("chefe")
		} else if contador > tamanho/2 {
			fmt.Print("ultron")
		} else {
			fmt.Print("pessoa")
		}
		if i == len(frase)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
