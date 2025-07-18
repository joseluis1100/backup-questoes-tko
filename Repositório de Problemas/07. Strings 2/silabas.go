package main

import (
	"fmt"
	"slices"
	"strings"
)
func main() {
    linha := make([]rune, 0)
    vogais := []rune{'a', 'e', 'i', 'o', 'u'}
    var x rune
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        linha = append(linha, x)
    }
    linha_min := []rune(strings.ToLower(string(linha)))
    for i := range linha {
        fmt.Printf("%c", linha[i])
        if i < len(linha)-1 && slices.Contains(vogais, linha_min[i]) && !slices.Contains(vogais, linha_min[i+1]) && linha_min[i+1] != ' ' {
            fmt.Print("-")
        }
    }
    fmt.Println()
}
