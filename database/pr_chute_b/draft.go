package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var min, max, c int = 0, 100, rand.Intn(100) + 1
	for {
		var comp string
		fmt.Printf("]%v, %v[ É %v?\n", min, max, c)
		fmt.Print("Acertei(=), É maior(>), É menor(<)? ")
		fmt.Scanf("%s", &comp)
		if comp == ">" {
			min = c
			c = rand.Intn(max-min-1) + min + 1
		} else if comp == "<" {
			max = c
			c = rand.Intn(max-min-1) + min + 1
		} else if comp == "=" {
			fmt.Println("ganhei")
			break
		} else {
			fmt.Println("Valor inválido, tente novamente")
		}
		if min + 2 == max {
			fmt.Println("perdi")
			break
		}
	}
}
