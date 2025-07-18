package main

import (
	"fmt"
	"math"
)

type Jogador struct {
	id int
	A  int
	B  int
}

func main() {
	var n, a, b int
	fmt.Scan(&n)
	var vencedor Jogador
	vencedor.id = -1
	for i := range n {
		fmt.Scan(&a, &b)
		if a < 10 || b < 10 {
			continue
		}
		if vencedor.id == -1 || math.Abs(float64(a)-float64(b)) < math.Abs(float64(vencedor.A)-float64(vencedor.B)) {
			vencedor = Jogador{i, a, b}
		}
	}
    if vencedor.id == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(vencedor.id)
    }
}
