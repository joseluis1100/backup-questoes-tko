package main

import (
	"fmt"
	"slices"
)

func main() {
	var quantidade, diferentes, maior_freq int
	mais_freq := make([]int, 0)
	fmt.Scan(&quantidade)
	mapa := make(map[int]int)
	vetor := make([]int, quantidade)
	for i := range vetor {
		fmt.Scan(&vetor[i])
		if _, existe := mapa[vetor[i]]; !existe {
			diferentes++
		}
		mapa[vetor[i]] += 1
	}
	fmt.Println(diferentes)
	for valor, freq := range mapa {
		if freq > maior_freq {
			maior_freq = freq
			mais_freq = []int{valor}
		} else if freq == maior_freq {
			mais_freq = append(mais_freq, valor)
		}
	}
	slices.Sort(mais_freq)
	for i := range mais_freq {
		fmt.Print(mais_freq[i])
		if i == len(mais_freq)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
