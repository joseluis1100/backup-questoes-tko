package main

import (
	"fmt"
	"slices"
)
func main() {
    v := make([]int, 5)
    for c := range v {
        fmt.Scan(&v[c])
    }
    slices.Sort(v)
    fmt.Println(v[0])
}
