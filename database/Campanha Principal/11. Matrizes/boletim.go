package main
import "fmt"
func main() {
    var soma int
    matriz := make([][]int, 2)
    for i := range matriz {
        matriz[i] = make([]int, 3)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
            soma += matriz[i][j]
        }
    }
    fmt.Println(soma)
}
