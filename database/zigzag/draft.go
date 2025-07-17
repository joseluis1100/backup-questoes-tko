package main

import "fmt"

func main() {
	var i, f int
	fmt.Scan(&i, &f)
	for ; i != f+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("zigzag")
		} else if i%3 == 0 {
			fmt.Println("zig")
		} else if i%5 == 0 {
			fmt.Println("zag")
		} else {
			fmt.Println(i)
		}
	}
}
