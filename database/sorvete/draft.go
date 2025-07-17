package main

import (
	"fmt"
	"slices"
)
func main() {
    frase := make([][]rune, 1)
    letras := make(map[rune]int)
    var x rune
    var linha, comum int
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        if x == ' ' {
            frase = append(frase, []rune{})
            linha++
            continue
        }
        letras[x]++
        frase[linha] = append(frase[linha], x)
    }
    for letra := range letras {
        existe := true
        for palavra := range frase {
            if !slices.Contains(frase[palavra], letra) {
                existe = false
            }
        }
        if existe == true {
            comum++
        }
    }
    fmt.Println(comum)
}
