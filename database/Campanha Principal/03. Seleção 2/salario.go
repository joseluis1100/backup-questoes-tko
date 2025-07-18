package main

import "fmt"

func main() {
	var s float32
	fmt.Scan(&s)
	if s <= 1000 {
		fmt.Printf("%.2f\n", s*1.2)
	} else if s <= 1500 {
		fmt.Printf("%.2f\n", s*1.15)
	} else if s <= 2000 {
		fmt.Printf("%.2f\n", s*1.1)
	} else if s > 2000 {
		fmt.Printf("%.2f\n", s*1.05)
	}
}
