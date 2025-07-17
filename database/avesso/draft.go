package main

import "fmt"

func main() {
	var q, n, x int
	fmt.Scan(&q)
	for range q {
		fmt.Scan(&n, &x)
		op := make([]int, n)
		for c := range op {
			fmt.Scan(&op[c])
		}
		for c := range op {
			if op[c] == x {
				if c == 0 {
					op[c+1] *= -1
				} else if c == len(op)-1 {
					op[c-1] *= -1
				} else {
					op[c+1] *= -1
					op[c-1] *= -1
				}
			}
		}
        fmt.Println(op)
	}
}
