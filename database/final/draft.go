package main

import "fmt"

func main() {
	var n1, n2, f, m float32
	fmt.Scan(&n1, &n2, &f)
	m = (n1 + n2) / 2
	if m >= 7 {
		fmt.Println("aprovado")
	} else if m < 7 && m >= 4 {
		if (m+f)/2 >= 5 {
			fmt.Println("aprovado na final")
		} else {
			fmt.Println("reprovado na final")
		}
	} else {
		fmt.Println("reprovado")
	}
}
