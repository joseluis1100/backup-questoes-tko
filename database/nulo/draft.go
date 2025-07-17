package main

import "fmt"

func main() {
	var v int
	fmt.Scan(&v)
	if v > 0 {
		fmt.Println("+")
	} else if v < 0 {
		fmt.Println("-")
	} else {
		fmt.Println("0")
	}
}
