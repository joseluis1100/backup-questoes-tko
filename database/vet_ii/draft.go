package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for c := 0; c != n; c++ {
		fmt.Scan(&l[c])
	}
	fmt.Print("[ ")
	for c := 0; c != n; c++ {
		fmt.Printf("%d ", l[c])
	}
	fmt.Println("]")
}
