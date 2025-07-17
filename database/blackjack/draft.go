package main

import "fmt"

func main() {
	var q, soma int
	fmt.Scan(&q)
	n := make([]int, q)
	for c := range n {
		fmt.Scan(&n[c])
		if n[c] > 10 {
			n[c] = 10
		}
		if soma < 11 && n[c] == 1 {
			n[c] = 11
		}
		soma += n[c]
		if soma > 21 {
			for c2 := range n {
				if n[c2] == 11 {
					n[c2] = 1
					soma -= 10
				}
			}
		}
	}
	fmt.Println(soma)
}
