package main

import (
	"fmt"
	"math/rand"
	"time"
)

func converter(aleatorio int) rune {
	if aleatorio == 1 {
		return 'A'
	}
	if aleatorio == 11 {
		return 'J'
	}
	if aleatorio == 12 {
		return 'Q'
	}
	if aleatorio == 13 {
		return 'K'
	}
	return rune(aleatorio) + '0'
}

func ajuste(aleatorio int, soma_player int) int {
	if aleatorio > 10 {
		return 10
	}
	if aleatorio == 1 && soma_player < 11 {
		return 11
	}
	return aleatorio
}

func main() {
	dinheiro := 100
	for i := 1; dinheiro >= 0; i++ {
		if dinheiro == 0 {
			fmt.Println("Você faliu!")
			return
		}
		var aposta int
		fmt.Printf("Rodada %v:\n", i)
		fmt.Printf("Dinheiro: %v\n", dinheiro)
		for {
			fmt.Print("Digite valor da aposta ou -1 para sair: ")
			fmt.Scan(&aposta)
			if aposta == -1 {
				return
			}
			if aposta > dinheiro || aposta <= 0 {
				fmt.Println("Valor inválido.")
			} else {
				break
			}
		}
		var soma_mesa, soma_mao, aleatorio int
		historico_mesa := make([]int, 0)
		historico_mao := make([]int, 0)
		fmt.Println("Iniciando Rodada:")
		time.Sleep(1 * time.Second)
		aleatorio = rand.Intn(13) + 1
		mesa := []rune{converter(aleatorio)}
		soma_mesa += ajuste(aleatorio, soma_mesa)
		historico_mesa = append(historico_mesa, ajuste(aleatorio, soma_mesa))
		fmt.Printf("# Mesa recebe %c - Total %v %c\n", converter(aleatorio), soma_mesa, mesa)
		time.Sleep(1 * time.Second)
		aleatorio = rand.Intn(13) + 1
		mao := []rune{converter(aleatorio)}
		soma_mao = ajuste(aleatorio, soma_mao)
		historico_mao = append(historico_mao, ajuste(aleatorio, soma_mao))
		fmt.Printf("# Voce recebe %c - Total %v %c\n", converter(aleatorio), soma_mao, mao)
		time.Sleep(1 * time.Second)
		estouorou := false
		for {
			var continua int
			aleatorio = rand.Intn(13) + 1
			mao = append(mao, converter(aleatorio))
			soma_mao += ajuste(aleatorio, soma_mao)
			historico_mao = append(historico_mao, ajuste(aleatorio, soma_mao))
			fmt.Printf("# Voce recebe %c - Total %v %c\n", converter(aleatorio), soma_mao, mao)
			if soma_mao > 21 {
				for i := range historico_mao {
					if historico_mao[i] == 11 {
						historico_mao[i] = 1
						soma_mao -= 10
					}
				}
			}
			if soma_mao > 21 {
				soma_mao = -1
				fmt.Println("ESTOUROU! Mesa ganha!")
				dinheiro -= aposta
				estouorou = true
				break
			}
			fmt.Println("Pedir = 1, Parar = 2")
			fmt.Scan(&continua)
			if continua == 2 {
				break
			}
		}
		if !estouorou {
			for {
				aleatorio = rand.Intn(13) + 1
				mesa = append(mesa, converter(aleatorio))
				soma_mesa += ajuste(aleatorio, soma_mesa)
				historico_mesa = append(historico_mesa, ajuste(aleatorio, soma_mesa))
				time.Sleep(1 * time.Second)
				fmt.Printf("# Mesa recebe %c - Total %v %c\n", converter(aleatorio), soma_mesa, mesa)
				if soma_mesa > 21 {
					fmt.Println("MESA ESTOURTOU! Você ganha!")
					dinheiro += aposta
					break
				}
				if soma_mesa >= soma_mao {
					fmt.Println("Mesa ganha!")
					dinheiro -= aposta
					break
				}
			}
		}
	}
}
