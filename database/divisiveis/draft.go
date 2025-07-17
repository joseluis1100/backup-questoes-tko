package main

import "fmt"

func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	if x1%3 == 0 && x2%3 == 0 || x1%5 == 0 && x2%5 == 0 {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}
