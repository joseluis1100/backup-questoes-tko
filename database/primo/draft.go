package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 2 {
		fmt.Println(1)
	} else {
		for c := 2; c < n; c++ {
			if n%c != 0 {
				if n == 2 || c == n-1 {
					fmt.Println(1)
				}
			} else {
				fmt.Println(0)
				break
			}
		}
	}
}
