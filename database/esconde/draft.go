package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for c := 1; c != n+1; c++ {
		if c%2 != 0 {
			fmt.Println(c)
		}
	}
	for c := -1; n != c; n-- {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}
