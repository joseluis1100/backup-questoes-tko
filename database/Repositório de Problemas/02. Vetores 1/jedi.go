package main

import "fmt"

func main() {
	var t, sj, ss int
	fmt.Scan(&t)
	j := make([]int, t/2)
	s := make([]int, t/2)
	for c := range j {
		fmt.Scan(&j[c])
		sj += j[c]
	}
	for c := range s {
		fmt.Scan(&s[c])
		ss += s[c]
	}
	if sj == ss {
		fmt.Println("Empate")
	} else if sj > ss {
		fmt.Println("Jedi")
	} else {
		fmt.Println("Sith")
	}
}
