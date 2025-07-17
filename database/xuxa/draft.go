package main

import "fmt"

func main() {
	letras := make([]rune, 0, 100)
	for range 100{ 
		var x rune
		fmt.Scanf("%c", &x)
		if (x == '\n'){
			break
		}
		letras = append(letras, x)

	}
	for c, f := 0, len(letras)-1; c < f; c, f = c+1, f-1 {
		letras[c], letras[f] = letras[f], letras[c]
	}
	fmt.Println(string(letras))
}
