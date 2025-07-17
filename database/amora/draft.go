package main

import (
	"fmt"
)

func main() {
	frase := make([]rune, 0)
	trecho := make([]rune, 0)
    var x rune
	contador := 0
	for {
		fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
		frase = append(frase, x)
	}
	for {
		fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
		trecho = append(trecho, x)
	}
	for i := range frase {
		if string(frase[i:i+len(trecho)]) == string(trecho) {
			contador++;
		}
	}
	fmt.Println(contador)
}
