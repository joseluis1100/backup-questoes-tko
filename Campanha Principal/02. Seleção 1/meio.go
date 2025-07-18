package main

import "fmt"

func comp(v1, v2, v3 int) int {
	if v1 <= v2 && v1 >= v3 || v1 >= v2 && v1 <= v3 {
		return v1
	} else if v2 <= v1 && v2 >= v3 || v2 >= v1 && v2 <= v3 {
		return v2
	} else {
		return v3
	}
}

func main() {
	var v1, v2, v3 int
	fmt.Scan(&v1, &v2, &v3)
	fmt.Println(comp(v1, v2, v3))
}
