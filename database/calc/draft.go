package main

import "fmt"

func main() {
	var x1, x2 int
	var o string
	fmt.Scan(&x1, &x2, &o)
	switch o {
	case "+":
		fmt.Println(x1 + x2)
	case "-":
		fmt.Println(x1 - x2)
	case "*":
		fmt.Println(x1 * x2)
	case "/":
		fmt.Println(x1 / x2)
	}
}
