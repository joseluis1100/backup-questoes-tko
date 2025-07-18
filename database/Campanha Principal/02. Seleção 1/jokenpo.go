package main

import "fmt"

func main() {
	var j1, j2 string
	fmt.Scan(&j1, &j2)
	if j1 == j2 {
		fmt.Println("empate")
	} else if j1 == "R" && j2 == "S" || j1 == "P" && j2 == "R" || j1 == "S" && j2 == "P" {
		fmt.Println("jog1")
	} else {
		fmt.Println("jog2")
	}
}
