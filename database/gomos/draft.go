package main

import "fmt"

type Pos struct {
	x int
	y int
}

func main() {
	var quantidade int
	var direcao rune
	fmt.Scanf("%d %c", &quantidade, &direcao)
	gomo := make([]Pos, quantidade)
	for i := range gomo {
		fmt.Scan(&gomo[i].x, &gomo[i].y)
	}
	if direcao == 'U' {
		fmt.Println(gomo[0].x, gomo[0].y-1)
	} else if direcao == 'D' {
		fmt.Println(gomo[0].x, gomo[0].y+1)
	} else if direcao == 'L' {
		fmt.Println(gomo[0].x-1, gomo[0].y)
	} else if direcao == 'R' {
		fmt.Println(gomo[0].x+1, gomo[0].y)
	}
	for i := range len(gomo) - 1 {
		fmt.Println(gomo[i].x, gomo[i].y)
	}
}
