package main
import "fmt"
func main() {
    var num_linhas, num_colunas, menores int
    fmt.Scan(&num_linhas, &num_colunas)
    matriz := make([][]int, num_linhas)
    for i := range matriz {
        matriz[i] = make([]int, num_colunas)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
        }
    }
    for i := range matriz {
        for j := range matriz[i] {
            if i > 0 && matriz[i][j] < matriz[i-1][j] {
                menores++
            }
        }
    }
    fmt.Println(menores)
}
