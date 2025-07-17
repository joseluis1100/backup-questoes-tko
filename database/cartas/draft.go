package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for c := range n {
		fmt.Scan(&x[c])
	}
	fmt.Print("[")
	for c := range n {
		if x[c] == 1 {
			fmt.Print("A")
		} else if x[c] == 11 {
			fmt.Print("J")
		} else if x[c] == 12 {
			fmt.Print("Q")
		} else if x[c] == 13 {
			fmt.Print("K")
		} else {
			fmt.Print(x[c])
		}
		if c != n-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
