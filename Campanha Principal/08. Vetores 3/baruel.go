package main

import (
	"fmt"
	"slices"
)

func main() {
	var t, q int
	fmt.Scan(&t, &q)
	f := make([]int, q)
	r := make([]int, 0)
	falta := make([]int, 0)
	for c := range f {
		fmt.Scan(&f[c])
		if c > 0 && f[c] == f[c-1] {
			r = append(r, f[c])
		}
	}
	if len(r) == 0 {
		fmt.Print("N")
	} else {
		for c := range r {
			if c == len(r)-1 {
				fmt.Print(r[c])
			} else {
				fmt.Printf("%d ", r[c])
			}
		}
	}
	fmt.Println()
	for c := 1; c < t+1; c++ {
		if !slices.Contains(f, c) {
			falta = append(falta, c)
		}
	}
	if len(falta) == 0 {
		fmt.Print("N")
	} else {
		for c := range falta {
			if c == len(falta)-1 {
				fmt.Print(falta[c])
			} else {
				fmt.Printf("%d ", falta[c])
			}
		}
	}
	fmt.Println()
}
