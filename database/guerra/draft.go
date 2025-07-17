package main

import "fmt"

type personagem struct {
	nome  []rune
	poder int
}

func main() {
	var qtd_ironman, qtd_cap, soma_ironman, soma_cap int
    var maior string
    var poder_maior int
	fmt.Scan(&qtd_ironman)
	pers_ironman := make([]personagem, qtd_ironman)
	var x rune
	for i := range pers_ironman {
		for {
			fmt.Scanf("%c", &x)
			if x == '\n' {
				break
			}
			pers_ironman[i].nome = append(pers_ironman[i].nome, x)
		}
        fmt.Scan(&pers_ironman[i].poder)
        if pers_ironman[i].poder > poder_maior {
            poder_maior = pers_ironman[i].poder
            maior = string(pers_ironman[i].nome)
        }
        soma_ironman += pers_ironman[i].poder
	}
    fmt.Scan(&qtd_cap)
    pers_cap := make([]personagem, qtd_cap)
    for i := range pers_cap {
		for {
			fmt.Scanf("%c", &x)
			if x == '\n' {
				break
			}
			pers_cap[i].nome = append(pers_cap[i].nome, x)
		}
        fmt.Scan(&pers_cap[i].poder)
        if pers_cap[i].poder > poder_maior {
            poder_maior = pers_cap[i].poder
            maior = string(pers_cap[i].nome)
        }
        soma_cap += pers_cap[i].poder
	}
    if soma_ironman == soma_cap {
        fmt.Println("Draw")
    } else if soma_cap > soma_ironman {
        fmt.Println("Team Captain Wins")
    } else {
        fmt.Println("Team Iron Wins")
    }
    fmt.Println(maior)
}
