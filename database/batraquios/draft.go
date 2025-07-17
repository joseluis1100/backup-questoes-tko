package main

import (
	"fmt"
	"slices"
)

func main() {
	var nv1, nv2 int
	fmt.Scan(&nv1)
	v1 := make([]int, nv1)
	for c := range v1 {
		fmt.Scan(&v1[c])
	}
	fmt.Scan(&nv2)
	v2 := make([]int, nv2)
	for c := range v2 {
		fmt.Scan(&v2[c])
	}
	for c := range v1 {
		if slices.Contains(v2, v1[c]) {
            if c == len(v1)-1 {
                fmt.Println("sim")
            }
            continue
		} else {
            fmt.Println("nao")
            break
        }
	} 
}
