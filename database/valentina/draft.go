package main

import "fmt"

func main() {
	var c1, o, c2 rune
	fmt.Scanf("%c\n%c\n%c", &c1, &o, &c2)
	if o == '+' {
		fmt.Printf("%c\n", (c1 + c2 - 194) % 26 + 97)
	} else {
        fmt.Printf("%c\n", (c1 - c2 + 26) % 26 + 97)
    }
}
