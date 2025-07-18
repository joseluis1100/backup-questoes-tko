package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func imprimirMatriz(matriz [][]string) {
	for i := range matriz {
		for j := range matriz[i] {
			if matriz[i][j] == "__" {
				fmt.Print("__ ")
			} else {
				valor, _ := strconv.Atoi(matriz[i][j])
				fmt.Printf("%02d ", valor)
			}
		}
		fmt.Println()
	}
}
func simulador_bingo() {
    num := 1
	roleta := make([][]string, 5)
	rack := make([][]string, 5)
	mapa := make(map[int]int)
	for i := range roleta {
		roleta[i] = make([]string, 15)
	}
	for i := range rack {
		rack[i] = make([]string, 15)
	}
	for i := range roleta {
		for j := range roleta[i] {
			roleta[i][j] = strconv.Itoa(num)
			mapa[num]++
			num++
		}
	}
	for i := range rack {
		for j := range rack[i] {
			rack[i][j] = "__"
		}
	}
	imprimirMatriz(roleta)
	for {
		for {
			fmt.Println("Escolha 1 para pedir bola e 0 para sair")
			var escolha int
			fmt.Scan(&escolha)
			if escolha == 1 {
				break
			}
			if escolha == 0 {
				return
			}
		}
		aleatorio := rand.Intn(75) + 1
		for {
			if _, existe := mapa[aleatorio]; existe {
				delete(mapa, aleatorio)
				break
			} else {
				aleatorio = rand.Intn(75) + 1
			}
		}
		fmt.Printf("Numero sorteado %v\n", aleatorio)
		fmt.Println("Roleta: ")
		for i := range roleta {
			for j := range roleta[i] {
				if roleta[i][j] == strconv.Itoa(aleatorio) {
					roleta[i][j] = "__"
					rack[i][j] = strconv.Itoa(aleatorio)
				}
			}
		}
		imprimirMatriz(roleta)
		fmt.Println("Rack: ")
		imprimirMatriz(rack)
	}
}

func gerarCartela() {
    cartela := make([][]string, 5)
    for i := range cartela {
        cartela[i] = make([]string, 5)
    }
    for i := range cartela {
        for j := range cartela[i] {
            valor := rand.Intn(15) + 1 + (j * 15)
            cartela[i][j] = strconv.Itoa(valor)
        }
    }
    cartela[2][2] = "##"
    for i := range cartela {
        for j := range cartela[i] {
            if cartela[i][j] == "##" {
                fmt.Print("## ")
            } else {
                valor, _ := strconv.Atoi(cartela[i][j])
                fmt.Printf("%02d ", valor)
            }
        }
        fmt.Println()
    }
}

func main() {
	for {
        fmt.Println("Escolha uma opção:")
        fmt.Println("[1] Simulador de Bingo")
        fmt.Println("[2] Gerador de cartelas")
        fmt.Println("[0] Sair")
        var opcao int
        fmt.Scan(&opcao)
        if opcao == 1 {
            simulador_bingo()
        } else if opcao == 2 {
            gerarCartela()
        } else if opcao == 0 {
            return
        } else {
            fmt.Println("Opção inválida, tente novamente.")
        }
    }
}
