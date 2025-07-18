package main

import "fmt"

func main() {
	var p, s, e int
	fmt.Scan(&p, &s, &e)
	for l := 0; l < p; l -= e {
		if l+s >= p {
			fmt.Println(l, "saiu")
			break
		} else if l+s < 0 {
			fmt.Println(l, "morreu")
			break
		} else {
			fmt.Println(l, l+s)
		}
		l += s
		if s > 0 {
			s -= 10
		}
	}
}
