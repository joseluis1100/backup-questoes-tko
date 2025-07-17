package main

import "fmt"

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	} else {
		return x2
	}
}
func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	fmt.Println(max(x1, x2))
}
