package main

import (
	"fmt"
	"slices"
)

func main() {
	var quantidade, numero int
    fmt.Scan(&quantidade)
    unico := make([]int, 0)
    mapa := make(map[int]int)
    for range quantidade {
        fmt.Scan(&numero)
        if _, existe := mapa[numero]; !existe {
            unico = append(unico, numero)
        }
        mapa[numero]++
    }
    slices.Sort(unico)
    for i := range unico {
        fmt.Print(unico[i])
        if i == len(unico)-1 {
            fmt.Println()
        } else {
            fmt.Print(" ")
        }
    }
}
