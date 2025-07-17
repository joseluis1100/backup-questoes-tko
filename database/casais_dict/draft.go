package main

import "fmt"

func main() {
	var quantidade, casais int
	fmt.Scan(&quantidade)
	vetor := make([]int, quantidade)
	mapa := make(map[int]int)
	for i := range vetor {
		fmt.Scan(&vetor[i])
		mapa[vetor[i]] += 1
	}
	for animal := range mapa {
		if _, existe := mapa[animal*-1]; existe {
            casais += min(mapa[animal], mapa[-animal])
		}
	}
    fmt.Println(casais/2)
}
