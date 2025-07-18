package main

import (
	"fmt"
	"slices"
)
func main() {
    num1 := make([]rune, 0)
    num2 := make([]rune, 0)
    var x rune
    for {
        fmt.Scanf("%c", &x)
        if x == ' ' {
            break
        }
        num1 = append(num1, x)
    }
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        num2 = append(num2, x)
    }
    permutacao := true
    if len(num1) != len(num2) {
        permutacao = false
    }
    for i := range num1 {
        if !slices.Contains(num2, num1[i]) {
            permutacao = false
        }
    }
    if permutacao {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }
}
