package main

import (
	"fmt"
	"math"
)
func main() {
    var lado, maior int
    fmt.Scan(&lado)
    soma_coluna := make([]int, lado)
    matriz := make([][]int, lado)
    for i := range matriz {
        matriz[i] = make([]int, lado)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
            soma_coluna[j] += int(math.Pow(float64(matriz[i][j]), 2))
        }
    }
    for i := range soma_coluna {
        if soma_coluna[i] > soma_coluna[maior] {
            maior = i
        }
    }
    fmt.Println(maior)
}
