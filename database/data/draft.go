package main

import "fmt"

func main() {
	var h, m, d, mm, a int
	fmt.Scan(&h, &m, &d, &mm, &a)
	fmt.Printf("%02d:%02d %02d/%02d/%02d\n", h, m, d, mm, a%100)
}
