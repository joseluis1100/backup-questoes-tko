package main

import "fmt"

func main() {
	var c1, l1, c2, l2 int
	fmt.Scan(&c1, &l1, &c2, &l2)
	if c1*l1 > c2*l2 {
		fmt.Println(c1 * l1)
	} else {
		fmt.Println(c2 * l2)
	}
}
