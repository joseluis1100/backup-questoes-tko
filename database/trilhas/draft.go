package main

import "fmt"

func main() {
	var quantidade, menor int
	fmt.Scan(&quantidade)
	esforco := make([]int, quantidade)
	for i := range quantidade {
		var qtd_pontos, indo, voltando int
		fmt.Scan(&qtd_pontos)
		pontos := make([]int, qtd_pontos)
		for j := range pontos {
			fmt.Scan(&pontos[j])
		}
		for j := range pontos {
			if j != len(pontos)-1 && pontos[j] < pontos[j+1] {
				indo += pontos[j+1] - pontos[j]
			}
		}
		for j := len(pontos) - 1; j > 0; j-- {
			if pontos[j] < pontos[j-1] {
				voltando += pontos[j-1] - pontos[j]
			}
		}
		if indo < voltando {
			esforco[i] = indo
		} else {
			esforco[i] = voltando
		}
		if i > 0 && esforco[i] < esforco[menor] {
			menor = i
		}
	}
	fmt.Println(menor+1)
}
