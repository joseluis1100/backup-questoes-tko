package main

import "fmt"

func main() {
	var p bool
	var d1, d2 int
	fmt.Scan(&p, &d1, &d2)
	if (d1+d2)%2 == 0 && !p || (d1+d2)%2 != 0 && p {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
