package main

import (
	"fmt"
	"slices"
)

func main() {
	n := make([]int, 5)
	for c := range n {
		fmt.Scan(&n[c])
	}
	fmt.Println(slices.Max(n) + slices.Min(n))
}
