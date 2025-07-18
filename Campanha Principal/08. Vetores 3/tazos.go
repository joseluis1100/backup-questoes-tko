package main

import "fmt"

func main() {
	var n, maiorf int
	fmt.Scan(&n)
	t := make([]int, n)
	for c := range t {
		fmt.Scan(&t[c])
	}
	maisf := make([]int, 0)
	maisf = append(maisf, t[0])
	for c := range t[len(t)-1] + 1 {
		contador := 0
		for c2 := range t {
			if t[c2] == c {
				contador++
			}
		}
		if contador > maiorf {
			maiorf = contador
			maisf = []int{c}
		} else if contador == maiorf {
			maisf = append(maisf, c)
		}
	}
	fmt.Print("[ ")
	for c := range maisf {
		fmt.Printf("%d ", maisf[c])
	}
	fmt.Println("]")
}
