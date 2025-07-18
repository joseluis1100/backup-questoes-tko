package main

import "fmt"

func main() {
	var v, t, c float32
	fmt.Scan(&v, &t, &c)
	fmt.Printf("%.2f\n", v*t/60/c)
}
