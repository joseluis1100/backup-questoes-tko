package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)
	s := make([]int, n)
	for c := range s {
		fmt.Scan(&s[c])
	}
	for c := range s {
		if c == 0 {
			if s[c] == 0 && s[c+1] == 0 {
				q++
			}
		} else if c == n-1 {
			if s[c] == 0 && s[c-1] == 0 {
				q++
			}
		} else {
			if s[c] == 0 && s[c+1] == 0 && s[c-1] == 0 {
				q++
			}
		}
	}
	fmt.Println(q)
}
