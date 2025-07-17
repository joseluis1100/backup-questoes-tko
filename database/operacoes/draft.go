package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("%d\n%d\n%d\n%.2f\n%d\n", a+b, a-b, a*b, float32(a)/float32(b), a%b)
}
