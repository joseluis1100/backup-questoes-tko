package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	f := make([]int, 0)
	for c := range n {
		if c == 0 || c == 1 {
			f = append(f, 1)
		} else {
			f = append(f, f[c-1]+f[c-2])
		}
	}
	fmt.Println(f[n-1])
}
