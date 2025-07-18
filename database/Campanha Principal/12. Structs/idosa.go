package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
	sexo  rune
}

func main() {
	var quantidade int
	fmt.Scan(&quantidade)
	var pessoa, idosa Pessoa
	for range quantidade + 2 {
		fmt.Scanf("%s %d %c", &pessoa.nome, &pessoa.idade, &pessoa.sexo)
		if pessoa.idade > idosa.idade && pessoa.sexo == 'f' {
			idosa = pessoa
		}
	}
	if idosa.idade == 0 {
		fmt.Println("nao tem mulher")
	} else {
		fmt.Println(idosa.nome)
	}
}
