package main

import "fmt"

func main() {
	var a, g, ra, rg float32
	fmt.Scan(&a, &g, &ra, &rg)
	if ra/a > rg/g {
		fmt.Println("A")
	} else {
		fmt.Println("G")
	}
}
