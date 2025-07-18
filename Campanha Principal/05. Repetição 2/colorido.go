package main

import "fmt"

func main() {
	var n int
	var p string
	fmt.Scan(&n, &p)
	fmt.Print("[ ")
	for c := 0; c != 11; c++ {
		if c == n {
			continue
		}
		if c == 10 {
			fmt.Print("ceu ")
		} else {
			fmt.Printf("%d%s ", c, p)
		}
		if p == "e" {
			p = "d"
		} else {
			p = "e"
		}
	}
	fmt.Println("]")
}
