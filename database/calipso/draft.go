package main

import (
	"fmt"
	"strings"
)
func main() {
    var quantidade int
    fmt.Scan(&quantidade)
    frases := make([][]rune, quantidade)
    var x rune
    var i int
	for i < quantidade {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			i++
		} else {
			frases[i] = append(frases[i], x)
		}
	}
    for i := range frases {
        qtd_letras := 0
        if frases[i][0] >= 97 && frases[i][0] <= 122 {
            for j := range frases[i] {
                if frases[i][j] != ' ' {
                    if qtd_letras % 2 == 0 {
                        fmt.Print(strings.ToLower(string(frases[i][j])))
                    } else {
                        fmt.Print(strings.ToUpper(string(frases[i][j])))
                    }
                    qtd_letras++
                } else {
                    fmt.Print(" ")
                }
            }
        } else {
            for j := range frases[i] {
                if frases[i][j] != ' ' {
                    if qtd_letras % 2 == 0 {
                        fmt.Print(strings.ToUpper(string(frases[i][j])))
                    } else {
                        fmt.Print(strings.ToLower(string(frases[i][j])))
                    }
                    qtd_letras++
                } else {
                    fmt.Print(" ")
                }
            }
        }
        fmt.Println()
    }
}
