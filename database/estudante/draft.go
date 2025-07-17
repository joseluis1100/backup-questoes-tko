package main

import (
	"fmt"
)

type aluno struct {
	nome  []rune
	n1    float32
	n2    float32
	n3    float32
	media float32
}

func main() {
	var quantidade int
	var x rune
	fmt.Scan(&quantidade)
	alunos := make([]aluno, quantidade)
	for i := range alunos {
		for {
			fmt.Scanf("%c", &x)
			if x == '\n' {
				break
			}
			alunos[i].nome = append(alunos[i].nome, x)
		}
		fmt.Scan(&alunos[i].n1, &alunos[i].n2, &alunos[i].n3)
		alunos[i].media = (alunos[i].n1 + alunos[i].n2 + alunos[i].n3) / 3
	}
	for i := range alunos {
		for j := range len(alunos)-1-i {
			if alunos[j].media < alunos[j+1].media {
				alunos[j], alunos[j+1] = alunos[j+1], alunos[j]
			}
		}
	}
	for i := range alunos {
		fmt.Printf("%v: %v\n   Media: %.2f\n   N1: %.2f, N2: %.2f, N3: %.2f\n", i, string(alunos[i].nome), alunos[i].media, alunos[i].n1, alunos[i].n2, alunos[i].n3)
	}
}
