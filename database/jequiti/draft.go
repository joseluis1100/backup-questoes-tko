package main

import (
	"fmt"
	"slices"
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
    chutes := make([]rune, 0)
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        chutes = append(chutes, x)
    }
    fmt.Scanf("%c", &x)
    nova_frase := []rune(strings.ToLower(string(frase)))
    for i := range nova_frase {
        if !slices.Contains(chutes, nova_frase[i]) && nova_frase[i] >= 'a' && nova_frase[i] <= 'z' {
            frase[i] = x
        }
    }
    fmt.Println(string(frase))
}
