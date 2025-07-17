package main

import (
	"fmt"
	"slices"
)

func main() {
	var quantidade, numero, soma int
	fmt.Scan(&quantidade)
	vetor := make([]int, 0)
	for range quantidade {
		fmt.Scan(&numero)
		if numero == 0 {
            if len(vetor) > 0 {
                vetor = slices.Delete(vetor, len(vetor)-1, len(vetor))
            }
		} else {
			vetor = append(vetor, numero)
		}
	}
    for i := range vetor{
        soma += vetor[i]
    }
	fmt.Println(soma)
}
