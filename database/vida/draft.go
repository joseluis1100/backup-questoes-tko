package main

import "fmt"

func ver_esquerda(matriz [][]rune,i int, j int) int {
	if matriz[i][j-1] == '#' {
		return 1
	}
	return 0
}
func ver_direita(matriz [][]rune, i int, j int) int {
	if matriz[i][j+1] == '#' {
		return 1
	}
	return 0
}
func ver_cima(matriz [][]rune, i int, j int) int {
	if matriz[i-1][j] == '#' {
		return 1
	}
	return 0
}
func ver_baixo(matriz [][]rune, i int, j int) int {
	if matriz[i+1][j] == '#' {
		return 1
	}
	return 0
}
func ver_diagonal_noroeste(matriz [][]rune, i int, j int) int {
	if matriz[i-1][j-1] == '#' {
		return 1
	}
	return 0
}
func ver_diagonal_nordeste(matriz [][]rune, i int, j int) int {
	if matriz[i-1][j+1] == '#' {
		return 1
	}
	return 0
}
func ver_diagonal_sudoeste(matriz [][]rune, i int, j int) int {
	if matriz[i+1][j-1] == '#' {
		return 1
	}
	return 0
}
func ver_diagonal_sudeste(matriz [][]rune, i int, j int) int {
	if matriz[i+1][j+1] == '#' {
		return 1
	}
	return 0
}

func main() {
	var l, c int
	fmt.Scan(&l, &c)
	matriz := make([][]rune, l)
	for i := range matriz {
		matriz[i] = make([]rune, c)
	}
	nova_matriz := make([][]rune, l)
	for i := range nova_matriz {
		nova_matriz[i] = make([]rune, c)
	}
	var x rune
	var linha int
	for linha < l {
		for i := range c {
			fmt.Scanf("%c", &x)
		if x == '\n' {
			linha++
			break
		}
        matriz[linha][i] = x
		}
	}
	for i := range matriz {
		for j := range matriz[i] {
			nova_matriz[i][j] = matriz[i][j]
		}
	}
	for i := range matriz {
		for j := range matriz[i] {
			var vivos int
			if i == 0 && j == 0 {
				vivos += ver_direita(matriz, i, j)
				vivos += ver_diagonal_sudeste(matriz, i, j)
				vivos += ver_baixo(matriz, i, j)
			} else if i == 0 && j == len(matriz[i])-1 {
				vivos += ver_esquerda(matriz, i, j)
				vivos += ver_diagonal_sudoeste(matriz, i, j)
				vivos += ver_baixo(matriz, i, j)
			} else if i == len(matriz)-1 && j == 0 {
				vivos += ver_cima(matriz, i, j)
				vivos += ver_diagonal_nordeste(matriz, i, j)
				vivos += ver_direita(matriz, i, j)
			} else if i == len(matriz)-1 && j == len(matriz[i])-1 {
				vivos += ver_cima(matriz, i, j)
				vivos += ver_diagonal_noroeste(matriz, i, j)
				vivos += ver_esquerda(matriz, i, j)
			} else if i == 0 {
				vivos += ver_direita(matriz, i, j)
				vivos += ver_diagonal_sudoeste(matriz, i, j)
				vivos += ver_esquerda(matriz, i, j)
				vivos += ver_diagonal_sudeste(matriz, i, j)
				vivos += ver_baixo(matriz, i, j)
			} else if i == len(matriz)-1 {
				vivos += ver_esquerda(matriz, i, j)
				vivos += ver_diagonal_noroeste(matriz, i, j)
				vivos += ver_cima(matriz, i, j)
				vivos += ver_diagonal_nordeste(matriz, i, j)
				vivos += ver_direita(matriz, i, j)
			} else if j == 0 {
				vivos += ver_cima(matriz, i, j)
				vivos += ver_diagonal_nordeste(matriz, i, j)
				vivos += ver_direita(matriz, i, j)
				vivos += ver_diagonal_sudeste(matriz, i, j)
				vivos += ver_baixo(matriz, i, j)
			} else if j == len(matriz[i])-1 {
				vivos += ver_cima(matriz, i, j)
				vivos += ver_diagonal_noroeste(matriz, i, j)
				vivos += ver_esquerda(matriz, i, j)
				vivos += ver_diagonal_sudoeste(matriz, i, j)
				vivos += ver_baixo(matriz, i, j)
			} else {
				vivos += ver_cima(matriz, i, j)
				vivos += ver_diagonal_noroeste(matriz, i, j)
				vivos += ver_esquerda(matriz, i, j)
				vivos += ver_diagonal_sudoeste(matriz, i, j)
				vivos += ver_baixo(matriz, i, j)
				vivos += ver_diagonal_sudeste(matriz, i, j)
				vivos += ver_direita(matriz, i, j)
				vivos += ver_diagonal_nordeste(matriz, i, j)
			}
			if matriz[i][j] == '#' && (vivos < 2 || vivos > 3) {
				nova_matriz[i][j] = '.'
			}
			if matriz[i][j] == '.' && vivos == 3 {
				nova_matriz[i][j] = '#'
			}
		}
	}
	for i := range nova_matriz {
		for j := range nova_matriz[i] {
			fmt.Printf("%c", nova_matriz[i][j])
		}
		fmt.Println()
	}
}
