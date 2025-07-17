package main

import "fmt"

type Fator struct {
	num int
	qtd int
}

func calc_fatores(num int) []Fator {
	valores := make([]Fator, 0)
	var contagem int
	i := 2
	for num != 1 {
		if num%i == 0 {
			num /= i
			contagem++
		} else {
			if contagem != 0 {
				valores = append(valores, Fator{num: i, qtd: contagem})
			}
			i++
			contagem = 0
		}
	}
	valores = append(valores, Fator{num: i, qtd: contagem})
	return valores
}
func main() {
	var n int
	fmt.Scan(&n)
	vetor := calc_fatores(n)
	for i := range vetor {
		fmt.Printf("%v %v\n", vetor[i].num, vetor[i].qtd)
	}

}
