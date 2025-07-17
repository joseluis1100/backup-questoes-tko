package main

import "fmt"

type aluno struct {
	id    []rune
	media int
}

func main() {
	var quantidade, maior int
	fmt.Scan(&quantidade)
	alunos := make([]aluno, quantidade)
	var x rune
	for i := range alunos {
		for range 2 {
			fmt.Scanf("%c", &x)
            if x == ' ' {
                fmt.Scanf("%c", &x)
            }
			alunos[i].id = append(alunos[i].id, x)
		}
		for range 4 {
			fmt.Scanf("%c", &x)
			alunos[i].media += int(x - '0')
		}
        if alunos[i].media == alunos[maior].media {
            if int(alunos[maior].id[0]-'0')*10+int(alunos[maior].id[1]-'0') < int(alunos[i].id[0]-'0')*10+int(alunos[i].id[1]-'0') {
                maior = i
            }
        } else if alunos[i].media > alunos[maior].media {
            maior = i
        }
	}
    fmt.Println(string(alunos[maior].id))
}
