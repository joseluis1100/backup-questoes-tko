package main

import "fmt"

func main() {
	var n, d, a int
	fmt.Scan(&n, &d, &a)
	fmt.Println(((d-a)%n + n) % n)
}
