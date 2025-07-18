package main

import "fmt"

func main() {
	var q int
	var x rune
	fmt.Scan(&q)
	for range q {
		chave := make([]rune, 0)
		var iguais, tamanho int
		for {
			fmt.Scanf("%c", &x)
			if x == '\n' {
				break
			}
			chave = append(chave, x)
		}
		for {
			fmt.Scanf("%c", &x)
			if x == '\n' {
				break
			}
			for c2 := range chave {
				if x == chave[c2] {
					iguais++
				}
			}
			tamanho++
		}
		if iguais == tamanho {
			fmt.Println("chefe")
		} else if iguais <= tamanho/2 {
			fmt.Println("pessoa")
		} else {
			fmt.Println("ultron")
		}
	}
}
