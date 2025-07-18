package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)
	letras := []rune(c)
	if letras[0] >= 65 && letras[0] <= 90 {
		letras[0] += 32
	} else if letras[0] >= 97 && letras[0] <= 122 {
		letras[0] -= 32
	}
	fmt.Printf("%c\n", letras[0])
}
