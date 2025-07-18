package main

import "fmt"

func main() {
	var lado, maior int
	fmt.Scan(&lado)
	soma_linha := make([]int, lado)
	soma_coluna := make([]int, lado)
	matriz := make([][]int, lado)
	for i := range matriz {
		matriz[i] = make([]int, lado)
	}
	for i := range matriz {
		for j := range matriz[i] {
			fmt.Scan(&matriz[i][j])
			soma_linha[i] += matriz[i][j]
			soma_coluna[j] += matriz[i][j]
		}
	}
	for i := range soma_linha {
		for j := range soma_coluna {
			soma := soma_linha[i] + soma_coluna[j] - matriz[i][j]*2
            if soma > maior {
                maior = soma
            }
		}
	}
	fmt.Println(maior)
}
