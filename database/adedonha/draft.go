package main

import "fmt"

func main() {
	var d rune
	fmt.Scan(&d)
	if d == 0 {
		fmt.Println("joguem de novo")
	} else {
		fmt.Printf("%c\n", ((d-1)%26+26)%26+97)
	}
}
