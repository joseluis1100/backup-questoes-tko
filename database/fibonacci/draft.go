package main

import "fmt"

func main() {
	var n, soma int
	fmt.Scan(&n)
	f := make([]int, 0)
	f = append(f, 1)
	f = append(f, 1)
	for c := 2; f[c-1]+f[c-2] < n; c++ {
		f = append(f, f[c-1]+f[c-2])
		if (f[c-1]+f[c-2])%2 == 0 {
			soma += f[c-1] + f[c-2]
		}
	}
	fmt.Println(soma)
}
