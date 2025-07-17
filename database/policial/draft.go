package main

import (
	"fmt"
	"slices"
)
func main() {
    var q int
    fmt.Scan(&q)
    n := make([]int, q)
    for c := range n {
        fmt.Scan(&n[c])
    }
    slices.Sort(n)
    for c := range n {
        fmt.Print(n[c])
        if c != len(n)-1 {
            fmt.Print(" ")
        } else {
            fmt.Println()
        }
    }
}
