package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for c := range x {
		fmt.Scan(&x[c])
	}
	for c := range x {
		if c > 0 && x[c-1] > x[c] {
			x[c], x[c-1] = x[c-1], x[c]
		}
	}
	for c := range slices.Max(x) + 1 {
		if slices.Contains(x, c) && c == slices.Max(x) {
			fmt.Println(c)
		} else if slices.Contains(x, c) {
			fmt.Printf("%d ", c)
		}
	}
}
