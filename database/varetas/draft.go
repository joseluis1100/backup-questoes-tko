package main

import "fmt"

func main() {
	var l1, l2, l3 int
	fmt.Scan(&l1, &l2, &l3)
	if l1 >= l2+l3 || l2 >= l1+l3 || l3 >= l1+l2 {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}