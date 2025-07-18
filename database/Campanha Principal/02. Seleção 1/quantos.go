package main

import "fmt"

func comp(x1, x2, x3 int) int {
	if x1 == x2 && x1 == x3 {
		return 3
	} else if x1 != x2 && x1 != x3 && x2 != x3 {
		return 0
	} else {
		return 2
	}
}

func main() {
	var x1, x2, x3 int
	fmt.Scan(&x1, &x2, &x3)
	fmt.Println(comp(x1, x2, x3))
}
