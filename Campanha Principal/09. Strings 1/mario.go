package main

import (
	"fmt"
	"slices"
)
func main() {
    var q, x int
    fmt.Scan(&q)
    n := make([]int, 0)
    for range q {
        fmt.Scan(&x)
        n = append(n, x)
    }
    maior := slices.Max(n)
    for l := maior; l > 0; l-- {
        for c := range n {
            if n[c] >= l {
                fmt.Print("#")
            } else {
                fmt.Print("_")
            }
        }
        fmt.Println()
    }
}
