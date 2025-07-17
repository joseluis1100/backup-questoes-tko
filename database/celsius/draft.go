package main

import "fmt"

func main() {
	var c float64
	fmt.Scan(&c)
	fmt.Printf("%.6f\n", 1.8*c+32)
}
