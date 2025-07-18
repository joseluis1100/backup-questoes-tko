package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for c := range x {
		fmt.Scan(&x[c])
	}
	for i, f := 0, n-1; i < f; i, f = i+1, f-1 {
		x[i], x[f] = x[f], x[i]
	}
	fmt.Print("[ ")
	for c := range x {
		fmt.Printf("%d ", x[c])
	}
	fmt.Println("]")
}
