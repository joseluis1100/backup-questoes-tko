package main

import "fmt"

func main() {
	var fn, q int
	fmt.Scan(&fn, &q)
	for c := 0; c != q; c++ {
		fmt.Println(fn)
		fn += 2
	}
}
