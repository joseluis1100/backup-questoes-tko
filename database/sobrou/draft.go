package main

import "fmt"

func main() {
	var q1, q2, q3, v1, v2, v3, d float32
	fmt.Scan(&q1, &q2, &q3, &v1, &v2, &v3, &d)
	fmt.Printf("%.2f\n", d-((q1*v1)+(q2*v2)+(q3*v3)))
}
