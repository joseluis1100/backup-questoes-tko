package main

import (
	"fmt"
	"slices"
)

func main() {
	frase := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		frase = append(frase, x)
	}
	var c int
	for c < len(frase) {
		if c > 0 && c < len(frase)-1 && frase[c] == ' ' && frase[c-1] == frase[c+1] {
			frase = slices.Delete(frase, c, c+2)
			c = 0
		} else {
            c++
        }
	}
	fmt.Println(string(frase))
}
