package main

import "fmt"

func main() {
	var n int
    fmt.Scan(&n)
    fatores := make(map[int]int)
    i := 2
    for n != 1 {
        if n % i == 0 {
            n /= i
            fatores[i]++
        } else {
            i++
        }
    }
    for valor, vezes := range fatores {
        fmt.Println(valor, vezes)
    }
}
