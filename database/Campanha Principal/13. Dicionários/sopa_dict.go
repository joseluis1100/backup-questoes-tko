package main

import "fmt"

var cache = make(map[int]int64)

func fibonacci(x int) int64 {
    if valor, existe := cache[x]; existe {
        return valor
    }
    if x == 0 {
        return 0
    }
    if x == 1 {
        return 1
    }
    cache[x] = fibonacci(x-1) + fibonacci(x-2)
	return cache[x]
}
func main() {
	var x int
    fmt.Scan(&x)
    fmt.Println(fibonacci(x))
}
