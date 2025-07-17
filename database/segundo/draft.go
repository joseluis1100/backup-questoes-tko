package main

import "fmt"

func main() {
	var h, m, s int
	fmt.Scan(&h, &m, &s)
	s++
	if s == 60 {
		s = 0
		m++
	}
	if m == 60 {
		m = 0
		h++
	}
	if h == 24 {
		h = 0
	}
	fmt.Printf("%02d %02d %02d\n", h, m, s)
}
