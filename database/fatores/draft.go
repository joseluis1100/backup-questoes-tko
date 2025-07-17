package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	d := 2
	c := 0
	for n != 1 {
		if n%d == 0 {
			n /= d
			c++
		} else {
			if c > 0 {
				fmt.Println(d, c)
			}
			d++
			c = 0
		}
	}
	fmt.Println(d, c)
}
