package main

import "fmt"

func main() {
	var lado, gladiador, condenado int
	fmt.Scan(&lado)
	matriz := make([][]string, lado)
	for i := range matriz {
		matriz[i] = make([]string, lado)
	}
	for i := range matriz {
		for j := range matriz[i] {
			fmt.Scan(&matriz[i][j])
		}
	}
	for i := range matriz {
		for j := range matriz[i] {
			if matriz[i][j] == "L" {
				for k := range matriz[i] {
					matriz[i][k] = "0"
				}
			}
		}
	}
	for i := range matriz {
		for j := range matriz[i] {
			if matriz[i][j] == "G" {
				gladiador += 2
			} else if matriz[i][j] == "C" {
				condenado++
			}
		}
	}
    for i := range matriz {
        if matriz[i][len(matriz)-i-1] == "C" {
            condenado++
        }
    }
	if gladiador == condenado {
		fmt.Println("Ninguem")
	} else if gladiador > condenado {
		fmt.Println("Gladiadores")
	} else {
		fmt.Println("Condenados a morte")
	}
}
