package main

import "fmt"

func main() {
	var quantidade int
	fmt.Scan(&quantidade)
	palavras := make([][]rune, quantidade)
	var linha int
	var x rune
	for linha < quantidade {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			linha++
		} else {
			palavras[linha] = append(palavras[linha], x)
		}
	}
	for i := range palavras {
		vogais := make([]rune, 0)
		var seguidos, maior_seguidos int
        palavras[i] = append(palavras[i], ' ')
		for j := range palavras[i] {
			if palavras[i][j] == 'a' || palavras[i][j] == 'e' || palavras[i][j] == 'i' || palavras[i][j] == 'o' || palavras[i][j] == 'u' {
				seguidos++
			} else {
				if seguidos > maior_seguidos {
					maior_seguidos = seguidos
					vogais = []rune{}
					for k := range maior_seguidos {
						vogais = append(vogais, palavras[i][j-maior_seguidos+k])
					}
				}
				seguidos = 0
			}
		}
		fmt.Println(string(vogais))
	}
}
