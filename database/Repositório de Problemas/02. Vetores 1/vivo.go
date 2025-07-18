package main

import "fmt"

func main() {
    teste := 1
	for {
		var p, r, n, s int
		fmt.Scan(&p, &r)
		jogadores := make([]int, p)
		status := make([]int, p)
		passa := make([]int, 0)
        if p == 0 && r == 0 {
                return
            }
		for i := range jogadores {
			fmt.Scan(&jogadores[i])
		}
		for range r {
			fmt.Scan(&n, &s)
			for j := range jogadores {
				fmt.Scan(&status[j])
				if status[j] == s {
					passa = append(passa, jogadores[j])
				}
			}
            jogadores = []int{}
            jogadores = append(jogadores, passa...)
            passa = []int{}
		}
        fmt.Printf("Teste %v\n%v\n", teste, jogadores[0])
        teste++
	}
}
