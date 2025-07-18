package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for c := 1; c <= n; c++ {
		for range n - c {
			fmt.Print(".")
		}
        for c2 := range c {
            fmt.Print(n)
            if c2 != c-1 {
                fmt.Print(".")
            }
        }
		for range n - c {
			fmt.Print(".")
		}
		fmt.Println()
	}
}
