package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print("[ ")
	for c := b; a != c+1; a++ {
		fmt.Printf("%d %d ", a, b)
		b--
	}
	fmt.Println("]")
}
