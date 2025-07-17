package main

import (
	"fmt"
	"slices"
	"strings"
)
func main() {
    pessoa1 := make([]rune, 0)
    pessoa2 := make([]rune, 0)
    var x rune
    var pontuacao, vogais1, vogais2 int
    vogais := []rune{'a', 'e', 'i', 'o', 'u'}
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        pessoa1 = append(pessoa1, x)
    }
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        pessoa2 = append(pessoa2, x)
    }
    pessoa1 = []rune(strings.ToLower(string(pessoa1)))
    pessoa2 = []rune(strings.ToLower(string(pessoa2)))
    for i := range pessoa1 {
        if slices.Contains(vogais, pessoa1[i]) {
            vogais1++
        }
    }
    for i := range pessoa2 {
        if slices.Contains(vogais, pessoa2[i]) {
            vogais2++
        }
    }
    if pessoa1[0] == pessoa2[0] {
        pontuacao += 20
    }
    if len(pessoa1) == len(pessoa2) {
        pontuacao += 30
    }
    if vogais1 == vogais2 {
        pontuacao += 30
    }
    if slices.Contains(vogais, pessoa1[len(pessoa1)-1]) && slices.Contains(vogais, pessoa2[len(pessoa2)-1]) || !slices.Contains(vogais, pessoa1[len(pessoa1)-1]) && !slices.Contains(vogais, pessoa2[len(pessoa2)-1]) {
        pontuacao += 20
    }
    if slices.Contains(vogais, pessoa1[len(pessoa1)-1]) && !slices.Contains(vogais, pessoa2[len(pessoa2)-1]) || !slices.Contains(vogais, pessoa1[len(pessoa1)-1]) && slices.Contains(vogais, pessoa2[len(pessoa2)-1]) {
        pontuacao -= 10
    }
    if pontuacao < 0 {
        pontuacao = 0
    }
    fmt.Printf("As chances do crush te dar bola sao: %v", pontuacao)
    fmt.Println("%!")
}
