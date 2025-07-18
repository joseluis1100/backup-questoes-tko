package main
import "fmt"
func main() {
    var l, c, maior int
    fmt.Scan(&l, &c)
    matriz := make([][]int, l)
    for i := range matriz {
        matriz[i] = make([]int, c)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
        }
    }
    for i := range matriz {
        var soma int
        for j := range matriz[i] {
            soma += matriz[i][j]
        }
        if soma > maior {
            maior = soma
        }
    }
    for i := range c {
        var soma int
        for j := range matriz {
            soma += matriz[j][i]
        }
        if soma > maior {
            maior = soma
        }
    }
    fmt.Println(maior)
}
