package main

import "fmt"

func main() {
	var a, b int
	soma := 0
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("invalido")
	} else {
		for ; a != b+1; a++ {
			if a%2 == 0 && a%3 == 0 {
				soma += a
			}
		}
		fmt.Println(soma)
	}
}
