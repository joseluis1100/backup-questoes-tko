package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	fmt.Printf("%d\n%d\n%.2f\n", n1/n2, n1%n2, float32(n1)/float32(n2))
}
