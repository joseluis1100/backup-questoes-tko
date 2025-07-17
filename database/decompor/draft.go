package main

import (
	"fmt"
)
func main() {
    var n int
    fmt.Scan(&n)
    numeros := make([]int, 0)
    for n > 0 {
        numeros = append(numeros, n%10)
        n /= 10
    }
    for c := len(numeros)-1; c >= 0; c-- {
        fmt.Print(numeros[c])
        if c != 0 {
            fmt.Print(" ")
        } else {
            fmt.Println()
        }
    }
}
