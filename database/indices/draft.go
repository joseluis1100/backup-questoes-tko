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
	nn := make([]int, 0)
	on := make([]int, q)
	copy(on, n)
	slices.Sort(on)
	for c := range on {
		for c2 := range n {
			if on[c] == n[c2] {
				nn = append(nn, c2)
				break
			}
		}
	}
	fmt.Print("[ ")
	for c := range nn {
		fmt.Printf("%d ", nn[c])
	}
	fmt.Println("]")
}
