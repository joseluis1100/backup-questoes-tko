package main
import "fmt"
func main() {
    soma_linha := make([]int, 3)
    soma_coluna := make([]int, 3)
    matriz := make([][]int, 3)
    for i := range matriz {
        matriz[i] = make([]int, 3)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
            soma_linha[i] += matriz[i][j]
            soma_coluna[j] += matriz[i][j]
        }
    }
    diagonal1 := matriz[0][0] + matriz[1][1] + matriz[2][2]
    diagonal2 := matriz[2][0] + matriz[1][1] + matriz[0][2]
    for i := range 3 {
        if soma_coluna[i] != soma_linha[i] || soma_coluna[i] != diagonal1 || soma_coluna[i] != diagonal2 {
            fmt.Println("nao")
            return
        }
    }
    fmt.Println("sim")
}
