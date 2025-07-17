package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var inf, sup int = 1, 100
	n := rand.Intn(sup) + 1
	var c int
	for {
		for {
			fmt.Printf("Diga um numero entre ]%v, %v[: ", inf, sup)
			fmt.Scan(&c)
			if c >= sup || c <= inf {
				println("Valor inválido, tente novamente!")
			} else {
                break
            }
		}
		if c == n {
			fmt.Println("ACERTOU!")
			break
		}
		if c < n {
			inf = c
		} else if c > n {
			sup = c
		}
		if inf+2 == sup {
			fmt.Printf("PERDEU! era %v\n", n)
			break
		}
	}
}
