package main

import "fmt"

var fatores = make(map[int]int)

func calc_fatores(n int, i int, fatores map[int]int) {
    if n == 1 {
        return
    }
    if n % i == 0 {
        fatores[i]++
        calc_fatores(n/i, i, fatores)
    } else {
        calc_fatores(n, i+1, fatores)
    }
}

func main() {
	var n int
    fmt.Scan(&n)
    calc_fatores(n, 2, fatores)
    //for n != 1 {
        //if n % i == 0 {
            //n /= i
            //fatores[i]++
        //} else {
            //i++
        //}
    //}
    for valor, vezes := range fatores {
        fmt.Println(valor, vezes)
    }
}
