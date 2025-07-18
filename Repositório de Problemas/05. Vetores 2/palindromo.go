package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    palindromos := make([]int, 0)
    menor := int(math.Pow(10, float64(n-1)))
    maior := int(math.Pow(10, float64(n))-1)
    for i := maior; i >= menor; i-- {
        for j := maior; j >= menor; j-- {
            normal := []rune(strconv.Itoa(i*j))
            inverso := make([]rune, len(normal))
            copy(inverso, normal)
            slices.Reverse(inverso)
            if string(normal) == string(inverso) && !slices.Contains(palindromos, i * j) {
                palindromos = append(palindromos, i*j)
            }
        }
    }
    slices.Sort(palindromos)
    slices.Reverse(palindromos)
    for i := range m {
        fmt.Println(palindromos[i])
    }
}