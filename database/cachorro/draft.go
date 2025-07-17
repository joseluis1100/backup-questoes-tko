package main

import "fmt"

func main() {
    var x, x2 rune
	letras := make([]rune, 0)
	for range 200 {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		if x == ' ' && x2 == ' ' {
			continue
		}
		x2 = x
		letras = append(letras, x)
	}
	fmt.Println(string(letras))
}
