package main

import "fmt"

func main() {
	var v int
	fmt.Scan(&v)
	if v%7 == 0 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
