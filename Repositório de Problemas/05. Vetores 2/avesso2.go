package main

import (
	"fmt"
	"math"
)

func main() {
	var n, qtd_x, x int
	fmt.Scan(&n, &qtd_x)
	jogadores := make([]int, n)
	for i := range jogadores {
		fmt.Scan(&jogadores[i])
	}
	for range qtd_x {
		fmt.Scan(&x)
		for j := range jogadores {
			if int(math.Abs(float64(jogadores[j]))) == x {
				if j == 0 {
					jogadores[1] *= -1
				} else if j == len(jogadores)-1 {
                    jogadores[len(jogadores)-2] *= -1
				} else {
                    jogadores[j-1] *= -1
                    jogadores[j+1] *= -1
                }
			}
		}
	}
    fmt.Println(jogadores)
}
