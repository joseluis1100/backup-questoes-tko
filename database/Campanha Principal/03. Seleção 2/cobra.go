package main

import "fmt"

func main() {
	var n, x, y, s int
	var c string
	fmt.Scan(&n, &x, &y, &c, &s)
	switch c {
	case "U":
		fmt.Println(x, ((y-s)%n+n)%n)
	case "D":
		fmt.Println(x, (y+s)%n)
	case "L":
		fmt.Println(((x-s)%n+n)%n, y)
	case "R":
		fmt.Println((x+s)%n, y)
	}
}
