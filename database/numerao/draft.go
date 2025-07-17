package main

import "fmt"

func main() {
	var numero string
	var soma_par, soma_impar int
	fmt.Scan(&numero)
	for i := range len(numero) {
		if i%2 == 0 {
			soma_impar += int(numero[i])-48
		} else {
			soma_par += int(numero[i])-48
		}
	}
	if soma_par-soma_impar == 0 || soma_par-soma_impar%11 == 0 {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}
