package main

import "fmt"

func main() {
	var s, q int
	fmt.Scan(&s, &q)
	caracteres := make([]rune, 0)
	partida := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		caracteres = append(caracteres, x)
	}
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		partida = append(partida, x)
	}
    mapa_subs:= make([]int, len(caracteres))
    for i := range mapa_subs {
        mapa_subs[i] = i
    }
    partida_int := make([]int, len(partida))
    for i := range partida {
        for j := range caracteres {
            if partida[i] == caracteres[j] {
                partida_int[i] = j
            }
        }
    }
	for range q {
        partida_int[s-1]++
        for i := s-1; i >= 0; i-- {
            if partida_int[i] < len(caracteres) {
                break
            } else {
                partida_int[i] = 0
                if i > 0{
                    partida_int[i-1]++
                }
            }
        }
		for i := range partida_int {
            fmt.Printf("%c", caracteres[partida_int[i]])
        }
        fmt.Println()
	}
}
