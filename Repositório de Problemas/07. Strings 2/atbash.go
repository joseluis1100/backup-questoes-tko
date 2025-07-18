package main

import (
	"fmt"
	"slices"
)

func main() {
	frase := make([]rune, 0)
	word1 := make([]rune, 0)
	word2 := make([]rune, 0)
	var x rune
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
		word1 = append(word1, x)
	}
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		word2 = append(word2, x)
	}
	for i := range frase {
		if slices.Contains(word1, frase[i]) {
			for j := range word1 {
				if frase[i] == word1[j] {
					frase[i] = word2[j]
					break
				}
			}
		} else if slices.Contains(word2, frase[i]) {
            for j := range word1 {
				if frase[i] == word2[j] {
					frase[i] = word1[j]
					break
				}
			}
        }
	}
    fmt.Println(string(frase))
}
