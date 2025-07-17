package main

import "fmt"

func main() {
	var n int
	var media float32
	fmt.Scan(&n)
	x := make([]float32, n)
	for c := range x {
		fmt.Scan(&x[c])
		media += x[c]
	}
	media /= float32(n)
	fmt.Printf("%.1f\n", media)
}
