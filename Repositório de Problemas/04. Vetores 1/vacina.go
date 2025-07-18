package main

import (
	"fmt"
	"slices"
)
func main() {
    var quantidade int
    fmt.Scan(&quantidade)
    vacina := make([]int, quantidade)
    pessoas := make([]int, quantidade)
    for i := range vacina {
        fmt.Scan(&vacina[i])
    }
    for i := range pessoas {
        fmt.Scan(&pessoas[i])
    }
    slices.Sort(pessoas)
    slices.Reverse(pessoas)
    for i := range vacina {
        var status bool
        for j := range pessoas {
            if pessoas[j] <= vacina[i] {
                status = true
                pessoas[j] = 9999
                break
            }
        }
        if status == false {
            fmt.Println("No")
            return
        }
    }
    fmt.Println("Yes")
}
