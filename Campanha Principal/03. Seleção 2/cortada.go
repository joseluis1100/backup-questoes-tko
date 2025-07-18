package main

import "fmt"

func main() {
	var b, t int
	fmt.Scan(&b, &t)
	if (b+t)*70/2 > 5600 {
		fmt.Println(1)
	} else if (b+t)*70/2 < 5600 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}
