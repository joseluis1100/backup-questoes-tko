package main

import "fmt"

func main() {
	var c, s int
	m := [50]int{}
	fmt.Scan(&c)
	temp_c := c
	for v := 0; temp_c > c-c*2; v++ {
		fmt.Scan(&m[v])
		temp_c -= m[v]
		s = v
	}
	temp_c = c
	for v := 0; v != s+1; v++ {
		temp_c -= m[v]
		if temp_c == c {
			fmt.Println("vazio")
		} else if temp_c > 0 {
			fmt.Println("ainda cabe")
		} else if temp_c > c-c*2 {
			fmt.Println("lotado")
		} else {
			fmt.Println("hora de partir")
		}
	}
}
