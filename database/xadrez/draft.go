package main

import "fmt"

func main() {
	var l, c int
	fmt.Scan(&l, &c)
	if l%2 == 0 && c%2 != 0 || l%2 != 0 && c%2 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
