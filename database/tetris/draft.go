package main

import "fmt"

func main() {
	var l, c int
	fmt.Scan(&l, &c)
	matriz := make([][]rune, l)
	for i := range matriz {
		matriz[i] = make([]rune, c)
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
	nova_matriz := make([][]rune, l)
	for i := range nova_matriz {
		nova_matriz[i] = make([]rune, c)
	}
	for i := range nova_matriz {
		for j := range nova_matriz[i] {
			nova_matriz[i][j] = matriz[i][j]
			if matriz[i][j] == 'o' {
				nova_matriz[i][j] = '.'
			}
		}
	}
	desce := true
	for j := range c {
		for i := range matriz {
			if i < len(matriz)-1 && matriz[i][j] == 'o' && matriz[i+1][j] == '#' || i == len(matriz)-1 && matriz[i][j] == 'o' {
				desce = false
				break
			}
		}
	}
	if desce {
		for i := range matriz {
			for j := range matriz[i] {
				if i < len(matriz)-1 && matriz[i][j] == '.' && matriz[i+1][j] == 'o' {
					nova_matriz[i+1][j] = '.'
				}
                if i > 0 && matriz[i-1][j] == 'o' {
                    nova_matriz[i][j] = 'o'
                }
			}
		}
		for i := range matriz {
			for j := range matriz[i] {
				fmt.Printf("%c", nova_matriz[i][j])
			}
			fmt.Println()
		}
	} else {
		for i := range matriz {
			for j := range matriz[i] {
				fmt.Printf("%c", matriz[i][j])
			}
			fmt.Println()
		}
	}
}
