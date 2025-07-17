package main

import (
	"fmt"
	"slices"
)
func main() {
    var q, soma1, soma2 int
    fmt.Scan(&q)
    n := make([]int, q)
    for c := range n {
        fmt.Scan(&n[c])
    }
    for c := range n {
        if c < len(n)-1 && n[c] < n [c+1] {
            soma1 += n[c+1] - n[c]
        }
    }
    slices.Reverse(n)
    for c := range n {
        if c < len(n)-1 && n[c] < n [c+1] {
            soma2 += n[c+1] - n[c]
        }
    }
    if soma1 < soma2 {
        fmt.Println(soma1)
    } else {
        fmt.Println(soma2)
    }
}
