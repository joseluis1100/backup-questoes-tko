package main

import "fmt"

func main() {
	var j1, j2, j3 int
	fmt.Scan(&j1, &j2, &j3)
	if j1 == j2 && j1 == j3 {
		fmt.Println("empate")
	} else if j2 == j3 && j1 != j2 {
		fmt.Println("jog1")
	} else if j1 == j3 && j2 != j1 {
		fmt.Println("jog2")
	} else {
		fmt.Println("jog3")
	}
}
