package main

import "fmt"

func numero(x rune) int {
	if x == '1' {
		return 2
	} else if x == '2' {
		return 5
	} else if x == '3' {
		return 5
	} else if x == '4' {
		return 4
	} else if x == '5' {
		return 5
	} else if x == '6' {
		return 6
	} else if x == '7' {
		return 3
	} else if x == '8' {
		return 7
	} else if x == '9' {
		return 6
	}
	return 6
}

func main() {
	var n int
	fmt.Scan(&n)
	leds := make([]int, n)
	v := make([]rune, 0)
	for c := range n {
		for  {
			var x rune
			fmt.Scanf("%c", &x)
			if x == '\n' {
				break
			}
			v = append(v, x)
			leds[c] += numero(x)
		}
	}
	for c := range n {
		fmt.Println(leds[c], "leds")
	}
}
