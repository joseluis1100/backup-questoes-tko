package main

import (
	"fmt"
	"strings"
)

type pares struct {
	x rune
	y rune
}

func prepararFrase() []pares {
	frase := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		if x == ' ' {
			continue
		}
		frase = append(frase, x)
	}
	frase_formatada := []rune(strings.ToUpper(string(frase)))

	for i := range frase_formatada {
		if i < len(frase_formatada)-1 && frase_formatada[i] == frase_formatada[i+1] && i%2 == 0 {
			frase_formatada = append(frase_formatada[:i+1], append([]rune{'X'}, frase_formatada[i+1:]...)...)
		}
	}
	if len(frase_formatada)%2 != 0 {
		frase_formatada = append(frase_formatada, 'X')
	}
	frase_dividida := make([]pares, 0)
	for i := range frase_formatada {
		if i%2 == 0 {
			frase_dividida = append(frase_dividida, pares{x: frase_formatada[i], y: frase_formatada[i+1]})
		}
	}
	return frase_dividida
}

func criar_matriz() [][]rune {
	usadas := make(map[rune]int)
	nao_usadas := make(map[rune]int)
	for i := 'A'; i <= 'Z'; i++ {
		if i != 'W' {
			nao_usadas[i]++
		}
	}
	chave := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		if x == ' ' || x == 'w' || x == 'W' {
			continue
		}
		x = rune(strings.ToUpper(string(x))[0])
		usadas[x]++
		delete(nao_usadas, x)
		chave = append(chave, x)
	}
	matriz := make([][]rune, 5)
	for i := range matriz {
		matriz[i] = make([]rune, 5)
	}
	posx, posy := 0, 0
	for _, letra := range chave {
		if _, existe := usadas[letra]; existe {
			matriz[posx][posy] = letra
			posy++
			if posy == 5 {
				posy = 0
				posx++
			}
			delete(usadas, letra)
		}
	}
	for letra := 'A'; letra <= 'Z'; letra++ {
		if _, existe := nao_usadas[letra]; existe {
			matriz[posx][posy] = letra
			posy++
			if posy == 5 {
				posy = 0
				posx++
			}
		}
	}
	return matriz
}

func main() {
	frase := prepararFrase()
	for _, par := range frase {
		fmt.Printf("%c%c ", par.x, par.y)
	}
	fmt.Println()
	matriz := criar_matriz()
	for i := range matriz {
		for j := range matriz[i] {
			fmt.Print(string(matriz[i][j]), " ")
		}
		fmt.Println()
	}
	posx_a, posy_a := 0, 0
	posx_b, posy_b := 0, 0
	for pos_par, par := range frase {
		for i := range matriz {
			for j := range matriz[i] {
				if matriz[i][j] == par.x {
					posx_a, posy_a = i, j
				}
				if matriz[i][j] == par.y {
					posx_b, posy_b = i, j
				}
			}
		}
		if posx_a == posx_b {
			posy_a = (posy_a + 1) % 5
			posy_b = (posy_b + 1) % 5
		} else if posy_a == posy_b {
			posx_a = (posx_a + 1) % 5
			posx_b = (posx_b + 1) % 5
		} else {
			posy_a, posy_b = posy_b, posy_a
		}
		frase[pos_par] = pares{x: matriz[posx_a][posy_a], y: matriz[posx_b][posy_b]}
	}
	for _, par := range frase {
		fmt.Printf("%c%c ", par.x, par.y)
	}
	fmt.Println()
}
