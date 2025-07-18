package main

import "fmt"

func main() {
	var p [4]string
	nota := 0
	for c := 0; c != 4; c++ {
		fmt.Scan(&p[c])
		if c == 0 && p[0] == "d" || c == 1 && p[1] == "a" || c == 2 && p[2] == "c" || c == 3 && p[3] == "d" {
			nota++
		}
	}
	switch nota {
	case 0:
		fmt.Println("Nunca assistiu")
	case 1:
		fmt.Println("Ja ouviu falar")
	case 2:
		fmt.Println("Interessado no assunto")
	case 3:
		fmt.Println("Fa")
	case 4:
		fmt.Println("Super Fa")
	}
}
