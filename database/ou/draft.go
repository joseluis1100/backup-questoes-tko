package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x == 3 || x == 5 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
