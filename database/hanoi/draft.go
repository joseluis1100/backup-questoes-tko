package main

import "fmt"

func hanoi(n int, partida rune, chegada rune, suporte rune){
	if n > 0 {
		hanoi(n-1, partida, suporte, chegada)
		fmt.Printf("%c -> %c\n",partida, chegada)
		hanoi(n-1, suporte, chegada, partida)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	hanoi(n, 'A', 'C', 'B')
}
