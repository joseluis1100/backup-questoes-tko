package main

import "fmt"

func main() {
	var x1, x2 float32
	fmt.Scan(&x1, &x2)
	fmt.Printf("%.1f\n", (x1+x2)/2)
}
