package main

import "fmt"

func main() {
	var q, soma float64
	fmt.Scan(&q)
	s := make([]float64, int(q))
	for c := range s {
		fmt.Scan(&s[c])
		soma += s[c]
	}
	media := soma / q
	fmt.Printf("%.2f\n", media)
	for c := range s {
		if s[c] == media {
			fmt.Print("M")
		} else if s[c] > media {
			fmt.Print("G")
		} else {
			fmt.Print("P")
		}
		if c != len(s)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
