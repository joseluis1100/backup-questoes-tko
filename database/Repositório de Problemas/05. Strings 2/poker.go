package main

import (
	"fmt"
	"slices"
)

func main() {
	var quantidade int
	fmt.Scan(&quantidade)
	for q := range quantidade {
		pares := make([]int, 0)
		var pontuacao int
		mao := make([]int, 5)
		cartas := make(map[int]int)
		for i := range mao {
			fmt.Scan(&mao[i])
			cartas[mao[i]]++
		}
		slices.Sort(mao)
		if mao[1] == mao[0]+1 && mao[2] == mao[1]+1 && mao[3] == mao[2]+1 && mao[4] == mao[3]+1 {
			pontuacao += mao[0] + 200
		} else {
			if len(cartas) == 2 {
				for carta, valor := range cartas {
					if valor == 4 {
						pontuacao += carta + 180
					}
					if valor == 3 {
						pontuacao += carta + 160
					}
				}
			} else if len(cartas) == 3 {
				for carta, valor := range cartas {
					if valor == 3 {
						pontuacao += carta + 140
					}
					if valor == 2 {
						pares = append(pares, carta)
					}
				}
                if len(pares) == 2 {
						slices.Sort(pares)
						pontuacao += 3*pares[1] + 2*pares[0] + 20
					}
			} else if len(cartas) == 4 {
				for carta, valor := range cartas {
					if valor == 2 {
						pontuacao += carta
					}
				}
			}
		}
		fmt.Printf("Teste %v\n", q+1)
		fmt.Println(pontuacao)
		fmt.Println()
	}
}
