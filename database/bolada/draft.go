package main

import "fmt"

func main() {
	var t string
	var f, p int
	fmt.Scan(&t, &f)
	switch t {
	case "b":
		p = ((f * 20) - 80) / 10
	case "c":
		p = ((f * 18) - 80) / 10
	}
	if p < 150 {
		fmt.Println("Fraco, nem passou")
	} else if p < 180 {
		fmt.Println("Perfeito")
	} else if p <= 210 {
		fmt.Println("Satisfeito")
	} else if p > 210 {
		fmt.Println("Muito forte, bola fora")
	}
}
