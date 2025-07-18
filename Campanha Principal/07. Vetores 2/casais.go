package main

import (
	"fmt"
	"math"
)

func main() {
	var n, casais int
	fmt.Scan(&n)
	m := make([]float64, 0)
	f := make([]float64, 0)
	x := make([]int, n)
	for c := range x {
		fmt.Scan(&x[c])
		if x[c] > 0 {
			m = append(m, float64(x[c]))
		} else {
			f = append(f, math.Abs(float64(x[c])))
		}
	}

	for c := range m {
		for c2 := range f {
			if m[c] == f[c2] {
				casais++
				f[c2] = 0
				break
			}
		}
	}
	fmt.Println(casais)
}
