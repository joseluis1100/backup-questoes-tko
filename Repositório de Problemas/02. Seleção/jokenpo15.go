package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a == b {
		fmt.Println("Empate")
	} else if ((b-a)%15+15)%15 < 8 {
		fmt.Println("Jogador 1")
	} else {
		fmt.Println("Jogador 2")
	}
}
