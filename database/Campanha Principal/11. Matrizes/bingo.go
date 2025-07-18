package main

import "fmt"

func main() {
	matriz := [][]int{
		{1, 9, 27, 23},
		{34, 20, 37, 47},
		{30, 87, 55, 69},
		{13, 60, 99, 66}}
	vetor := make([]int, 6)
	var contador int
	for _, valor := range vetor {
		fmt.Scan(&valor)
		for i := range matriz {
			for j := range matriz[i] {
				if valor == matriz[i][j] {
					contador++
				}
			}
		}
	}
    fmt.Println(contador)
}
