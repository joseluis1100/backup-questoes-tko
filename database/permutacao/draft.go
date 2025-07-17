package main

import "fmt"

func main() {
	partida := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		partida = append(partida, x)
	}
	var avancos int
	fmt.Scan(&avancos)
	for range avancos {
		i := len(partida) - 2
		for i >= 0 && partida[i] >= partida[i+1] {
			i--
		}
		if i < 0 {
			fmt.Println(string(partida))
			return
		}
		j := len(partida) - 1
		for partida[j] <= partida[i] {
			j--
		}
		partida[i], partida[j] = partida[j], partida[i]
		for k, l := i+1, len(partida)-1; k < l; k, l = k+1, l-1 {
			partida[k], partida[l] = partida[l], partida[k]
		}
	}
	for i := range partida {
		fmt.Printf("%c", partida[i])
	}
	fmt.Println()
}
