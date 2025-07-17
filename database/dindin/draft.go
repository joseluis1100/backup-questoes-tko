package main

import "fmt"

func main() {
	var q, qc, ql, qm, qt int
	fmt.Scan(&q)
	s := make([]string, q)
    t := make([]string, q)
	for c := range s {
		fmt.Scan(&s[c], &t[c])
		if s[c] == "c" {
			qc++
		} else {
			ql++
		}
		if t[c] == "m" {
			qm++
		} else {
			qt++
		}
	}
	if qc == ql {
		fmt.Println("empate")
	} else if qc > ql {
		fmt.Println("c")
	} else {
		fmt.Println("l")
	}
	if qm == qt {
		fmt.Println("empate")
	} else if qm < qt {
		fmt.Println("m")
	} else {
		fmt.Println("t")
	}
}
