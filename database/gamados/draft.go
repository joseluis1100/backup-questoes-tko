package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
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
    palavras := strings.Split(string(frase), " ")
    original := make([]string, len(palavras))
    copy(original, palavras)
    sort.Strings(palavras)
    if slices.Equal(palavras, original) {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
