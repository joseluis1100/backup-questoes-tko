package main

import (
	"fmt"
	"slices"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	j := make([]int, n)
	for c := range j {
		j[c] = c + 1
	}
    x--
    for len(j) > 1 {
        x = (x+1) % len(j)
        j = slices.Delete(j, x, x+1)
        if x == len(j) {
            x = 0
        }
    }
    fmt.Println(j[0])
}
