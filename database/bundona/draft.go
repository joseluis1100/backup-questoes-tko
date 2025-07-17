package main

import "fmt"

func main() {
	var h, m, d, total int
	var s string
	fmt.Scanf("%d\n%d\n%s\n%d", &h, &m, &s, &d)
	if s == "H" {
		total = h*60 + m + 10*d
	} else {
		total = h*60 + m - 10*d
	}
	fmt.Printf("%02d %02d\n", (((total)%720+720)%720)/60, (((total)%720+720)%720)%60)
}
