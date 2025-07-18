package main

import "fmt"

func main() {
	s := make([]rune, 0, 100)
	for c := 0; c < 100; c++ {
		var x rune
		fmt.Scan(x)
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		s = append(s, x)
	}
	var i, q int
	fmt.Scan(&i, &q)
	for c := i; c < i + q; c++ {
        if c > len(s) - 1 {
            break
        }
		fmt.Printf("%c", s[c])
	}
	fmt.Println()
}
