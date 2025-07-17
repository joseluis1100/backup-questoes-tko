package main

import "fmt"

func main() {
	numeros := make([]rune, 0)
	inverso := make([]rune, 0)
	for { 
		var x rune
		fmt.Scanf("%c", &x)
		if (x == '\n'){
			break
		}
		numeros = append(numeros, x)
        inverso = append(inverso, x)
	}
	for c, f := 0, len(numeros)-1; c < f; c, f = c+1, f-1 {
		inverso[f], inverso[c] = numeros[c], numeros[f] 
	}
	if string(numeros) == string(inverso) {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}
