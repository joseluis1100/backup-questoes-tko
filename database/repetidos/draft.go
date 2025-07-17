package main

import "fmt"

func main() {
	var x, n, qx int
	fmt.Scan(&x, &n)
	m := make([]int, n)
	for c := range m {
		fmt.Scan(&m[c])
		if m[c] == x {
			qx++
		}
	}
	fmt.Println(qx)
}
