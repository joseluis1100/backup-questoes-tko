package main
import "fmt"
func main() {
    matriz := make([][]int, 5)
    for i := range matriz {
        matriz[i] = make([]int, 5)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
        }
    }
    principal := matriz[0][0] + matriz[1][1] + matriz[2][2] + matriz[3][3] + matriz[4][4]
    secundaria := matriz[0][4] + matriz[1][3] + matriz[2][2] + matriz[3][1] + matriz[4][0]
    fmt.Println(principal - secundaria)
}
