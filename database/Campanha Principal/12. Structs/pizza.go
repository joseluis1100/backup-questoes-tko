package main

import "fmt"

type restaurante struct {
	nome string
	nota int
}

func main() {
	var q int
	fmt.Scan(&q)
	m := restaurante{"", 0}
	for range q {
		var nome string
		var nota int
		fmt.Scan(&nome, &nota)
		r := restaurante{nome, nota}
		if r.nota >= m.nota {
			m.nome, m.nota = r.nome, r.nota
		}
	}
	fmt.Println(m.nome)
}
