package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	v := make([]int, n)
	for c := range v {
		fmt.Scan(&v[c])
	}
	fmt.Print("[")
	for c := range v {
		if c == n-1 {
			fmt.Printf("%d", v[c])
		} else {
			fmt.Printf("%d, ", v[c])
		}
	}
	fmt.Println("]")
}
