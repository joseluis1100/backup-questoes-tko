package main

import "fmt"

func main() {
	var n string
	var i int
	fmt.Scan(&n, &i)
	if i < 12 {
		fmt.Printf("%s eh crianca\n", n)
	} else if i < 18 {
		fmt.Printf("%s eh jovem\n", n)
	} else if i < 65 {
		fmt.Printf("%s eh adulto\n", n)
	} else if i < 1000 {
		fmt.Printf("%s eh idoso\n", n)
	} else {
		fmt.Printf("%s eh mumia\n", n)
	}
}
