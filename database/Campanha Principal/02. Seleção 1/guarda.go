package main

import "fmt"

func main() {
	var w, l, a bool
	fmt.Scan(&w, &l, &a)
	if !w {
		fmt.Println("you must connect to wifi")
	} else if !l {
		fmt.Println("you need to login first")
	} else if !a {
		fmt.Println("you must to login as admin")
	} else {
		fmt.Println("done")
	}
}
