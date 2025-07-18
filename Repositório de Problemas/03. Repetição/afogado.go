package main

import "fmt"

func main() {
	var p, e int
	fmt.Scan(&p, &e)
	for f := e; ; f++ {
		l := 0
		for v := 1; ; v++ {
            if l < 0 {
                break
            }
			if v == 1 {
				l += (f - 10*(v-1))
			} else {
				l += (f - 10*(v-1)) - e
			}
			if l >= p {
				fmt.Println(f)
				return
			}
		}
	}
}
