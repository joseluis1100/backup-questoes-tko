package main

import "fmt"

func main() {
	var d, h int
	fmt.Scan(&d, &h)
	if d > 1 && d < 7 && (h >= 8 && h <= 11 || h >= 14 && h <= 17) || d == 7 && h >= 8 && h <= 11 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
