package main

import (
	"fmt"
	"slices"
)
func main() {
    frase := make([][]rune, 0)
	frase = append(frase, []rune{})
	var linha int
    var x rune
    status := "int"
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
        for j := range frase[i] {
            if frase[i][j] >= 65 && frase[i][j] <= 122 {
                status = "str"
                break
            } else if slices.Contains(frase[i], '.') {
                status = "float"
            } else {
                status = "int"
            }
        }
        fmt.Print(status)
        if i == len(frase)-1 {
            fmt.Println()
            return
        }
        fmt.Print(" ")
    }
}
