package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print("[ ")
	for c := 0; c != 11; c++ {
		if c == n {
			continue
		} else if c == 10 {
			fmt.Print("ceu ")
		} else {
			fmt.Printf("%d ", c)
		}
	}
	fmt.Println("]")
}
