package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	numeros := make([]int, n)
	sequencia := make([]int, m)
	contagem := make([]int, 10)
	for i := range numeros {
		fmt.Scan(&numeros[i])
	}
	for i := range sequencia {
		fmt.Scan(&sequencia[i])
		sequencia[i]--
	}
    contagem[numeros[0]]++
	for i := range len(sequencia)-1 {
		if sequencia[i] <= sequencia[i+1] {
			for j := sequencia[i]+1; j <= sequencia[i+1]; j++ {
				contagem[numeros[j]]++
			}
		} else {
			for j := sequencia[i]-1; j >= sequencia[i+1]; j-- {
				contagem[numeros[j]]++
			}
		}
	}
    fmt.Print("[ ")
	for i := range contagem {
        fmt.Print(contagem[i], " ")
    }
    fmt.Println("]")
}
