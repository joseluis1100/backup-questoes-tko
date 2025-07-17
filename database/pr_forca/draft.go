package main

import (
	"math/rand"
	"fmt"
	"slices"
)
func main() {
    palavras := []string{"legal", "teste", "batata"}
    aleatorio := rand.Intn(len(palavras))
    palavra := []rune(palavras[aleatorio])
    palavra_codificada := make([]rune, len(palavra))
    chutes := make([]rune, 0)
    disponiveis := make(map[rune]int)
    var x rune
    for i := range palavra_codificada {
        palavra_codificada[i] = '*'
    }
    letra := 'a'-1
    for range 26 {
        letra++
        disponiveis[letra]++
    }
    for i := 6; i >= 0; i-- {
        fmt.Println("-----------------------------------------------------------------------------------")
        if i == 0 {
            fmt.Printf("Perdeu! a palavra era %v!\n", string(palavra))
            break
        }
        fmt.Printf("Chances: %v\n", i)
        fmt.Printf("Chutes: %c\n", chutes)
        fmt.Print("Disponiveis: [ ")
        for letra := 'a'; letra <= 'z'; letra++ {
            if _, existe := disponiveis[letra]; existe {
                fmt.Printf("%c ", letra)
            }
        }
        fmt.Println("]")
        fmt.Printf("Palavra codificada: %v\n", string(palavra_codificada))
        for {
            fmt.Print("Digitar chute: ")
            fmt.Scanf("%c\n", &x)
            if slices.Contains(chutes, x) {
                fmt.Println("Palavra já usada, tente novamente.")
            } else {
                break
            }
        }
        for i := range len(palavra) {
            if x == palavra[i] {
                palavra_codificada[i] = palavra[i]
            }
        }
        if !slices.Contains(palavra_codificada, '*') {
            fmt.Println("-----------------------------------------------------------------------------------")
            fmt.Printf("A palavra é %v, voce ganhou!\n", string(palavra))
            break
        }
        chutes = append(chutes, x)
        delete(disponiveis, x)
    }
}
