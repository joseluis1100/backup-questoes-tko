package main

import "fmt"

func main() {
	var c1, v float64
	var c2 string
	fmt.Scan(&c1, &c2, &v)
	if c2 == "m" && c1 > v || c2 == "M" && c1 < v {
		fmt.Println("segundo")
	} else {
		fmt.Println("primeiro")
	}
}
