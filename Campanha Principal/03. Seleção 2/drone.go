package main

import "fmt"

func main() {
	var a, b, c, h, l int
	fmt.Scan(&a, &b, &c, &h, &l)
	if a <= h && b <= l || a <= h && c <= l || b <= h && a <= l || b <= h && c <= l || c <= h && a <= l || c <= h && b <= l {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
