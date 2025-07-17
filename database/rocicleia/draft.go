package main

import (
	"fmt"
	"slices"
)
func main() {
    palavra1 := make([]rune, 0)
    palavra2 := make([]rune, 0)
    var x rune
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        palavra1 = append(palavra1, x)
    }
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        palavra2 = append(palavra2, x)
    }
    permutacao := true
    if len(palavra1) != len(palavra2) {
        permutacao = false
    }
    for i := range palavra2 {
        if !slices.Contains(palavra1, palavra2[i]) {
            permutacao = false
        }
    }
    if permutacao {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
