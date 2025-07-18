package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3, t float64
	fmt.Scan(&n1, &n2, &n3, &t)
	if (n1+n2+n3+t-math.Min(math.Min(n1, n2), n3))/3 >= 7 {
		fmt.Printf("Aprovado com %.1f\n", (n1+n2+n3+t-math.Min(math.Min(n1, n2), n3))/3)
	} else {
		fmt.Printf("Final com %.1f\n", (n1+n2+n3+t-math.Min(math.Min(n1, n2), n3))/3)
	}
}
