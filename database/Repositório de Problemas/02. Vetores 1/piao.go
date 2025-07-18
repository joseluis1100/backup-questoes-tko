package main

import (
	"fmt"
	"math"
)

func main() {
	var l, n int
	vencedor, perdedor := -1, 0
	fmt.Scan(&l, &n)
	jogadores := make([]int, n)
	for i := range jogadores {
		fmt.Scan(&jogadores[i])
		if int(math.Abs(float64(jogadores[i]))) <= l {
			if vencedor == -1 {
				vencedor = 0
			}
			if int(math.Abs(float64(jogadores[i]))) <= int(math.Abs(float64(jogadores[vencedor]))) {
				vencedor = i
			}
		}
		if int(math.Abs(float64(jogadores[i]))) >= int(math.Abs(float64(jogadores[perdedor]))) {
			perdedor = i
		}
	}
	if vencedor == -1 {
		fmt.Println("nenhum")
	} else {
		fmt.Println(vencedor)
	}
	fmt.Println(perdedor)
}
