package main

import (
	"fmt"
	"math"
)

func main() {
	var v, c1, c2 float64
	fmt.Scan(&v, &c1, &c2)
	if math.Abs(v-c1) < math.Abs(v-c2) {
		fmt.Println("primeiro")
	} else if math.Abs(v-c1) > math.Abs(v-c2) {
		fmt.Println("segundo")
	} else {
		fmt.Println("empate")
	}
}
