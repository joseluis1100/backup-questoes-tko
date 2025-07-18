package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	n := make([]float64, q)
	for c := range n {
		fmt.Scan(&n[c])
	}
	if len(n)%2 == 0 {
        fmt.Printf("%.1f\n", (float64(n[(len(n)/2)-1]) + float64(n[len(n)/2])) / 2)
	} else {
		fmt.Printf("%.1f\n", float64(n[(len(n)/2)]))
	}
}
