package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for c := 0; c != n; c++ {
		fmt.Scan(&l[c])
	}
	for c := 0; c != n; c++ {
		fmt.Println(l[c])
	}
}
