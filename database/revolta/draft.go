package main

import "fmt"

func main() {
	var n, sp, si int
	fmt.Scan(&n)
	x := make([]int, n)
	for c := range n {
		fmt.Scan(&x[c])
		if x[c]%2 == 0 {
			sp += x[c]
		} else {
			si += x[c]
		}
	}
	if sp == si {
		fmt.Println("empate")
	} else if sp < si {
		fmt.Println("soldados")
	} else {
		fmt.Println("rebeldes")
	}
}
