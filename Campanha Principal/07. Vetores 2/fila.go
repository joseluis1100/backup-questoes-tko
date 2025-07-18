package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for c := range x {
		fmt.Scan(&x[c])
	}
	fmt.Print("[ ")
	for c := range x {
		if x[c]%2 != 0 {
			fmt.Printf("%d ", x[c])
		}
	}
	fmt.Println("]")
	fmt.Print("[ ")
	for c := range x {
		if x[c]%2 == 0 {
			fmt.Printf("%d ", x[c])
		}
	}
	fmt.Println("]")
}
