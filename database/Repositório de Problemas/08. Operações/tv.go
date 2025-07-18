package main

import "fmt"

func main() {
	var v, p float32
	fmt.Scan(&v, &p)
	switch p {
	case 1:
		fmt.Printf("%.2f\n%.2f\n", v, v)
	case 2:
		fmt.Printf("%.2f\n%.2f\n", v*1.05/2, v*1.05)
	case 3:
		fmt.Printf("%.2f\n%.2f\n", v*1.1/3, v*1.1)
	case 4:
		fmt.Printf("%.2f\n%.2f\n", v*1.15/4, v*1.15)
	case 5:
		fmt.Printf("%.2f\n%.2f\n", v*1.2/5, v*1.2)
	case 6:
		fmt.Printf("%.2f\n%.2f\n", v*1.25/6, v*1.25)
	case 7:
		fmt.Printf("%.2f\n%.2f\n", v*1.3/7, v*1.3)
	case 8:
		fmt.Printf("%.2f\n%.2f\n", v*1.35/8, v*1.35)
	case 9:
		fmt.Printf("%.2f\n%.2f\n", v*1.4/9, v*1.4)
	case 10:
		fmt.Printf("%.2f\n%.2f\n", v*1.45/10, v*1.45)
	}
}
