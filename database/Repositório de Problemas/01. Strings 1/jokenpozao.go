package main

import "fmt"

func string_to_num(jogador string) int {
	if jogador == "paper" {
		return 0
	} else if jogador == "air" {
		return 1
	} else if jogador == "water" {
		return 2
	} else if jogador == "gun" {
		return 3
	} else if jogador == "rock" {
		return 4
	} else if jogador == "fire" {
		return 5
	} else if jogador == "scissors" {
		return 6
	} else if jogador == "human" {
		return 7
	} else if jogador == "sponge" {
		return 8
	}
	return 0
}

func main() {
	var str_jog1, str_jog2 string
	fmt.Scan(&str_jog1, &str_jog2)
	jog1 := string_to_num(str_jog1)
	jog2 := string_to_num(str_jog2)
	if jog1 == jog2 {
		fmt.Println("empate")
	} else if (jog2-jog1+9)%9 <= 4 {
		fmt.Println("jog1")
	} else {
		fmt.Println("jog2")
	}
}
