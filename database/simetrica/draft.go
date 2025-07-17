package main
import "fmt"
func main() {
    matriz := make([][]int, 3)
    for i := range matriz {
        matriz[i] = make([]int, 3)
    }
    matriz2 := make([][]int, 3)
    for i := range matriz2 {
        matriz2[i] = make([]int, 3)
    }
    for i := range matriz {
        for j := range matriz[i] {
            fmt.Scan(&matriz[i][j])
        }
    }
    for i := range matriz2 {
        for j := range matriz2[i] {
            matriz2[i][j] = matriz[j][i]
        }
    }
    igual := true
    for i := range matriz {
        for j := range matriz[i] {
            if !(matriz[i][j] == matriz2[i][j]) {
                igual = false
            }
        }
    }
    if igual {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
